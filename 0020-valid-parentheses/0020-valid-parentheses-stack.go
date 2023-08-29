type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}	
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}

const symbols = "[({])}"

func isValid(s string) bool {
	openSymbols := map[byte]int{
		symbols[0]: 0,
		symbols[1]: 1,
		symbols[2]: 2,
	}

	closeSymbols := map[byte]int{
		symbols[3]: 0,
		symbols[4]: 1,
		symbols[5]: 2,
	}

	var stack = New()

	for i := 0; i < len(s); i++ {
		val, found := closeSymbols[s[i]]
		if found {
			if stack.Len() == 0 {
				return false
			}

			curr := stack.Pop().(byte)
			if openSymbols[curr] != val {
				return false
			}
		} else {
			_, found = openSymbols[s[i]]
			if found {
				stack.Push(s[i])
			} else {
				return false
			}
		}
	}

	return stack.Len() == 0
}
