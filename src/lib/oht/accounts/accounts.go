package accounts

import (
	"lib/oht/core"
)

type Accounts struct {
	oht       *oht.OHT
	manager   *Manager
	Interface *Interface
}

func InitializeAccounts(oht *oht.OHT) *Accounts {
	am := NewManager(oht.Interface.GenerateEncryptedKeyStore())
	return &Accounts{
		oht:       oht,
		manager:   am,
		Interface: &Interface{Manager: am},
	}
}
