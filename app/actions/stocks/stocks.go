package stocks

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/wafer-bw/discobottest/app/interactions"
)

type payload struct {
	Share  float64
	Strike float64
	Ask    float64
}

// ExtrinsicRisk calculates the extrinsic risk of an option provided the
// `share` price, `strike` price, & `ask` price
func ExtrinsicRisk(request *interactions.InteractionRequest) (*interactions.InteractionResponse, error) {
	p, err := getPayload(request.Data.Options)
	if err != nil {
		return nil, err
	}
	extrinsicRisk := ((p.Ask - (p.Share - p.Strike)) / p.Share) * 100
	return &interactions.InteractionResponse{
		Type: interactions.ChannelMessageWithSource,
		Data: &interactions.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("%.2f", extrinsicRisk),
		},
	}, nil
}

func getPayload(options []*interactions.ApplicationCommandInteractionDataOption) (*payload, error) {
	if len(options) != 3 {
		return nil, errors.New("missing required options")
	}
	share, err := strconv.ParseFloat(options[0].Value, 64)
	if err != nil {
		return nil, err
	}
	strike, err := strconv.ParseFloat(options[1].Value, 64)
	if err != nil {
		return nil, err
	}
	ask, err := strconv.ParseFloat(options[2].Value, 64)
	if err != nil {
		return nil, err
	}
	return &payload{Share: share, Strike: strike, Ask: ask}, nil
}
