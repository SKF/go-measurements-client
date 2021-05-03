package gomeasurementsclient

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SKF/go-measurements-client/rest/models"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

type MeasurementsClient interface {
	GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentType []string) (models.ModelNodeDataResponse, error)
	PostNodeData(ctx context.Context, nodeData []models.ModelNodeDataRequest) error
	GetNodeData(ctx context.Context, measurementID uuid.UUID, excludeCoordinates bool) (models.ModelNodeData, error)
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
		SetHeader("Accept", "application/json")

	var response models.ModelNodeDataResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.ModelNodeDataResponse{}, fmt.Errorf("failed to get latest measurements: %w", err)
	}

	return response, nil
}

func (c *client) PostNodeData(ctx context.Context, nodeData []models.ModelNodeDataRequest) error {
	request := rest.Post("/node-data").
		SetHeader("Accept", "application/json").
		WithJSONPayload(nodeData)

	if _, err := c.Do(ctx, request); err != nil {
		return fmt.Errorf("failed to post measurement(s): %w", err)
	}

	return nil
}

func (c *client) GetNodeData(ctx context.Context, measurementID uuid.UUID, excludeCoordinates bool) (models.ModelNodeData, error) {
	request := rest.Get("/node-data/{measurementID}/{?exclude_coordinates*}").
		Assign("measurementID", measurementID.String()).
		Assign("exclude_coordinates", strconv.FormatBool(excludeCoordinates)).
		SetHeader("Accept", "application/json")

	var response models.ModelNodeData
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.ModelNodeData{}, fmt.Errorf("failed to get measurement: %w", err)
	}

	return response, nil
}
