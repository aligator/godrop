package graph

import (
	"github.com/aligator/godrop/server/log"
	"github.com/aligator/godrop/server/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Logger          log.GoDropLogger
	FileNodeService *service.FileNodeService
}
