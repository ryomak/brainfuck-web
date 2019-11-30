package bf

import (
	"regexp"
)

// state is bf lexer
type State struct {
	cs            *Commands
	inputCommands []string
	memory        []byte
	pointer       int
}

var r *regexp.Regexp

// InitState is setting state
func InitState(c Commands, maxCommandNum int) *State {
	state := new(State)
	state.cs = &c
	state.memory = make([]byte, maxCommandNum)
	state.pointer = 0
	state.parse()
	return state
}

// parse is inputing commands with regexp
func (s *State) parse() {
	r = regexp.MustCompile(`(` + s.cs.Ops() + `)`)
}

// Start is excuting bf file
func (s *State) Start(inputStr string) string {
	var output string
	inputCommands := r.FindAllString(inputStr, -1)
	turn := 0
	for turn < len(inputCommands) {
		switch s.inputCommands[turn] {
		case s.cs.NEXT:
			s.pointer++
		case s.cs.PREV:
			s.pointer--
		case s.cs.INC:
			s.memory[s.pointer]++
		case s.cs.DEC:
			s.memory[s.pointer]--
		case s.cs.WRITE:
			output += string(s.memory[s.pointer])
		case s.cs.OPEN:
			if s.memory[s.pointer] == 0 {
				for depth := 1; depth > 0; {
					turn++
					nCommand := s.inputCommands[turn]
					if nCommand == s.cs.OPEN {
						depth++
					}
					if nCommand == s.cs.CLOSE {
						depth--
					}
				}
			}
		case s.cs.CLOSE:
			for depth := 1; depth > 0; {
				turn--
				nCommand := s.inputCommands[turn]
				if nCommand == s.cs.OPEN {
					depth--
				}
				if nCommand == s.cs.CLOSE {
					depth++
				}
			}
			turn--
		}
		turn++
	}
	return output
}
