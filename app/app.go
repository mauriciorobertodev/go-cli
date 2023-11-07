package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// NewCli retorna uma aplicação de linha de comando pronta para ser usada
func NewCli() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Cli"
	app.Description = "Busca IPs de endereços na internet"

	app.Commands = []*cli.Command {
		{
			Name:        "ip",
			Aliases:     []string{},
			Usage:       "Busca os ips de um endereço na internet",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name: "host",
					Value: "google.com.br",
				},
			},
			Action: searchIps,
		},
		{
			Name:        "server",
			Aliases:     []string{},
			Usage:       "Busca os servidores de um endereço na internet",
			Flags: []cli.Flag {
				&cli.StringFlag{
					Name: "host",
					Value: "google.com.br",
				},
			},
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil { 
		log.Fatal(err)
		return err
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}
func searchServers(c *cli.Context) error {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil { 
		log.Fatal(err)
		return err
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

	return nil
}