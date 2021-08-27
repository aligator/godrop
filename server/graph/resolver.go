package graph

import "github.com/aligator/godrop/server/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	FileNodeService *service.FileNodeService
}
