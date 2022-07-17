package access

import (
	"pandora/ent"
)

type RBAC interface {
	HasAccess(user, url, method string) bool
}

func NewRBAC(db *ent.Client) RBAC {
	return NewCasbinRBAC(db)
}

//var ProviderSet = wire.NewSet(NewRBAC)
