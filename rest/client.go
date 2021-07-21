package gomeasurementsclient

import (
	"context"
	"fmt"

	"github.com/SKF/go-measurements-client/rest/models"
	"github.com/SKF/go-measurements-client/rest/workaround"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/uuid"

	"github.com/SKF/go-utility/v2/stages"
)

type MeasurementsClient interface {
	GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentType []string) (models.ModelNodeDataResponse, error)
	PostNodeData(ctx context.Context, nodeData []models.ModelNodeDataRequest) error
	DeleteNodeData(ctx context.Context, nodeID uuid.UUID, deleteNodeDataRequest models.ModelDeleteNodeDataRequest) error

	GetBandOverall(ctx context.Context, measurementID uuid.UUID, startFrequency, stopFrequency int64) (models.ModelMeasurementBandOverallResponse, error)
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

func (c *client) GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentTypes []string) (models.ModelNodeDataResponse, error) {
	request := rest.Get("nodes/{nodeID}/node-data/recent{?content_type*}").
		Assign("nodeID", nodeID.String()).
		Assign("content_type", workaround.FormatContentTypes(contentTypes)).
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

func (c *client) DeleteNodeData(ctx context.Context, nodeID uuid.UUID, deleteNodeDataRequest models.ModelDeleteNodeDataRequest) error {
	request := rest.Delete("nodes/{nodeID}/node-data").
		Assign("nodeID", nodeID.String()).
		SetHeader("Accept", "application/json").
		WithJSONPayload(deleteNodeDataRequest)

	if _, err := c.Do(ctx, request); err != nil {
		return fmt.Errorf("failed to delete measurement(s): %w", err)
	}

	return nil
}

func (c *client) GetBandOverall(ctx context.Context, measurementID uuid.UUID, startFrequency, stopFrequency int64) (models.ModelMeasurementBandOverallResponse, error) {
	request := rest.Get("/node-data/{measurementId}/band/overall{?startFrequency,stopFrequency*}").
		Assign("measurementId", measurementID.String()).
		Assign("startFrequency", startFrequency).
		Assign("stopFrequency", stopFrequency).
		SetHeader("Accept", "application/json")

	var response models.ModelMeasurementBandOverallResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.ModelMeasurementBandOverallResponse{}, fmt.Errorf("failed to get overall band value for measurement: %w", err)
	}

	return response, nil
}
