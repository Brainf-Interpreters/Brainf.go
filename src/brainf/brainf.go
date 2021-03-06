package brainf

import (
	"fmt"
	"io/ioutil"
)

func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = append((*alist)[:f-1])
	return rv
}

// Node Struct
type Node struct {
	kind byte
	op   int
}

// Parse code
func Parse(code string) []Node {
	var stack []int
	var result []Node
	i := 0

	st := ""
	for _, chr := range code {
		if chr == '+' || chr == '-' ||
			chr == '<' || chr == '>' ||
			chr == '.' || chr == ',' ||
			chr == '[' || chr == ']' {
			st += string(chr)
		}
	}

	code = st

	for i < len(code) {
		chr := code[i]

		if chr == '+' || chr == '-' ||
			chr == '<' || chr == '>' ||
			chr == '.' || chr == ',' {
			result = append(result, Node{kind: chr})
		} else {
			if chr == '[' {
				stack = append(stack, i)
				result = append(result, Node{kind: chr})
			} else {
				if chr == ']' {
					if len(stack) == 0 {
						panic("Missing '[' before ']'")
					}
					last := pop(&stack)
					(&result[last]).op = i
					result = append(result, Node{kind: chr, op: last})
				}
			}
		}

		i++
	}

	return result
}

// Run code
func Run(nodes []Node) (map[int]int, int) {
	data := map[int]int{0: 0}
	pointer := 0

	i := 0
	for i < len(nodes) {
		node := nodes[i]
		switch node.kind {
		case '+':
			if data[pointer] == 255 {
				data[pointer] = 0
			} else {
				data[pointer]++
			}

		case '-':
			if data[pointer] == 0 {
				data[pointer] = 255
			} else {
				data[pointer]--
			}

		case '>':
			if pointer == 255 {
				pointer = 0
			} else {
				pointer++
			}
			_, a := data[pointer]
			if a == false {
				data[pointer] = 0
			}

		case '<':
			if pointer == 0 {
				pointer = 255
			} else {
				pointer--
			}
			_, a := data[pointer]
			if a == false {
				data[pointer] = 0
			}

		case '.':
			fmt.Print(string(data[pointer]))

		case ',':
			var x []byte
			fmt.Scan(&x)
			data[pointer] = int(x[0])

		case '[':
			if data[pointer] == 0 {
				i = node.op
			}

		case ']':
			if data[pointer] != 0 {
				i = node.op
			}
		}

		i++
	}

	fmt.Println()

	return data, pointer
}

// RunCode runs a code
func RunCode(code string) {
	fmt.Println(Run(Parse(code)))
}

// RunFile runs a file
func RunFile(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot open file:", filename)
	} else {
		RunCode(string(dat))
	}
}
