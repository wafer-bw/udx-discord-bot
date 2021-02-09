package slashcommands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/wafer-bw/discobottest/app/config"
	"github.com/wafer-bw/discobottest/app/models"
)

// Deps defines `SlashCommands` dependencies
type Deps struct{}

// impl implements `SlashCommands` properties
type impl struct {
	deps    *Deps
	conf    *config.Config
	apiURL  string
	headers map[string]string
}

// SlashCommands interfaces `SlashCommands` methods
type SlashCommands interface {
	ListGlobalApplicationCommands() ([]*models.ApplicationCommand, error)
	DeleteGlobalApplicationCommand(commandID string) error
	ListGuildApplicationCommands(guildID string) ([]*models.ApplicationCommand, error)
	DeleteGuildApplicationCommand(guildID string, commandID string) error
}

// New returns a new `SlashCommands` interface
func New(deps *Deps, conf *config.Config) SlashCommands {
	return &impl{
		deps:   deps,
		conf:   conf,
		apiURL: fmt.Sprintf("%s/%s/applications/%s", conf.DiscordAPI.BaseURL, conf.DiscordAPI.APIVersion, conf.Credentials.ClientID),
		headers: map[string]string{
			"Authorization": fmt.Sprintf("Bot %s", conf.Credentials.Token),
		},
	}
}

func (impl *impl) ListGlobalApplicationCommands() ([]*models.ApplicationCommand, error) {
	url := fmt.Sprintf("%s/commands", impl.apiURL)
	return impl.listApplicationCommands(url)
}

func (impl *impl) DeleteGlobalApplicationCommand(commandID string) error {
	url := fmt.Sprintf("%s/commands/%s", impl.apiURL, commandID)
	return impl.deleteApplicationCommands(url)
}

func (impl *impl) ListGuildApplicationCommands(guildID string) ([]*models.ApplicationCommand, error) {
	url := fmt.Sprintf("%s/guilds/%s/commands", impl.apiURL, guildID)
	return impl.listApplicationCommands(url)
}

func (impl *impl) DeleteGuildApplicationCommand(guildID string, commandID string) error {
	url := fmt.Sprintf("%s/guilds/%s/commands/%s", impl.apiURL, guildID, commandID)
	return impl.deleteApplicationCommands(url)
}

func (impl *impl) listApplicationCommands(url string) ([]*models.ApplicationCommand, error) {
	response, err := httpRequest(http.MethodGet, url, impl.headers, nil)
	if err != nil {
		return nil, err
	}
	commands := &[]*models.ApplicationCommand{}
	if err := unmarshal(response.Body, commands); err != nil {
		return nil, err
	}
	return *commands, nil
}

func (impl *impl) deleteApplicationCommands(url string) error {
	response, err := httpRequest(http.MethodDelete, url, impl.headers, nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.New(fmt.Sprint("Error: ", string(body)))
	}
	return nil
}

func unmarshal(responseBody io.ReadCloser, v interface{}) error {
	body, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}

func httpRequest(method string, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		request.Header.Set(key, val)
	}

	return client.Do(request)
}
