//
// go.notify/freedesktop :: notify.go
//
//   Copyright (c) 2017 Akinori Hattori <hattya@gmail.com>
//
//   Permission is hereby granted, free of charge, to any person
//   obtaining a copy of this software and associated documentation files
//   (the "Software"), to deal in the Software without restriction,
//   including without limitation the rights to use, copy, modify, merge,
//   publish, distribute, sublicense, and/or sell copies of the Software,
//   and to permit persons to whom the Software is furnished to do so,
//   subject to the following conditions:
//
//   The above copyright notice and this permission notice shall be
//   included in all copies or substantial portions of the Software.
//
//   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//   EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
//   MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//   NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
//   BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
//   ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
//   CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//   SOFTWARE.
//

// Package freedesktop implements the Desktop Notifications Specification
// version 1.2.
//
// See https://developer.gnome.org/notification-spec/ for details.
package freedesktop

import (
	"fmt"
	"image"
	"math"
	"sync"

	"github.com/godbus/dbus"
)

const (
	addMatch                           = "org.freedesktop.DBus.AddMatch"
	notifications                      = "org.freedesktop.Notifications"
	notificationClosed                 = "org.freedesktop.Notifications.NotificationClosed"
	actionInvoked                      = "org.freedesktop.Notifications.ActionInvoked"
	objectPath         dbus.ObjectPath = "/org/freedesktop/Notifications"
)

// for testing
var (
	sessionBus  = dbus.SessionBus
	testHookNew func(*Client)
)

// Client is a notification client.
type Client struct {
	NotificationClosed chan NotificationClosed
	ActionInvoked      chan ActionInvoked

	conn   *dbus.Conn
	busObj dbus.BusObject
	obj    dbus.BusObject
	c      chan *dbus.Signal
	wg     sync.WaitGroup

	mu   sync.Mutex
	done chan struct{}
}

// New returns a new Client connected to the session bus.
func New() (*Client, error) {
	conn, err := sessionBus()
	if err != nil {
		return nil, err
	}
	c := &Client{
		NotificationClosed: make(chan NotificationClosed),
		ActionInvoked:      make(chan ActionInvoked),
		conn:               conn,
		busObj:             conn.BusObject(),
		obj:                conn.Object(notifications, objectPath),
		c:                  make(chan *dbus.Signal),
		done:               make(chan struct{}),
	}
	if testHookNew != nil {
		testHookNew(c)
	}
	// signal
	c.conn.Signal(c.c)
	for _, sig := range []string{"NotificationClosed", "ActionInvoked"} {
		if err := c.addMatch(sig); err != nil {
			return nil, err
		}
	}
	c.wg.Add(1)
	go c.signal()
	return c, nil
}

// Close closes the D-Bus connection.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	select {
	case <-c.done:
		return nil
	default:
	}

	close(c.done)
	c.wg.Wait()
	return c.conn.Close()
}

// CloseNotification closes and removes the notification of the specified ID.
func (c *Client) CloseNotification(id uint32) error {
	call := c.obj.Call("CloseNotification", 0, id)
	return call.Err
}

// GetCapabilities retrieves capabilities that the server implements.
//
// See https://developer.gnome.org/notification-spec/#command-get-capabilities
// for available capabilities.
func (c *Client) GetCapabilities() (caps []string, err error) {
	call := c.obj.Call("GetCapabilities", 0)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&caps)
	}
	return
}

// GetServerInformation retrieves the information of the server.
func (c *Client) GetServerInformation() (si ServerInfo, err error) {
	call := c.obj.Call("GetServerInformation", 0)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&si.Name, &si.Vendor, &si.Version, &si.SpecVersion)
	}
	return
}

