package handlers

import "todoapp/services"

// Injection of Services for Handlers usage
type Handlers struct {
	Services services.IServices
}

// Constructor (used in router.go)
func NewHandlers(services services.IServices) *Handlers {
	return &Handlers{
		Services: services,
	}
}
