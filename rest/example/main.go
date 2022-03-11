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

	accessToken := "<access token>"

	measurementsClient := measurements.NewClient(
		measurements.WithStage(stages.StageSandbox),
		client.WithDatadogTracing("go-measurements-client-example"),
		client.WithTokenProvider(auth.RawToken(accessToken)),
	)

	contentTypes := []string{
		"DATA_POINT",
	}

	nodeData, err := measurementsClient.GetNodeDataRecent(ctx, "c7e2fe61-4051-4029-8329-f733db081b89", contentTypes, true, 0)
	if err != nil {
		panic(err)
	}

	for _, nd := range nodeData.Data {
		fmt.Printf("nd; %+v\n", nd)
	}
}
