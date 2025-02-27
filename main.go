package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/wiremock/go-wiremock"
	. "github.com/wiremock/wiremock-testcontainers-go"
)

const serverPort = 8080

func main() {
	ctx := context.Background()
	var emptyCustomizers []testcontainers.ContainerCustomizer
	container, err := RunContainer(ctx, emptyCustomizers...)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Panicf("failed to terminate container: %s", err)
		}
	}()

	err = container.Client.StubFor(
		wiremock.Get(wiremock.URLEqualTo("/hello")).
			WillReturnResponse(
				wiremock.NewResponse().
					WithBody(time.Now().Local().String()).
					WithHeader("Content-Type", "application/json").
					WithStatus(http.StatusOK),
			),
	)
	if err != nil {
		log.Panic(err)
	}

	uri, err := GetURI(ctx, container)
	if err != nil {
		log.Fatalf("Fail to obtain uri from wiremock container: %s", err)
	}

	res, err := http.Get(fmt.Sprintf("%s/hello", uri))
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %s\n", err)
	}
	fmt.Printf("response body: %s\n", string(resBody))
}
