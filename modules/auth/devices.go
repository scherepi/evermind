package auth

import (
	"strings"
	"bytes"
	"uuid"
)

type DeviceIdentity struct {
	DeviceID int;
	UUID uuid.UUID
	Description string // a brief description identifying this device, like "guac's phone" or "the M5Stick used for local inference"
}
