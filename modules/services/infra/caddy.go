package infra

import (
	"encoding/json"
	"os"
	"syscall"
)

type CaddyService struct {
	Status ServiceStatus
	CaddyPort int // the port for the Caddy admin interface
	subdomainMap map[string]int // the map of subdomains to ports!
}