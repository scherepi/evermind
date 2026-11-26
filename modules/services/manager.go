package services

import (
	"uuid"
)

type ServiceManager struct {

}

type ServiceStatus int;

type Service interface {
	GetStatus() ServiceStatus
	GetServiceID() uuid.UUID
	Start() bool
	Suspend() bool
	Stop() bool
}