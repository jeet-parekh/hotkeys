## Global HotKeys for Go (for Windows)

---

### About hotkeys

- Use hotkeys to set and listen for global hotkey events (for Windows).

---

### Example

- This program sets `alt + a` and `alt + b` hotkeys. It then listens for 10 hotkey events.

```go
package main

import (
	"fmt"
	"github.com/jeet-parekh/hotkeys"
)

func main() {
	hks := hotkeys.NewHotKeys(4)
	hotkeyMessages := hks.GetMessageChannel()
	hks.RegisterHotKey(1, hotkeys.MOD_ALT, hotkeys.VK_A)
	hks.RegisterHotKey(2, hotkeys.MOD_ALT, hotkeys.VK_B)
	hks.Start()

	for i := 0; i < 10; i+=1 {
		fmt.Printf("%+v\n", (<- hotkeyMessages).ID)
	}

	hks.Stop()
	hks.UnregisterAllHotKeys()
}
```

---

### Notes

- Modifier key codes and virtual key codes are provided as constants.

---
