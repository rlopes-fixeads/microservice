// Copyright 2017 The OLX Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"os"
	"gopkg.in/urfave/cli.v2"

	"microproxy/modules/server"
)

const (
	NAME string = "MicroProxy"
	VER  string = "1.0.0"
)


func main() {
	app := cli.App{
		Name:    NAME,
		Version: VER,
		Commands: []*cli.Command{
			server.CmdServer,
		},
	}

	app.Run(os.Args)
}