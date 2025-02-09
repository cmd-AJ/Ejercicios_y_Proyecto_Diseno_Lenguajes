package postfix

type Symbol struct {
	// Raw character
	value string
	// Precedence. The bigger, the more to the left is place postfix notation.
	precedence int
	// If the symbol its an operator
	isOperator bool
	// Number of operands
	operands int
}

func (s *Symbol) String() string {
	return s.value
}

const ESCAPE_SYMBOL string = "\\"
const CONCAT_SYMBOL string = "·"

var OPERATORS = map[string]Symbol{
	")": {value: ")", precedence: 10, isOperator: true, operands: 1},
	"(": {value: "(", precedence: 10, isOperator: true, operands: 0},
	"]": {value: "]", precedence: 10, isOperator: true, operands: 1},
	"[": {value: "[", precedence: 10, isOperator: true, operands: 0},
	"|": {value: "|", precedence: 20, isOperator: true, operands: 2},
	"·": {value: "·", precedence: 30, isOperator: true, operands: 2},
	"?": {value: "?", precedence: 40, isOperator: true, operands: 1},
	"*": {value: "*", precedence: 40, isOperator: true, operands: 1},
	"+": {value: "+", precedence: 40, isOperator: true, operands: 1},
	"^": {value: "^", precedence: 50, isOperator: true, operands: 2},
}
