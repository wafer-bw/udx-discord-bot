package chstrat

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	exitCode := m.Run()
	os.Exit(exitCode)
}

// todo use mocks so it doesnt spam the api
// func TestChstrat(t *testing.T) {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		log.Println("Warning: could not load .env file")
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		request := &discord.InteractionRequest{
// 			Data: &discord.ApplicationCommandInteractionData{
// 				Options: []*discord.ApplicationCommandInteractionDataOption{
// 					{Name: "symbol", Value: "AAPL"},
// 				},
// 			},
// 		}
// 		response, err := chstrat(request)
// 		require.Nil(t, err)
// 		require.Equal(t, "", response.Data.Content)
// 	})
// }
