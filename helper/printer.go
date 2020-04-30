package helper

import (
	"fmt"
	"time"
)

const (
	NOTICE int = iota
	LOG    int = iota
)

type seqF func() int

func seqnumFunc() seqF {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Printer struct {
	seqnum seqF
	ID     int64
}

func NewPrinter() *Printer {
	p := Printer{}
	p.Init()
	return &p
}

func (t *Printer) Init() {
	t.seqnum = seqnumFunc()
	t.ID = time.Now().UnixNano()
}

func (t *Printer) Trace(str string) {
	fmt.Printf("[%v-%v] %v\n", t.ID, t.seqnum(), str)
}
