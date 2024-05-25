package core

import (
	"fmt"
	"github.com/ohwin/core/global"
	"github.com/ohwin/core/types"
	"reflect"
	"testing"
)

func TestRunWindowsServer(t *testing.T) {

	Init()
	global.Redis.Set("sdd", "s", 0)
	RunWindowsServer([]types.RouterFunc{})
}

func CC(a interface{}) any {
	fmt.Println(reflect.TypeOf(a))
	return nil
}
