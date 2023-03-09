package main

type calculator struct {
	numberOne int
	numberTwo int
}

func newCalculator() calculator {
	c := calculator{}
	return c
}

func (c *calculator) setNumberOne(numberOne int) {
	c.numberOne = numberOne
}
func (c *calculator) setNumberTwo(numberTwo int) {
	c.numberTwo = numberTwo
}
func (c *calculator) getNumberTwo() int {
	return c.numberTwo
}
func (c *calculator) getNumberOne() int {
	return c.numberOne
}

func (c *calculator) AddNumbers() int {
	return c.numberOne + c.numberTwo
}
func (c *calculator) MinusNumbers() int {
	return c.numberOne - c.numberTwo
}
func (c *calculator) MultiplyNumbers() int {
	return c.numberOne * c.numberTwo
}
func (c *calculator) DivideNumbers() int {
	return c.numberOne / c.numberTwo
}
