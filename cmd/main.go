package main

import (
	"context"
	"log"

	"github.com/flying-robot/gitserver/adapter/filesystem"
	"github.com/flying-robot/gitserver/adapter/git"
	"github.com/flying-robot/gitserver/service"
)

func main() {
	request := service.CloneRequest{
		Upstream: "https://github.com/flying-robot/commit-sink.git",
		Local:    "/tmp/commit-sink.git",
	}

	log.Println("cloning with default adapters")
	cloneRepoService := &service.CloneRepositoryService{
		Filesystem: &filesystem.Adapter{},
		Git:        &git.Adapter{},
	}
	cloneRepoService.Clone(context.Background(), request)

	log.Println("cloning with tracing adapters")
	cloneRepoService.Git = &git.TracingAdapter{
		Trace:            true,
		TracePackAccess:  true,
		TracePacket:      true,
		TracePerformance: true,
		TraceSetup:       true,
	}
	cloneRepoService.Clone(context.Background(), request)
}
