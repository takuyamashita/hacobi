package container_test

import (
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
)

type Intf1 interface {
	ABC() int
	GetN() int
}

type Impl1 struct {
	Intf2 Intf2
	N     int
}

func (impl1 *Impl1) ABC() int {
	impl1.N++
	return impl1.Intf2.DEF()
}

func (impl1 Impl1) GetN() int {
	return impl1.N
}

type Intf2 interface {
	DEF() int
}

type Impl2 struct {
	n int
}

func (impl2 *Impl2) DEF() int {
	impl2.n = impl2.n + 1
	return impl2.n
}

func TestBind(t *testing.T) {

	c := container.NewContainer()

	c.Bind(func() Intf2 {
		return &Impl2{}
	})

	c.Bind(func(intf2 Intf2) Intf1 {
		return &Impl1{
			Intf2: intf2,
		}
	})

	var intf1 Intf1

	c.Make(&intf1)

	n := intf1.ABC()
	if n != 1 {
		t.Errorf("intf1.ABC() = %v, want %v", n, 1)
	}

	n = intf1.ABC()
	if n != 2 {
		t.Errorf("intf1.ABC() = %v, want %v", n, 2)
	}
}

func TestBindSingle(t *testing.T) {

	c := container.NewContainer()

	c.Bind(func() Intf2 {
		return &Impl2{}
	})

	c.BindSingle(func(intf2 Intf2) Intf1 {
		return &Impl1{
			Intf2: intf2,
		}
	})

	var intf1_1 Intf1

	c.Make(&intf1_1)

	var intf1_2 Intf1

	c.Make(&intf1_2)

	intf1_1.ABC()
	intf1_1.ABC()

	if intf1_1.GetN() != intf1_2.GetN() {
		t.Errorf("intf1_1 = %d, intf1_2 = %d", intf1_1.GetN(), intf1_2.GetN())
	}
}
