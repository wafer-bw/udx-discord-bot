package actions

import (
	"github.com/wafer-bw/discobottest/app/actions/stocks"
	"github.com/wafer-bw/discobottest/app/errs"
	"github.com/wafer-bw/discobottest/app/interactions"
)

// Action is executed to resolve an interaction request
type Action func(request *interactions.InteractionRequest) (*interactions.InteractionResponse, error)

// todo - provide way of registering actions
// keys must be all lowercase
var actions = map[string]Action{
	"extrinsicrisk": stocks.ExtrinsicRisk,
}

// Run the action for the provided interaction request
func Run(interaction *interactions.InteractionRequest) (*interactions.InteractionResponse, error) {
	action, ok := actions[interaction.Data.Name]
	if !ok {
		return nil, errs.ErrNotImplemented
	}
	return action(interaction)
}
