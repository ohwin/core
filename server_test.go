package core

import (
	"fmt"
	"github.com/ohwin/core/global"
	"reflect"
	"testing"
)

func TestRunWindowsServer(t *testing.T) {

	a := make([]int, 0)
	c := CC(a)
	fmt.Println("--->", c.(string))
	Init()
	global.Redis.Set("sdd", "s", 0)

}

func CC(a interface{}) any {
	fmt.Println(reflect.TypeOf(a))
	return nil
}
