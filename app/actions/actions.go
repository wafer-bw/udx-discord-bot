package actions

import (
	"github.com/wafer-bw/discobottest/app/actions/debug"
	"github.com/wafer-bw/discobottest/app/actions/stocks"
	"github.com/wafer-bw/discobottest/app/errs"
	"github.com/wafer-bw/discobottest/app/models"
)

// Action is executed to resolve an interaction request
type Action func(request *models.InteractionRequest) (*models.InteractionResponse, error)

// todo - provide way of registering actions
// keys must be all lowercase
var actions = map[string]Action{
	"extrinsicrisk": stocks.ExtrinsicRisk,
	"debug":         debug.Debug,
}

// Run the action for the provided interaction request
func Run(interaction *models.InteractionRequest) (*models.InteractionResponse, error) {
	action, ok := actions[interaction.Data.Name]
	if !ok {
		return nil, errs.ErrNotImplemented
	}
	return action(interaction)
}
