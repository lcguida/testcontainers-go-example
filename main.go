package main

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func main () {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "fsouza/fake-gcs-server",
		Cmd: []string{"-scheme http", "-port 4443"},
		ExposedPorts: []string{"4443/tcp"},
		WaitingFor: wait.ForLog("started"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
	})
	if err != nil {
		panic(err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(host)
}