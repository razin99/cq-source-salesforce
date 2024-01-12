package client

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/simpleforce/simpleforce"
)

type Client struct {
	logger     zerolog.Logger
	Spec       Spec
	SalesForce *simpleforce.Client
}

func (c *Client) ID() string {
	return "SalesForce"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	sfc := simpleforce.NewClient(s.Endpoint, simpleforce.DefaultClientID, "v59.0")

	if err := sfc.LoginPassword(s.Username, s.Password, s.Token); err != nil {
		return Client{}, err
	}

	c := Client{
		logger:     logger,
		Spec:       *s,
		SalesForce: sfc,
	}

	return c, nil
}
