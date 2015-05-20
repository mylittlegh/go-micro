package cmd

import (
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/myodc/go-micro/broker"
	"github.com/myodc/go-micro/registry"
	"github.com/myodc/go-micro/server"
	"github.com/myodc/go-micro/store"
	"github.com/myodc/go-micro/transport"
)

var (
	Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server_address",
			EnvVar: "MICRO_SERVER_ADDRESS",
			Value:  ":0",
			Usage:  "Bind address for the server. 127.0.0.1:8080",
		},
		cli.StringFlag{
			Name:   "broker",
			EnvVar: "MICRO_BROKER",
			Value:  "http",
			Usage:  "Broker for pub/sub. http, nats, etc",
		},
		cli.StringFlag{
			Name:   "broker_address",
			EnvVar: "MICRO_BROKER_ADDRESS",
			Usage:  "Comma-separated list of broker addresses",
		},
		cli.StringFlag{
			Name:   "registry",
			EnvVar: "MICRO_REGISTRY",
			Value:  "consul",
			Usage:  "Registry for discovery. kubernetes, consul, etc",
		},
		cli.StringFlag{
			Name:   "registry_address",
			EnvVar: "MICRO_REGISTRY_ADDRESS",
			Usage:  "Comma-separated list of registry addresses",
		},
		cli.StringFlag{
			Name:   "store",
			EnvVar: "MICRO_STORE",
			Value:  "consul",
			Usage:  "Store used as a basic key/value store using consul, memcached, etc",
		},
		cli.StringFlag{
			Name:   "store_address",
			EnvVar: "MICRO_STORE_ADDRESS",
			Usage:  "Comma-separated list of store addresses",
		},
		cli.StringFlag{
			Name:   "transport",
			EnvVar: "MICRO_TRANSPORT",
			Value:  "http",
			Usage:  "Transport mechanism used; http, rabbitmq, etc",
		},
		cli.StringFlag{
			Name:   "transport_address",
			EnvVar: "MICRO_TRANSPORT_ADDRESS",
			Usage:  "Comma-separated list of transport addresses",
		},
	}
)

func Setup(c *cli.Context) error {
	server.Address = c.String("server_address")

	bAddrs := strings.Split(c.String("broker_address"), ",")

	switch c.String("broker") {
	case "http":
		broker.DefaultBroker = broker.NewHttpBroker(bAddrs)
	case "nats":
		broker.DefaultBroker = broker.NewNatsBroker(bAddrs)
	}

	rAddrs := strings.Split(c.String("registry_address"), ",")

	switch c.String("registry") {
	case "kubernetes":
		registry.DefaultRegistry = registry.NewKubernetesRegistry(rAddrs)
	case "consul":
		registry.DefaultRegistry = registry.NewConsulRegistry(rAddrs)
	}

	sAddrs := strings.Split(c.String("store_address"), ",")

	switch c.String("store") {
	case "memcached":
		store.DefaultStore = store.NewMemcacheStore(sAddrs)
	case "memory":
		store.DefaultStore = store.NewMemoryStore(sAddrs)
	case "etcd":
		store.DefaultStore = store.NewEtcdStore(sAddrs)
	}

	tAddrs := strings.Split(c.String("transport_address"), ",")

	switch c.String("transport") {
	case "http":
		transport.DefaultTransport = transport.NewHttpTransport(tAddrs)
	case "rabbitmq":
		transport.DefaultTransport = transport.NewRabbitMQTransport(tAddrs)
	case "nats":
		transport.DefaultTransport = transport.NewNatsTransport(tAddrs)
	}

	return nil
}

func Init() {
	cli.AppHelpTemplate = `
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`

	cli.HelpPrinter = func(templ string, data interface{}) {
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		t := template.Must(template.New("help").Parse(templ))
		err := t.Execute(w, data)
		if err != nil {
			panic(err)
		}
		w.Flush()
		os.Exit(2)
	}

	app := cli.NewApp()
	app.HideVersion = true
	app.Usage = "a go micro app"
	app.Action = func(c *cli.Context) {}
	app.Before = Setup
	app.Flags = Flags
	app.RunAndExitOnError()
}