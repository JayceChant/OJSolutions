package main

import "strconv"

// 0ms, 100%
// 2.3MB, 100%
func isNumber(s string) bool {
	c := numChecker{str: s}
	return c.isNum()
}

type numChecker struct {
	str string
	i   int
}

func (c *numChecker) isNum() bool {
	return c.start() && c.mustReal() && c.maybeExponent() && c.end()
}

// consume spaces and judge if end only
func (c *numChecker) start() bool {
	c.trim()

	if c.i >= len(c.str) {
		return false
	}

	return true
}

// consume spaces and judge if end only
func (c *numChecker) end() bool {
	c.trim()

	if c.i != len(c.str) {
		return false
	}

	return true
}

func (c *numChecker) trim() {
	for c.i < len(c.str) && c.str[c.i] == ' ' {
		c.i++
	}
}

func (c *numChecker) mustReal() bool {
	c.signed()
	hasInt := c.mustInteger()
	c.dot()
	hasDecimal := c.mustInteger()
	return hasInt || hasDecimal
}

// match complete exponent or none
func (c *numChecker) maybeExponent() bool {
	if c.i == len(c.str) || c.str[c.i] != 'e' {
		return true
	}
	c.i++

	c.signed()
	return c.mustInteger()
}

// consume at most 1 sign only
func (c *numChecker) signed() {
	if c.i < len(c.str) && (c.str[c.i] == '+' || c.str[c.i] == '-') {
		c.i++
	}
}

// consume at least 1 digit
func (c *numChecker) mustInteger() bool {
	if c.i == len(c.str) || c.str[c.i] < '0' || c.str[c.i] > '9' {
		return false
	}

	c.i++

	for c.i < len(c.str) && (c.str[c.i] >= '0' && c.str[c.i] <= '9') {
		c.i++
	}
	return true
}

func (c *numChecker) dot() {
	if c.i < len(c.str) && c.str[c.i] == '.' {
		c.i++
	}
}

// ========== baseline ==========

func isNumHelper(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
