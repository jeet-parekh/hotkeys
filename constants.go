package hotkeys

const (
	_SUCCESS string = "The operation completed successfully."

	MOD_ALT      uintptr = 0x0001
	MOD_CONTROL  uintptr = 0x0002
	MOD_SHIFT    uintptr = 0x0004
	MOD_WIN      uintptr = 0x0008
	MOD_NOREPEAT uintptr = 0x4000

	VK_0 uintptr = 0x30
	VK_1 uintptr = 0x31
	VK_2 uintptr = 0x32
	VK_3 uintptr = 0x33
	VK_4 uintptr = 0x34
	VK_5 uintptr = 0x35
	VK_6 uintptr = 0x36
	VK_7 uintptr = 0x37
	VK_8 uintptr = 0x38
	VK_9 uintptr = 0x39

	VK_A uintptr = 0x41
	VK_B uintptr = 0x42
	VK_C uintptr = 0x43
	VK_D uintptr = 0x44
	VK_E uintptr = 0x45
	VK_F uintptr = 0x46
	VK_G uintptr = 0x47
	VK_H uintptr = 0x48
	VK_I uintptr = 0x49
	VK_J uintptr = 0x4A
	VK_K uintptr = 0x4B
	VK_L uintptr = 0x4C
	VK_M uintptr = 0x4D
	VK_N uintptr = 0x4E
	VK_O uintptr = 0x4F
	VK_P uintptr = 0x50
	VK_Q uintptr = 0x51
	VK_R uintptr = 0x52
	VK_S uintptr = 0x53
	VK_T uintptr = 0x54
	VK_U uintptr = 0x55
	VK_V uintptr = 0x56
	VK_W uintptr = 0x57
	VK_X uintptr = 0x58
	VK_Y uintptr = 0x59
	VK_Z uintptr = 0x5A

	VK_LBUTTON             uintptr = 0x01
	VK_RBUTTON             uintptr = 0x02
	VK_CANCEL              uintptr = 0x03
	VK_MBUTTON             uintptr = 0x04
	VK_XBUTTON1            uintptr = 0x05
	VK_XBUTTON2            uintptr = 0x06
	VK_BACK                uintptr = 0x08
	VK_TAB                 uintptr = 0x09
	VK_CLEAR               uintptr = 0x0C
	VK_RETURN              uintptr = 0x0D
	VK_SHIFT               uintptr = 0x10
	VK_CONTROL             uintptr = 0x11
	VK_MENU                uintptr = 0x12
	VK_PAUSE               uintptr = 0x13
	VK_CAPITAL             uintptr = 0x14
	VK_KANA                uintptr = 0x15
	VK_HANGUEL             uintptr = 0x15
	VK_HANGUL              uintptr = 0x15
	VK_JUNJA               uintptr = 0x17
	VK_FINAL               uintptr = 0x18
	VK_HANJA               uintptr = 0x19
	VK_KANJI               uintptr = 0x19
	VK_ESCAPE              uintptr = 0x1B
	VK_CONVERT             uintptr = 0x1C
	VK_NONCONVERT          uintptr = 0x1D
	VK_ACCEPT              uintptr = 0x1E
	VK_MODECHANGE          uintptr = 0x1F
	VK_SPACE               uintptr = 0x20
	VK_PRIOR               uintptr = 0x21
	VK_NEXT                uintptr = 0x22
	VK_END                 uintptr = 0x23
	VK_HOME                uintptr = 0x24
	VK_LEFT                uintptr = 0x25
	VK_UP                  uintptr = 0x26
	VK_RIGHT               uintptr = 0x27
	VK_DOWN                uintptr = 0x28
	VK_SELECT              uintptr = 0x29
	VK_PRINT               uintptr = 0x2A
	VK_EXECUTE             uintptr = 0x2B
	VK_SNAPSHOT            uintptr = 0x2C
	VK_INSERT              uintptr = 0x2D
	VK_DELETE              uintptr = 0x2E
	VK_HELP                uintptr = 0x2F
	VK_LWIN                uintptr = 0x5B
	VK_RWIN                uintptr = 0x5C
	VK_APPS                uintptr = 0x5D
	VK_SLEEP               uintptr = 0x5F
	VK_NUMPAD0             uintptr = 0x60
	VK_NUMPAD1             uintptr = 0x61
	VK_NUMPAD2             uintptr = 0x62
	VK_NUMPAD3             uintptr = 0x63
	VK_NUMPAD4             uintptr = 0x64
	VK_NUMPAD5             uintptr = 0x65
	VK_NUMPAD6             uintptr = 0x66
	VK_NUMPAD7             uintptr = 0x67
	VK_NUMPAD8             uintptr = 0x68
	VK_NUMPAD9             uintptr = 0x69
	VK_MULTIPLY            uintptr = 0x6A
	VK_ADD                 uintptr = 0x6B
	VK_SEPARATOR           uintptr = 0x6C
	VK_SUBTRACT            uintptr = 0x6D
	VK_DECIMAL             uintptr = 0x6E
	VK_DIVIDE              uintptr = 0x6F
	VK_F1                  uintptr = 0x70
	VK_F2                  uintptr = 0x71
	VK_F3                  uintptr = 0x72
	VK_F4                  uintptr = 0x73
	VK_F5                  uintptr = 0x74
	VK_F6                  uintptr = 0x75
	VK_F7                  uintptr = 0x76
	VK_F8                  uintptr = 0x77
	VK_F9                  uintptr = 0x78
	VK_F10                 uintptr = 0x79
	VK_F11                 uintptr = 0x7A
	VK_F12                 uintptr = 0x7B
	VK_F13                 uintptr = 0x7C
	VK_F14                 uintptr = 0x7D
	VK_F15                 uintptr = 0x7E
	VK_F16                 uintptr = 0x7F
	VK_F17                 uintptr = 0x80
	VK_F18                 uintptr = 0x81
	VK_F19                 uintptr = 0x82
	VK_F20                 uintptr = 0x83
	VK_F21                 uintptr = 0x84
	VK_F22                 uintptr = 0x85
	VK_F23                 uintptr = 0x86
	VK_F24                 uintptr = 0x87
	VK_NUMLOCK             uintptr = 0x90
	VK_SCROLL              uintptr = 0x91
	VK_LSHIFT              uintptr = 0xA0
	VK_RSHIFT              uintptr = 0xA1
	VK_LCONTROL            uintptr = 0xA2
	VK_RCONTROL            uintptr = 0xA3
	VK_LMENU               uintptr = 0xA4
	VK_RMENU               uintptr = 0xA5
	VK_BROWSER_BACK        uintptr = 0xA6
	VK_BROWSER_FORWARD     uintptr = 0xA7
	VK_BROWSER_REFRESH     uintptr = 0xA8
	VK_BROWSER_STOP        uintptr = 0xA9
	VK_BROWSER_SEARCH      uintptr = 0xAA
	VK_BROWSER_FAVORITES   uintptr = 0xAB
	VK_BROWSER_HOME        uintptr = 0xAC
	VK_VOLUME_MUTE         uintptr = 0xAD
	VK_VOLUME_DOWN         uintptr = 0xAE
	VK_VOLUME_UP           uintptr = 0xAF
	VK_MEDIA_NEXT_TRACK    uintptr = 0xB0
	VK_MEDIA_PREV_TRACK    uintptr = 0xB1
	VK_MEDIA_STOP          uintptr = 0xB2
	VK_MEDIA_PLAY_PAUSE    uintptr = 0xB3
	VK_LAUNCH_MAIL         uintptr = 0xB4
	VK_LAUNCH_MEDIA_SELECT uintptr = 0xB5
	VK_LAUNCH_APP1         uintptr = 0xB6
	VK_LAUNCH_APP2         uintptr = 0xB7
	VK_OEM_1               uintptr = 0xBA
	VK_OEM_PLUS            uintptr = 0xBB
	VK_OEM_COMMA           uintptr = 0xBC
	VK_OEM_MINUS           uintptr = 0xBD
	VK_OEM_PERIOD          uintptr = 0xBE
	VK_OEM_2               uintptr = 0xBF
	VK_OEM_3               uintptr = 0xC0
	VK_OEM_4               uintptr = 0xDB
	VK_OEM_5               uintptr = 0xDC
	VK_OEM_6               uintptr = 0xDD
	VK_OEM_7               uintptr = 0xDE
	VK_OEM_8               uintptr = 0xDF
	VK_OEM_102             uintptr = 0xE2
	VK_PROCESSKEY          uintptr = 0xE5
	VK_PACKET              uintptr = 0xE7
	VK_ATTN                uintptr = 0xF6
	VK_CRSEL               uintptr = 0xF7
	VK_EXSEL               uintptr = 0xF8
	VK_EREOF               uintptr = 0xF9
	VK_PLAY                uintptr = 0xFA
	VK_ZOOM                uintptr = 0xFB
	VK_NONAME              uintptr = 0xFC
	VK_PA1                 uintptr = 0xFD
	VK_OEM_CLEAR           uintptr = 0xFE
)