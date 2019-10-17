package enum

// enumに使用する型宣言
type State int

// Stateに使用する定数を宣言
const (
	_ State = iota
	OPEN
	DOING
	CLOSE
)

// Stringer interfaceを実装
//func (state State) String() string  {
//	switch state {
//	case OPEN:
//		return "OPEN"
//	case DOING:
//		return "DOING"
//	case CLOSE:
//		return "CLOSE"
//	}
//	panic("Unknown State value.")
//}

// ツールを使用して実装
// $ go get golang.org/x/tools/cmd/stringer
// $ stringer -type State state.go
// enum/state_string.go
