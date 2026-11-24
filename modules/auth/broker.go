package auth

import (
	"fmt"
	"time"
)

// records every single transaction that happens with regards to global auth
type AuthLedger struct {

}

type AuthLedgerEntry struct {
	TxType 
}

// the AuthBroker reads and exercises auth rules, and is the arbiter for who can access what services and devices, and which devices can interact with which services.

type AuthBroker struct {

}