// Notify sends a notification to the server.
func (c *Client) Notify(n *Notification) (id uint32, err error) {
	hints := make(map[string]dbus.Variant)
	if 0 < len(n.Hints) {
		var si ServerInfo
		si, err = c.GetServerInformation()
		if err != nil {
			return
		}
		var major, minor int
		if _, err = fmt.Sscanf(si.SpecVersion, "%v.%v", &major, &minor); err != nil {
			return
		}
		for k, v := range n.Hints {
			switch k {
			case "image-data":
				switch {
				case major > 1 || (major == 1 && minor >= 2):
					// >= 1.2
				case major == 1 && minor == 1:
					// =  1.1
					k = "image_data"
				default:
					// <  1.1
					k = "icon_data"
				}
			case "image-path":
				switch {
				case major > 1 || (major == 1 && minor >= 2):
					// >= 1.2
				default:
					// <  1.2
					k = "image_path"
				}
			}
			hints[k] = dbus.MakeVariant(v)
		}
	}

	call := c.obj.Call("Notify", 0, n.Name, n.ID, n.Icon, n.Summary, n.Body, n.Actions, hints, n.Timeout)
	if call.Err != nil {
		err = call.Err
	} else {
		err = call.Store(&id)
	}
	return
}

func (c *Client) addMatch(sig string) error {
	call := c.busObj.Call(addMatch, 0, fmt.Sprintf(`type='signal',interface='%v',member='%v'`, notifications, sig))
	return call.Err
}

func (c *Client) signal() {
	defer c.wg.Done()

	var closed chan NotificationClosed
	var invoked chan ActionInvoked
	closedBuf := make([]NotificationClosed, 1)
	invokedBuf := make([]ActionInvoked, 1)

	for {
		select {
		case sig := <-c.c:
			if sig != nil && sig.Path == objectPath {
				switch sig.Name {
				case notificationClosed:
					if closed == nil {
						closed = c.NotificationClosed
						closedBuf = closedBuf[1:]
					}
					closedBuf = append(closedBuf, NotificationClosed{
						ID:     sig.Body[0].(uint32),
						Reason: Reason(sig.Body[1].(uint32)),
					})
				case actionInvoked:
					if invoked == nil {
						invoked = c.ActionInvoked
						invokedBuf = invokedBuf[1:]
					}
					invokedBuf = append(invokedBuf, ActionInvoked{
						ID:  sig.Body[0].(uint32),
						Key: sig.Body[1].(string),
					})
				}
			}
		case closed <- closedBuf[0]:
			if len(closedBuf) == 1 {
				closed = nil
			} else {
				closedBuf = closedBuf[1:]
			}
		case invoked <- invokedBuf[0]:
			if len(invokedBuf) == 1 {
				invoked = nil
			} else {
				invokedBuf = invokedBuf[1:]
			}
		case <-c.done:
			return
		}
	}
}

// Notification represents a notification.
//
// See https://developer.gnome.org/notification-spec/#basic-design for details.
type Notification struct {
	Name    string                 // Application Name
	ID      uint32                 // Replaces ID
	Icon    string                 // Notification Icon
	Summary string                 // Summary
	Body    string                 // Body
	Actions []string               // Actions
	Hints   map[string]interface{} // Hints
	Timeout int32                  // Expiration Timeout
}

// Action adds (or replaces) the specified action to the Notification.
func (n *Notification) Action(key, label string) {
	for i := 0; i < len(n.Actions); i += 2 {
		if n.Actions[i] == key {
			n.Actions[i+1] = label
			return
		}
	}
	n.Actions = append(n.Actions, key, label)
}

