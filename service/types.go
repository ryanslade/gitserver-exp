package service

import (
	"io"
	"os"
)

// BaseArgs are provided to every subcommand and control the operating environment.
type BaseArgs struct {
	Dir      string                      // The working directory of the subcommand.
	Env      []string                    // The key-value environment values, if any.
	FlowRate int64                       // The write rate for Stdout.
	Stdout   func(w io.Writer) io.Writer // The standard output stream.
}

// InitArgs configure the behavior of the Init subcommand.
type InitArgs struct {
	BaseArgs // Arguments shared by all subcommands.
}

// FetchArgs configure the behavior of the Fetch subcommand.
type FetchArgs struct {
	BaseArgs        // Arguments shared by all subcommands.
	Upstream string // The location of the repository to fetch from.
}

// MkdirAllArgs configure the behavior of the MkdirAll command.
type MkdirAllArgs struct {
	Path string      // The complete path to create.
	Mode os.FileMode // The path's mode and permission bits.
}

// A CloneRequest is provided as input to the CloneRepository service.
type CloneRequest struct {
	Upstream string // The location of the repository to fetch from.
	Local    string // The location where the repository will reside.
}
