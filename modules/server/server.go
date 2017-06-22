// Copyright 2017 The OLX Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package server

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"
)

var CmdServer = &cli.Command{
	Name:    "server",
	Aliases: []string{"s"},
	Usage:   "Run server",
	Action:  runServer,
}

func runServer(c *cli.Context) error {
	fmt.Println("Server running")
	return nil
}