// Hint adds (or replaces) the specified hint to the Notification.
//
// See https://developer.gnome.org/notification-spec/#hints for available
// hints.
func (n *Notification) Hint(name string, value interface{}) error {
	if n.Hints == nil {
		n.Hints = make(map[string]interface{})
	}
	switch name {
	case "image-data", "image_data", "icon_data":
		name = "image-data"
		switch v := value.(type) {
		case *ImageData:
		case ImageData:
			value = &v
		case image.Image:
			var err error
			if value, err = NewImageData(v); err != nil {
				return err
			}
		}
	case "image-path", "image_path":
		name = "image-path"
	case "x", "y":
		int2i := func(i int64) (int32, bool) {
			if math.MinInt32 <= i && i <= math.MaxInt32 {
				return int32(i), true
			}
			return 0, false
		}
		uint2i := func(u uint64) (int32, bool) {
			if u <= math.MaxInt32 {
				return int32(u), true
			}
			return 0, false
		}
		var i int32
		var ok bool
		switch v := value.(type) {
		case int:
			i, ok = int2i(int64(v))
		case int8:
			i, ok = int2i(int64(v))
		case int16:
			i, ok = int2i(int64(v))
		case int32:
			i, ok = v, true
		case int64:
			i, ok = int2i(v)
		case uint:
			i, ok = uint2i(uint64(v))
		case uint8:
			i, ok = uint2i(uint64(v))
		case uint16:
			i, ok = uint2i(uint64(v))
		case uint32:
			i, ok = uint2i(uint64(v))
		case uint64:
			i, ok = uint2i(v)
		default:
			return fmt.Errorf("%q is not int32: %T", name, value)
		}
		if !ok {
			return fmt.Errorf("%q overflows int32 range: %v", name, value)
		}
		value = i
	case "urgency":
		int2y := func(i int64) (uint8, bool) {
			if 0 <= i && i <= math.MaxUint8 {
				return uint8(i), true
			}
			return 0, false
		}
		uint2y := func(u uint64) (uint8, bool) {
			if u <= math.MaxUint8 {
				return uint8(u), true
			}
			return 0, false
		}
		var y uint8
		var ok bool
		switch v := value.(type) {
		case int:
			y, ok = int2y(int64(v))
		case int8:
			y, ok = int2y(int64(v))
		case int16:
			y, ok = int2y(int64(v))
		case int32:
			y, ok = int2y(int64(v))
		case int64:
			y, ok = int2y(v)
		case uint:
			y, ok = uint2y(uint64(v))
		case uint8:
			y, ok = v, true
		case uint16:
			y, ok = uint2y(uint64(v))
		case uint32:
			y, ok = uint2y(uint64(v))
		case uint64:
			y, ok = uint2y(v)
		default:
			return fmt.Errorf("%q is not byte: %T", name, value)
		}
		if !ok {
			return fmt.Errorf("%q overflows byte range: %v", name, value)
		}
		value = y
	}
	n.Hints[name] = value
	return nil
}

// ImageData represents a raw image data structure of signature (iiibiiay).
//
// See https://developer.gnome.org/notification-spec/#icons-and-images for
// details.
type ImageData struct {
	Width         int32
	Height        int32
	Stride        int32
	Alpha         bool
	BitsPerSample int32
	NumChannels   int32
	Data          []byte
}

// NewImageData returns a new raw image data structure from the specified img.
//
// img should be either an *image.Gray or an *image.NRGBA.
func NewImageData(img image.Image) (*ImageData, error) {
	size := img.Bounds().Size()
	data := &ImageData{
		Width:         int32(size.X),
		Height:        int32(size.Y),
		BitsPerSample: 8,
	}
	switch img := img.(type) {
	case *image.Gray:
		data.Stride = int32(img.Stride)
		data.NumChannels = 1
		data.Data = img.Pix
	case *image.NRGBA:
		data.Stride = int32(img.Stride)
		data.Alpha = true
		data.NumChannels = 4
		data.Data = img.Pix
	default:
		return nil, fmt.Errorf("unsupported image: %T", img)
	}
	return data, nil
}

// ServerInfo represents the information of a server.
type ServerInfo struct {
	Name        string
	Vendor      string
	Version     string
	SpecVersion string
}

// NotificationClosed represents a NotificationClosed signal.
type NotificationClosed struct {
	ID     uint32
	Reason Reason
}

// Reason represents a reason of the NotificationClosed signal.
type Reason uint32

// List of reasons for the NotificationClosed signal.
const (
	ReasonExpired Reason = iota + 1
	ReasonDismissed
	ReasonClosed
	ReasonUndefined
)

func (r Reason) String() string {
	switch r {
	case ReasonExpired:
		return "the notification expired"
	case ReasonDismissed:
		return "the notification was dismissed by the user"
	case ReasonClosed:
		return "the notification was closed by a call to CloseNotification"
	}
	return "undefined/reserved reasons"
}

// ActionInvoked represents an ActionInvoked signal.
type ActionInvoked struct {
	ID  uint32
	Key string
}
