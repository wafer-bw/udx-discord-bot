package stocks

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/wafer-bw/discobottest/app/models"
)

type payload struct {
	Share  float64
	Strike float64
	Ask    float64
}

// ExtrinsicRisk calculates the extrinsic risk of an option
// provided the `share` price, `strike` price, & `ask` price
func ExtrinsicRisk(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	p, err := getPayload(request.Data.Options)
	if err != nil {
		fmt.Println(err)
		return responseErr("Error parsing command :("), nil
	}
	risk := calcExtrinsicRisk(p)
	return responseOk(risk), nil
}

func calcExtrinsicRisk(p *payload) float64 {
	return ((p.Ask - (p.Share - p.Strike)) / p.Share) * 100
}

func responseOk(risk float64) *models.InteractionResponse {
	return &models.InteractionResponse{
		Type: models.InteractionResponseTypes.ChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("%.2f%%", risk),
		},
	}
}

func responseErr(msg string) *models.InteractionResponse {
	return &models.InteractionResponse{
		Type: models.InteractionResponseTypes.ChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{Content: msg},
	}
}

func getPayload(options []*models.ApplicationCommandInteractionDataOption) (*payload, error) {
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
