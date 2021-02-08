package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wafer-bw/discobottest/app/actions"
	"github.com/wafer-bw/discobottest/app/config"
	"github.com/wafer-bw/discobottest/app/errs"
	"github.com/wafer-bw/discobottest/app/interactions"
)

// Deps defines `Handler` dependencies
type Deps struct {
	Interactions interactions.Interactions
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

func (impl *impl) Handle(w http.ResponseWriter, r *http.Request) {
	response, err := impl.resolve(r)
	if err != nil {
		impl.respond(w, nil, err)
		return
	}

	body, err := impl.marshal(response)
	if err != nil {
		impl.respond(w, nil, err)
		return
	}

	impl.respond(w, body, nil)
}

func (impl *impl) resolve(r *http.Request) (*interactions.InteractionResponse, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if !impl.deps.Interactions.Verify(body, r.Header, impl.conf.Credentials.PublicKey) {
		return nil, errs.ErrUnauthorized
	}

	interaction, err := impl.unmarshal(body)
	if err != nil {
		return nil, err
	}

	switch interaction.Type {
	case interactions.Ping:
		return &interactions.InteractionResponse{Type: interactions.Pong}, nil
	case interactions.ApplicationCommand:
		return actions.Run(interaction)
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

func (impl *impl) unmarshal(data []byte) (*interactions.InteractionRequest, error) {
	interaction := &interactions.InteractionRequest{}
	if err := json.Unmarshal(data, interaction); err != nil {
		return nil, err
	}
	return interaction, nil
}

func (impl *impl) marshal(response *interactions.InteractionResponse) ([]byte, error) {
	return json.Marshal(response)
}
