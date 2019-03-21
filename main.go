package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "k8sdaysindia"
	app.Usage = "Using client-go effectively with Kubernetes api"
	app.Version = "0.2"

	app.Commands = []cli.Command{
		{Name: "crud", Usage: "Run CRUD example", Action: crudOperation},
		{Name: "lister", Usage: "Run lister example", Action: lister},
		{Name: "informer", Usage: "Run informer example", Action: informer},
		{Name: "workqueue", Usage: "Run workqueue example", Action: workqueue_example},
	}

	app.Run(os.Args)
}
