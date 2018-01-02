package plane

// A Button represents, like in Pixel, either keyboard or mouse
// button.
type Button uint

const (
	MouseUnknown Button = iota
	MouseOne
	MouseTwo
	MouseThree
)

const (
	KeyUnknown Button = iota

	KeyUp
	KeyDown
	KeyLeft
	KeyRight

	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ

	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9

	KeySpace
	KeyEnter
	KeySlash
	KeySemicolon
	KeyMinus
	KeyEquals
	KeyBackslash
	KeyComma
	KeyPeriod
	KeyTab
	KeyEscape

	KeyCapsLock
	KeyNumLock
	KeyScrollLock

	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12

	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPageUp
	KeyPageDown

	KeyShift
	KeyControl
	KeyAlt
)
