package queuetwostacks

type CQueue struct {
	primary  stack
	overflow stack
}

func NewCQueue() CQueue {
	return CQueue{Stack(), Stack()}
}

type stack struct {
	data *[]int
}

func Stack() stack {
	data := make([]int, 0)
	return stack{data: &data}
}

func (s *stack) pop() int {
	if len(*s.data) < 1 {
		return -1
	}
	var e int
	e, *s.data = (*s.data)[len(*s.data)-1:][0], (*s.data)[:len(*s.data)-1]
	return e
}

func (s *stack) push(val int) {
	*s.data = append(*s.data, val)
}

func (s *stack) top() int {
	if len(*s.data) < 1 {
		return -1
	}
	return (*s.data)[len(*s.data)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s.data) == 0
}

func (c *CQueue) Pop() int {
	if c.primary.isEmpty() {
		return -1
	}
	for !c.primary.isEmpty() {
		c.overflow.push(c.primary.pop())
	}
	res := c.overflow.pop()
	for !c.overflow.isEmpty() {
		c.primary.push(c.overflow.pop())
	}
	return res
}

func (c *CQueue) Push(val int) {
	c.primary.push(val)
}

func (c *CQueue) Empty() bool {
	return c.primary.isEmpty()
}

func (c *CQueue) Peek() int {
	if c.primary.isEmpty() {
		return -1
	}
	for !c.primary.isEmpty() {
		c.overflow.push(c.primary.pop())
	}
	res := c.overflow.top()
	for !c.overflow.isEmpty() {
		c.primary.push(c.overflow.pop())
	}
	return res
}
