package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

// Deps defines `Client` dependencies
type Deps struct{}

// impl implements `Client` properties
type impl struct {
	deps    *Deps
	conf    *config.Config
	apiURL  string
	headers map[string]string
}

// Client interfaces `Client` methods
type Client interface {
	ListGlobalApplicationCommands() ([]*models.ApplicationCommand, error)
	CreateGlobalApplicationCommand(command *models.ApplicationCommand) error
	DeleteGlobalApplicationCommand(commandID string) error
	ListGuildApplicationCommands(guildID string) ([]*models.ApplicationCommand, error)
	CreateGuildApplicationCommand(guildID string, command *models.ApplicationCommand) error
	DeleteGuildApplicationCommand(guildID string, commandID string) error
}

// New returns a new `Client` interface
func New(deps *Deps, conf *config.Config) Client {
	return &impl{
		deps:   deps,
		conf:   conf,
		apiURL: fmt.Sprintf("%s/%s/applications/%s", conf.DiscordAPI.BaseURL, conf.DiscordAPI.APIVersion, conf.Credentials.ClientID),
		headers: map[string]string{
			"Authorization": fmt.Sprintf("Bot %s", conf.Credentials.Token),
			"Content-Type":  "application/json",
		},
	}
}

func (impl *impl) ListGlobalApplicationCommands() ([]*models.ApplicationCommand, error) {
	url := fmt.Sprintf("%s/commands", impl.apiURL)
	return impl.listApplicationCommands(url)
}

func (impl *impl) CreateGlobalApplicationCommand(command *models.ApplicationCommand) error {
	url := fmt.Sprintf("%s/commands", impl.apiURL)
	return impl.createApplicationCommand(url, command)
}

func (impl *impl) DeleteGlobalApplicationCommand(commandID string) error {
	url := fmt.Sprintf("%s/commands/%s", impl.apiURL, commandID)
	return impl.deleteApplicationCommands(url)
}

func (impl *impl) ListGuildApplicationCommands(guildID string) ([]*models.ApplicationCommand, error) {
	url := fmt.Sprintf("%s/guilds/%s/commands", impl.apiURL, guildID)
	return impl.listApplicationCommands(url)
}

func (impl *impl) CreateGuildApplicationCommand(guildID string, command *models.ApplicationCommand) error {
	url := fmt.Sprintf("%s/guilds/%s/commands", impl.apiURL, guildID)
	return impl.createApplicationCommand(url, command)
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

func (impl *impl) createApplicationCommand(url string, command *models.ApplicationCommand) error {
	body, err := marshal(command)
	if err != nil {
		return err
	}
	response, err := httpRequest(http.MethodPost, url, impl.headers, body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.New(fmt.Sprint("error - command not created: ", string(body)))
	}
	return nil
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
		return errors.New(fmt.Sprint("error - command not deleted: ", string(body)))
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

func marshal(v interface{}) (io.Reader, error) {
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
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
