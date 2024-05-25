package auth

import (
	"github.com/ohwin/core/global"
)

type Police struct {
	Role   string
	Url    string
	Method string
}

// AddPolicy 添加Police
func AddPolicy(p Police) (bool, error) {
	return global.Enforce.AddPolicy(p.Role, p.Url, p.Method)
}

// DelPolicy 删除Policy
func DelPolicy(p Police) (bool, error) {
	return global.Enforce.RemovePolicy(p.Role, p.Url, p.Method)
}

// UpdatePolicy 更新Policy
func UpdatePolicy(old Police, new Police) (bool, error) {
	return global.Enforce.UpdatePolicy([]string{old.Role, old.Url, old.Method}, []string{new.Role, new.Url, new.Method})
}
