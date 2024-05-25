package core

import (
	"github.com/ohwin/core/global"
	"testing"
)

func TestRunWindowsServer(t *testing.T) {

	Init()
	global.Redis.Set("sdd", "s", 0)

}
