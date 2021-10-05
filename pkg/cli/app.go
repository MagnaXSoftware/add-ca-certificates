package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"magnax.ca/add-ca-certificates/pkg/management"
)

const (
	CertBundlePath = "/etc/ssl/certs/ca-certificates.crt"
	LocalCertsPath = "/usr/local/share/ca-certificates/"
)

func NewApp() *cli.App {
	var certBundlePath string
	var localCertsPath = cli.NewStringSlice()

	app := &cli.App{
		Name:  "add-ca-certificates",
		Usage: "Add new certificates to the global ca-certificates file.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "bundle",
				Value:       CertBundlePath,
				Usage:       "Path to the certificate bundle",
				Destination: &certBundlePath,
			},
			&cli.StringSliceFlag{
				Name:        "local-path",
				Value:       cli.NewStringSlice(LocalCertsPath),
				Usage:       "Path to the local certificates folder. All certificates in this tree will be trusted. Can be specified multiple times.",
				Destination: localCertsPath,
			},
		},
		Action: func(c *cli.Context) error {
			manager := management.NewManager(certBundlePath, localCertsPath.Value())
			err := manager.BuildBundle()
			if err != nil {
				return err
			}
			n, err := manager.WriteBundle()
			fmt.Printf("Printed %d certificates to the bundle\n", n)
			return err
		},
	}

	return app
}
