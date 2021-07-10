package telegram

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Config struct {
	Token string `yaml:"token"`
}

type Client struct {
	Bot *tb.Bot
}

func (c *Client) init(config Config) error {
	var err error
	c.Bot, err = tb.NewBot(tb.Settings{
		Token:  config.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	return err
}

func NewConnection(config Config) (*Client, error) {
	c := &Client{}
	err := c.init(config)
	return c, err
}
