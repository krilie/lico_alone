package dig

import (
	"go.uber.org/dig"
)

type StructContainer struct {
	*dig.Container
}

var Container = &StructContainer{
	Container: dig.New(),
}

func (c *StructContainer) MustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	CheckErr(c.Container.Provide(constructor, opts...))
}
func (c *StructContainer) MustInvoke(function interface{}, opts ...dig.InvokeOption) {
	CheckErr(c.Container.Invoke(function, opts...))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
