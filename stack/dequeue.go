package stack

type Dequeue struct {
	stackIn, stackOut *Stack
}

func NewDequeue() *Dequeue {
	return &Dequeue{
		stackIn:  NewStack(),
		stackOut: NewStack(),
	}
}
func (d *Dequeue) Push(v int64) {
	d.stackIn.Push(v)
}

func (d *Dequeue) Pop() int64 {
	if d.stackOut.Empty() {
		for !d.stackIn.Empty() {
			d.stackOut.Push(d.stackIn.Pop())
		}
	}

	return d.stackOut.Pop()
}

func (d *Dequeue) Empty() bool {
	return d.stackIn.Empty() && d.stackOut.Empty()
}
