package auth

import (
	"uuid"
)

type UserPrivilegeLevel int

// guests are people who will only be present temporarily or only need access to one service, 
// authorized would be people who are more frequent visitors but still shouldn't be able to manage devices, 
// and admin is self-explanatory
const (
	Guest UserPrivilegeLevel = iota
	Authorized
	Admin
)

var privilegeMap = map[UserPrivilegeLevel]string{
	Guest: "guest",
	Authorized: "authorized",
	Admin: "admin",
}

type UserIdentity struct {
	UserID int;
	UUID uuid.UUID;
	Username string;
	Privilege UserPrivilegeLevel
}