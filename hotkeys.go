package hotkeys

import (
	"github.com/jeet-parekh/winapi"
)

type (
	// HotKey contains information about a single hotkey.
	HotKey struct {
		ID             uintptr
		ModifierKeys   uintptr
		VirtualKeyCode uintptr
	}

	// HotKeys provides an interface to use set and use hotkeys.
	HotKeys struct {
		registeredHotKeys map[uintptr]*HotKey
		messages          chan *HotKeyEvent
		enabled           bool
	}

	// HotKeyEvent contains information about a hotkey event.
	HotKeyEvent struct {
		ID          uintptr
		KeysPressed uintptr
		Time        uint32
		CursorX     int32
		CursorY     int32
	}
)

// NewHotKeys creates and returns a new HotKeys object.
// The parameter is the buffer size of the channel through which the messages would be sent.
func NewHotKeys(bufferSize uint) *HotKeys {
	return &HotKeys{
		registeredHotKeys: make(map[uintptr]*HotKey),
		messages:          make(chan *HotKeyEvent, bufferSize),
	}
}

// GetMessageChannel returns the channel through which the messages would be sent.
func (hks *HotKeys) GetMessageChannel() <-chan *HotKeyEvent {
	return hks.messages
}

// GetRegisteredHotKeys returns a ID -> HotKey map of registered hotkeys.
func (hks *HotKeys) GetRegisteredHotKeys() map[uintptr]*HotKey {
	return hks.registeredHotKeys
}

// IsIDRegistered returns a boolean indicating whether or not a hotkey with the given ID is registered.
func (hks *HotKeys) IsIDRegistered(id uintptr) bool {
	if hks.registeredHotKeys[id] == nil {
		return false
	}
	return true
}

// RegisterHotKey registers a hotkey.
// First parameter is the ID of the hotkey.
// Second parameter is the modifier keys of the hotkey. They are the keys that must be pressed in combination with the key specified by the third parameter in order to activate the hotkey.
// Third parameter is the virtual key code of the hotkey.
func (hks *HotKeys) RegisterHotKey(id uintptr, modifierKeys uintptr, virtualKeyCode uintptr) error {
	_, err := winapi.RegisterHotKey(0, id, modifierKeys, virtualKeyCode)
	if err.Error() != _SUCCESS {
		return err
	}

	hks.registeredHotKeys[id] = &HotKey{
		ID:             id,
		ModifierKeys:   modifierKeys,
		VirtualKeyCode: virtualKeyCode,
	}
	return nil
}

// UnregisterHotKey unregisters a hotkey by its ID.
func (hks *HotKeys) UnregisterHotKey(id uintptr) error {
	if hks.IsIDRegistered(id) {
		_, err := winapi.UnregisterHotKey(0, id)
		if err.Error() != _SUCCESS {
			return err
		}
		delete(hks.registeredHotKeys, id)
	}
	return nil
}

// UnregisterAllHotKeys unregisters all the registered hotkeys.
func (hks *HotKeys) UnregisterAllHotKeys() error {
	for id := range hks.registeredHotKeys {
		err := hks.UnregisterHotKey(id)
		if err != nil {
			return err
		}
	}
	return nil
}

// Start receiving messages through the channel.
func (hks *HotKeys) Start() {
	hks.enabled = true
	go func() {
		for {
			var msg winapi.MSG
			winapi.GetMessage(&msg, 0, 0, 0)
			if msg.Message == 0x0312 {
				if hks.enabled {
					hks.messages <- &HotKeyEvent{
						ID:          msg.WParam,
						KeysPressed: msg.LParam,
						Time:        msg.Time,
						CursorX:     msg.Pt.X,
						CursorY:     msg.Pt.Y,
					}
				} else {
					break
				}
			}
		}
	}()
}

// Stop receiving messages through the channel - the message channel is closed. This function does not unregister the hotkeys.
func (hks *HotKeys) Stop() {
	hks.enabled = false
	close(hks.messages)
}
