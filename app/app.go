package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate return cli application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "My first app cli in go"
	app.Usage = "Search IP and names server"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "example.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP address",
			Flags:  flags,
			Action: searchIPS,
		},

		{
			Name:   "servers",
			Usage:  "Search name server",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app

}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, name := range servers {
		fmt.Println(name.Host)
	}

}

func searchIPS(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
