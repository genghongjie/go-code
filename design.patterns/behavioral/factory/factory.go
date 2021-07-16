package factory

import "fmt"

type Company interface {
	Goodjob()
}

type Thinry struct {
}

func (m *Thinry) Goodjob() {
	fmt.Println("Thinry 9点开始进入工作")
}

type Shine struct {
}

func (m *Shine) Goodjob() {
	fmt.Println("Shine 9点30开始进入工作")
}
func NewConpany(cType string) Company {

	switch cType {
	case "jit":
		return &Thinry{}
	case "shine":
		return &Shine{}
	}
	return nil
}
