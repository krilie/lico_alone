package dig

import (
	"fmt"
	"go.uber.org/dig"
	"testing"
)

type Hello interface{ Hello() }
type HelloOne struct{}
type HelloTwo struct{}

func NewHelloOne() Hello   { return &HelloOne{} }
func NewHelloTwo() Hello   { return &HelloTwo{} }
func (h *HelloOne) Hello() { fmt.Println("form hello one") }
func (h *HelloTwo) Hello() { fmt.Println("form hello two") }

func TestDig(t *testing.T) {
	container := dig.New()
	err := container.Provide(NewHelloOne)
	t.Log(err)
	err = container.Invoke(func(hello Hello) {
		hello.Hello()
	})
	t.Log(err)
	err = container.Provide(NewHelloTwo)
	t.Log(err)
	err = container.Invoke(func(hello Hello) {
		hello.Hello()
	})
	t.Log(err)
}
