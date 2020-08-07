package main

import "strconv"

// 0ms
// 2.6MB
func evaluate(expression string) int {
	expr = expression
	val, _ := eval(&bindings{
		data: make(map[string]int),
	}, 0)
	return val
}

var (
	expr string
)

// [start, end)
func eval(bind *bindings, start int) (int, int) {
	defer bind.returnLending()
	start++ // skip '('

	fn, start := readName(start)

	if fn == "let" {
		return let(bind, start)
	}

	return calc(bind, start, fn == "mult")
}

func let(bind *bindings, start int) (val int, end int) {
	for {
		var name string
		var err error

		start++ // skip ' '
		if expr[start] == '(' {
			val, start = eval(bind.subBind(), start)
		} else {
			name, start = readName(start)
			val, err = strconv.Atoi(name)
			if err != nil {
				val = bind.getVal(name)
			}
		}

		if expr[start] == ')' {
			end = start + 1 // skip ')'
			return
		}

		start++ // skip ' '
		if expr[start] == '(' {
			val, start = eval(bind.subBind(), start)
		} else {
			var tempName string
			tempName, start = readName(start)
			val, err = strconv.Atoi(tempName)
			if err != nil {
				val = bind.getVal(tempName)
			}
		}
		bind.setVal(name, val)
	}
}

// add or mult
func calc(bind *bindings, start int, mult bool) (val int, end int) {
	var temp int
	var name string
	var err error

	if mult {
		val = 1
	}

	for expr[start] != ')' {
		start++ // skip ' '
		if expr[start] == '(' {
			temp, start = eval(bind.subBind(), start)
		} else {
			name, start = readName(start)
			temp, err = strconv.Atoi(name)
			if err != nil {
				temp = bind.getVal(name)
			}
		}

		if mult {
			val *= temp
		} else { // add
			val += temp
		}
	}
	end = start + 1 // skip ')'
	return
}

// [start, end)
func readName(start int) (name string, end int) {
	end = start
	for expr[end] != ' ' && expr[end] != ')' {
		end++
	}
	name = expr[start:end]
	return
}

type bindings struct {
	parent *bindings
	data   map[string]int
	lent   int
}

func (b *bindings) setVal(name string, val int) {
	b.data[name] = val
}

func (b *bindings) getVal(name string) int {
	val, ok := b.data[name]
	if !ok && b.parent != nil {
		val = b.parent.getVal(name)
	}
	return val
}

func (b *bindings) subBind() *bindings {
	if len(b.data) == 0 {
		b.lent++
		return b
	}

	return &bindings{
		parent: b,
		data:   make(map[string]int),
	}
}

func (b *bindings) returnLending() {
	if b.lent > 0 {
		if len(b.data) > 0 {
			b.data = make(map[string]int)
		}
		b.lent--
	}
}
