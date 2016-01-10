package lesson07

//Original Problem:
//https://codility.com/programmers/task/nesting/
func Nesting(S string) int {

	if len(S) == 0 {
		return 1
	}

	if len(S) == 1 {
		return 0
	}

	queue := StackADT{Stack: ``}

	dequeuedElement := ``

	i := 0
	for i < len(S) {
		currentChar := S[i : i+1]
		if currentChar == `(` {
			queue.enStack(currentChar)
		}

		if currentChar == `)` {
			dequeuedElement = queue.deStack()
			if dequeuedElement == `` {
				return 0
			}
		}
		i++
	}

	if queue.isEmpty() == true {
		return 1
	}
	return 0
}

type StackADT struct {
	Stack string
}

func (q *StackADT) enStack(char string) {
	q.Stack += char
}

func (q *StackADT) deStack() string {
	if q.isEmpty() == true {
		return ``
	}

	element := q.Stack[len(q.Stack)-1:]
	q.Stack = q.Stack[:len(q.Stack)-1]
	return element
}

func (q *StackADT) size() int {
	return len(q.Stack)
}

func (q *StackADT) isEmpty() bool {
	if q.size() == 0 {
		return true
	}
	return false
}
