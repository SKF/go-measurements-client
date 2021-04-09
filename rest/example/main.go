package main

import (
	"context"
	"fmt"

	measurements "github.com/SKF/go-measurements-client/rest"

	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
)

func main() {
	ctx := context.Background()

	p := myProvider{token: "<access token>"}

	client := measurements.NewClient(
		measurements.WithStage(stages.StageSandbox),
		client.WithDatadogTracing("go-measurements-client-example"),
		client.WithTokenProvider(&p),
	)

	contentTypes := []string{
		"DATA_POINT",
	}

	nodeData, err := client.GetNodeDataRecent(ctx, "c7e2fe61-4051-4029-8329-f733db081b89", contentTypes)
	if err != nil {
		panic(err)
	}

	for _, nd := range nodeData.Data {
		fmt.Printf("nd; %+v\n", nd)
	}
}

type myProvider struct {
	token string
}

func (p *myProvider) GetRawToken(ctx context.Context) (auth.RawToken, error) {
	return auth.RawToken(p.token), nil
}
