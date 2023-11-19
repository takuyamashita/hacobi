package container

import (
	"log"
	"reflect"
)

type Container interface {

	// BindSingleは、resolverが返すインスタンスをシングルトンとして登録します
	// 1つのコンテナインスタンスに対してシングルになります(=複数のコンテナインスタンスを作成した場合、それぞれのコンテナインスタンスに対してシングルトンになります)
	BindSingle(resolver interface{})

	// Bindは、resolverが返すインスタンスをresolverの返り値の型で登録します
	Bind(resolver interface{})

	// Makeは、intfに対して、コンテナに登録されているresolverを使ってインスタンスを生成しintfにセットします
	Make(intf interface{})
}

func NewContainer() container {
	return container{}
}

type resolverInfo struct {
	isSingleton bool
	resolver    func() reflect.Value
	incetance   interface{}
}

type container map[reflect.Type]*resolverInfo

func (c container) BindSingle(resolver interface{}) {
	c.bind(resolver, true)
}
func (c container) Bind(resolver interface{}) {
	c.bind(resolver, false)
}

func (c container) bind(resolver interface{}, isSingleton bool) {

	r := reflect.TypeOf(resolver)

	if r == nil || r.Kind() != reflect.Func || r.NumOut() != 1 || r.Out(0).Kind() != reflect.Interface {
		log.Fatal("resolverは引数が0個または1つ以上のinterfaceで、返り値がinterface1つの関数である必要があります")
	}

	c[r.Out(0)] = &resolverInfo{
		resolver: func() reflect.Value {

			args := make([]reflect.Value, r.NumIn())
			for i := 0; i < r.NumIn(); i++ {
				args[i] = c.retrieveInsetance(r.In(i))
			}

			return reflect.ValueOf(resolver).Call(args)[0]
		},
		isSingleton: isSingleton,
	}
}

func (c container) Make(intf interface{}) {

	intfReflect := reflect.TypeOf(intf)

	generatedIncetance := c.retrieveInsetance(intfReflect.Elem())

	reflect.ValueOf(intf).Elem().Set(generatedIncetance)
}

func (c container) retrieveInsetance(key reflect.Type) reflect.Value {

	if _, ok := c[key]; !ok {
		log.Fatalf("%sに対応するresolverがありません", key.String())
	}

	if c[key].isSingleton {
		if c[key].incetance == nil {
			c[key].incetance = c[key].resolver().Interface()
		}
		return reflect.ValueOf(c[key].incetance)
	}

	return c[key].resolver()
}
