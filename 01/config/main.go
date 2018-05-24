package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Client struct {
	consulIP   string
	connString string
}

func (c *Client) String() string {
	return fmt.Sprintf("ConsulIP: %s, connection String: %s",
		c.consulIP, c.connString)
}

var defaultClient = Client{
	consulIP:   "localhost:9000",
	connString: "postgres://localhost:5432",
}

type ConfigFunc func(opt *Client)

func FromFile(path string) ConfigFunc {
	return func(opt *Client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		decoder := json.NewDecoder(f)

		fop := struct {
			ConsulIP string `json:"consul_ip"`
		}{}

		err = decoder.Decode(&fop)

		if err != nil {
			panic(err)
		}

		opt.consulIP = fop.ConsulIP
	}
}

func FromEnv() ConfigFunc {
	return func(opt *Client) {
		conn, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connString = conn
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client {
	client := defaultClient
	for _, fn := range opts {
		fn(&client)
	}
	return &client
}

func main() {
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}
