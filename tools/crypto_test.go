package tools

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	md5 := MD5("123456")
	fmt.Println(md5)
}
