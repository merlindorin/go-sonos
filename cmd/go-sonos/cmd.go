package main

import (
	"github.com/alecthomas/kong"
	c "github.com/hoomy-official/go-shared/pkg/cmd"
	"github.com/hoomy-official/go-sonos/cmd/go-sonos/generate"
	"github.com/joho/godotenv"
)

const (
	name        = "sonos-api-generator"
	description = "A Sonos API generator based on @svrooij work: https://sonos.svrooij.io/"
)

//nolint:gochecknoglobals //used for building the app
var (
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

type CMD struct {
	*c.Commons

	Generate *generate.Generate `cmd:"generate"`
}

func main() {
	_ = godotenv.Load()

	cli := CMD{
		Commons: &c.Commons{
			Version: c.NewVersion(name, version, commit, buildSource, date),
			Licence: c.NewLicence(license),
		},
		Generate: &generate.Generate{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name(name),
		kong.Description(description),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Commons))
}
