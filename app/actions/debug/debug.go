package debug

import (
	"fmt"

	"github.com/wafer-bw/discobottest/app/models"
)

// Debug is a debug command
func Debug(request *models.InteractionRequest) (*models.InteractionResponse, error) {
	fmt.Println("JOINED AT:", request.Member.JoinedAt)
	fmt.Println("NICK:", request.Member.Nick)
	// fmt.Println("USER:", request.Member.User)
	return &models.InteractionResponse{
		Type: models.InteractionResponseTypeChannelMessageWithSource,
		Data: &models.InteractionApplicationCommandCallbackData{
			Content: "received",
		},
	}, nil
}
