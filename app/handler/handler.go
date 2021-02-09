package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wafer-bw/udx-discord-bot/app/auth"
	"github.com/wafer-bw/udx-discord-bot/app/commands"
	"github.com/wafer-bw/udx-discord-bot/app/config"
	"github.com/wafer-bw/udx-discord-bot/app/errs"
	"github.com/wafer-bw/udx-discord-bot/app/models"
)

// Deps defines `Handler` dependencies
type Deps struct {
	Commands commands.Commands
	Auth     auth.Authorization
}

// impl implements `Handler` properties
type impl struct {
	deps *Deps
	conf *config.Config
}

// Handler interfaces `Handler` methods
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

// New returns a new `Handler` interface
func New(deps *Deps, conf *config.Config) Handler {
	return &impl{deps: deps, conf: conf}
}

var pongResponse = &models.InteractionResponse{
	Type: models.InteractionResponseTypePong,
}

func (impl *impl) Handle(w http.ResponseWriter, r *http.Request) {
	interactionRequest, err := impl.resolve(r)
	if err != nil {
		impl.respond(w, nil, err)
		return
	}

	interactionResponse, err := impl.execute(interactionRequest)
	if err != nil {
		impl.respond(w, nil, err)
		return
	}

	body, err := impl.marshal(interactionResponse)
	if err != nil {
		impl.respond(w, nil, err)
		return
	}

	impl.respond(w, body, nil)
}

func (impl *impl) resolve(r *http.Request) (*models.InteractionRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if !impl.deps.Auth.Verify(body, r.Header, impl.conf.Credentials.PublicKey) {
		return nil, errs.ErrUnauthorized
	}

	return impl.unmarshal(body)
}

func (impl *impl) execute(interaction *models.InteractionRequest) (*models.InteractionResponse, error) {
	switch interaction.Type {
	case models.InteractionTypePing:
		return pongResponse, nil
	case models.InteractionTypeApplicationCommand:
		return impl.deps.Commands.Run(interaction)
	default:
		return nil, errs.ErrInvalidInteractionType
	}
}

func (impl *impl) respond(w http.ResponseWriter, body []byte, err error) {
	switch err {
	case nil:
		if _, err = w.Write(body); err != nil {
			impl.respond(w, nil, err)
		}
	case errs.ErrUnauthorized:
		http.Error(w, err.Error(), http.StatusUnauthorized)
	case errs.ErrInvalidInteractionType:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errs.ErrNotImplemented:
		http.Error(w, err.Error(), http.StatusNotImplemented)
	default:
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (impl *impl) unmarshal(data []byte) (*models.InteractionRequest, error) {
	interaction := &models.InteractionRequest{}
	if err := json.Unmarshal(data, interaction); err != nil {
		return nil, err
	}
	return interaction, nil
}

func (impl *impl) marshal(response *models.InteractionResponse) ([]byte, error) {
	return json.Marshal(response)
}
