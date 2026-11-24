package auth

import (
	"strings"
	"bytes"
	"uuid"
)

type DeviceIdentity struct {
	DeviceID int;
	UUID uuid.UUID
}
