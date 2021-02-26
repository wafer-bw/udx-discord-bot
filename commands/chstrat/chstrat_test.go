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
// 	t.Run("success", func(t *testing.T) {
// 		request := &models.InteractionRequest{
// 			Data: &models.ApplicationCommandInteractionData{
// 				Options: []*models.ApplicationCommandInteractionDataOption{
// 					{Name: "symbol", Value: "AAPL"},
// 					{Name: "Asset Class", Value: "stocks"},
// 				},
// 			},
// 		}
// 		response, err := chstrat(request)
// 		require.Nil(t, err)
// 		require.Equal(t, "", response.Data.Content)
// 	})
// }
