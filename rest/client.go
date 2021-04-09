package gomeasurementsclient

import (
	"context"
	"fmt"

	"github.com/SKF/go-measurements-client/rest/models"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/uuid"

	"github.com/SKF/go-utility/v2/stages"
)

type MeasurementsClient interface {
	GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentType []string) (models.ModelNodeDataResponse, error)
}

type client struct {
	*rest.Client
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://measurement-api.iot.enlight.skf.com")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://measurement-api.%s.iot.enlight.skf.com", stage))
}

func NewClient(opts ...rest.Option) MeasurementsClient {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
		}, opts...)...,
	)

	return &client{Client: restClient}
}

func (c *client) GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentType []string) (models.ModelNodeDataResponse, error) {
	request := rest.Get("nodes/{nodeID}/node-data/recent{?content_type*}").
		Assign("nodeID", nodeID.String()).
		Assign("content_type", contentType).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	var response models.ModelNodeDataResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.ModelNodeDataResponse{}, fmt.Errorf("failed to get latest measurements: %w", err)
	}

	return response, nil
}
