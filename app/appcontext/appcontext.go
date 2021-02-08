package appcontext

import (
	"github.com/wafer-bw/discobottest/app/config"
	"github.com/wafer-bw/discobottest/app/interactions"
)

// Deps defines `AppContext` dependencies
type Deps struct{}

// AppContext implements `AppContext` properties
type AppContext struct {
	Interactions interactions.Interactions
}

// New returns a new `AppContext` struct
func New(deps *Deps, conf *config.Config) *AppContext {
	return &AppContext{
		Interactions: interactions.New(),
	}
}
