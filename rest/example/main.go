package main

import (
	"context"
	"fmt"

	hierarchy "github.com/SKF/go-hierarchy-client/rest"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
)

const serviceName = "user_service"

func main() {
	ctx := context.Background()

	p := myProvider{token: "eyJraWQiOiJ3TktkUUtMQURMdmVoRzR5V2h0RjRHZSsyUW9Rdm1DXC9vRzdFWkU0cVI2ND0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiI0Nzg0MzBiMC1hODM2LTQyZjktOWFjNS0xNDI4YzhkNmJhODUiLCJldmVudF9pZCI6IjU0NjQzZmFlLTdjZGItNGQxMS1hMGQ5LWU5ODc1ZDE2MGRiYiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE2MTc4NjA5ODAsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC5ldS13ZXN0LTEuYW1hem9uYXdzLmNvbVwvZXUtd2VzdC0xX0RnN3BURGpXYyIsImV4cCI6MTYxNzk1MTM5NywiaWF0IjoxNjE3OTQ3Nzk3LCJqdGkiOiIyY2ExNThhMy05NDM5LTQxZDctOTFmZi1lZDZkMGZmOThmYmEiLCJjbGllbnRfaWQiOiIxOThhaXEwdXBwajBtbzZjMjh1bHZ1aWYxYiIsInVzZXJuYW1lIjoicmlja2FyZC5lbmdsdW5kQHNrZi5jb20ifQ.T1DDw9lLv4ccm-4U85laEj8bh05_FR0OJF8_8gqclZ1PVM-L-x0rXy4-JiV8jthgFnj0gba4PsDh2Yv56TKLl23VtzXhH-W6gzfc21DV4SAcYGvA1eXjzIgiAuwkr8OaNIZHdZQ3M6zF5aiRPsg-rWAOpDVzc0ixm0nbAddzS6hIud5jTJCUDCmY0NEGXhnFGeQ2gVcNPpzSxgl230EC9oWrN8aG6fy_OzjXoIBPeVSu3UqBTtzrYTddQps0cgrVzS9FmDQeDTCL9-lbg9hc1AyuuupyowM9u7JUfgMzkvIfdy1cgMQLMLF_KhTEHQUzWMe6lZaYFDiY0MvpgF-Y1w"}
	client := hierarchy.NewClient(
		hierarchy.WithStage(stages.StageSandbox),
		client.WithDatadogTracing(serviceName),
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
