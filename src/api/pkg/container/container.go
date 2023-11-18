package main

import (
	"fmt"
	"reflect"
)

func NewContainer() container {
	return container{}
}

type container map[reflect.Type]func() reflect.Value

func (c container) Bind(resolver interface{}) {

	r := reflect.TypeOf(resolver)

	args := make([]reflect.Value, r.NumIn())

	for i := 0; i < r.NumIn(); i++ {
		args[i] = c.callResolver(r.In(i))
	}
	for i := 0; i < r.NumOut(); i++ {
	}

	c[r.Out(0)] = func() reflect.Value {
		return reflect.ValueOf(resolver).Call(args)[0]
	}
	fmt.Println("bind", r.Out(0))
}

func (c container) Make(intf interface{}) {

	intfReflect := reflect.TypeOf(intf)

	generatedIncetance := c.callResolver(intfReflect.Elem())

	reflect.ValueOf(intf).Elem().Set(generatedIncetance)
}

func (c container) callResolver(key reflect.Type) reflect.Value {

	return c[key]()
}

func main() {

	c := NewContainer()

	c.Bind(func() TestIntf2 {
		return &TestImpl2{}
	})

	c.Bind(func(test2 TestIntf2) TestIntf {
		return &TestImpl{
			test2: test2,
		}
	})
	var t TestIntf
	c.Make(&t)

	t.M()
}

type TestIntf interface {
	M()
}

type TestImpl struct {
	test2 TestIntf2
}

func (t TestImpl) M() {
	l("hello")
	t.test2.M2()
}

type TestIntf2 interface {
	M2()
}

type TestImpl2 struct {
}

func (TestImpl2) M2() {
	l("hello2")
}

func l(v ...interface{}) {
	fmt.Println(v...)
}
