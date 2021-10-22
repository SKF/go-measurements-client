package gomeasurementsclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SKF/go-measurements-client/rest/models"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/uuid"

	"github.com/SKF/go-utility/v2/stages"
)

type MeasurementsClient interface {
	GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentTypes []string, excludeCoordinates bool, limit int) (models.ModelNodeDataResponse, error)
	PostNodeData(ctx context.Context, nodeData []models.ModelNodeDataRequest) error
	PostNodeDataVerbose(ctx context.Context, nodeData []models.ModelNodeDataRequest) (models.ModelIngestNodeDataResponse, error)
	DeleteNodeData(ctx context.Context, nodeID uuid.UUID, deleteNodeDataRequest models.ModelDeleteNodeDataRequest) error

	GetMeasurement(ctx context.Context, measurementID uuid.UUID, contentType string, excludeCoordinates bool) (models.ModelMeasurementResponse, error)

	GetBandOverall(ctx context.Context, measurementID uuid.UUID, startFrequency, stopFrequency float64) (models.ModelMeasurementBandOverallResponse, error)

	GetLastCollectedAt(ctx context.Context, nodeID uuid.UUID) (*models.ModelStringResponse, error)
}

const (
	ContentTypeDataPoint       = "DATA_POINT"
	ContentTypeSpectrum        = "SPECTRUM"
	ContentTypeTimeSeries      = "TIME_SERIES"
	ContentTypeNote            = "NOTE"
	ContentTypeMissingValue    = "MISSING_VALUE"
	ContentTypeQuestionAnswers = "QUESTION_ANSWERS"
)

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

func (c *client) GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentTypes []string, excludeCoordinates bool, limit int) (models.ModelNodeDataResponse, error) {
	request := rest.Get("nodes/{nodeID}/node-data/recent{?content_type,exclude_coordinates,limit}").
		Assign("nodeID", nodeID.String()).
		Assign("content_type", contentTypes).
		Assign("exclude_coordinates", excludeCoordinates).
		Assign("limit", limit).
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

// PostNodeDataVerbose does not guarantee that the returned measurement ID
// exists. Do not use this functionality in production use cases.
func (c *client) PostNodeDataVerbose(ctx context.Context, nodeData []models.ModelNodeDataRequest) (models.ModelIngestNodeDataResponse, error) {
	request := rest.Post("/node-data{?verbose}").
		SetHeader("Accept", "application/json").
		Assign("verbose", "true").
		WithJSONPayload(nodeData)

	var response models.ModelIngestNodeDataResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.ModelIngestNodeDataResponse{}, fmt.Errorf("failed to post measurement(s): %w", err)
	}

	return response, nil
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

func (c *client) GetMeasurement(ctx context.Context, measurementID uuid.UUID, contentType string, excludeCoordinates bool) (models.ModelMeasurementResponse, error) {
	request := rest.Get("node-data/{measurementID}/{?exclude_coordinates,contentType}").
		SetHeader("Accept", "application/json").
		Assign("measurementID", measurementID.String()).
		Assign("exclude_coordinates", excludeCoordinates)

	if contentType != "" {
		request = request.Assign("contentType", contentType)
	}

	var response models.ModelMeasurementResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.ModelMeasurementResponse{}, fmt.Errorf("failed to get measurement: %w", err)
	}

	return response, nil
}

func (c *client) GetBandOverall(ctx context.Context, measurementID uuid.UUID, startFrequency, stopFrequency float64) (models.ModelMeasurementBandOverallResponse, error) {
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

func (c *client) GetLastCollectedAt(ctx context.Context, nodeID uuid.UUID) (*models.ModelStringResponse, error) {
	request := rest.Get("nodes/{nodeID}/last-collected-at").
		Assign("nodeID", nodeID.String()).
		Assign("content_type", "application/json").
		SetHeader("Accept", "application/json")

	httpResponse, err := c.Do(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to call measurements for last collected at: %w", err)
	}

	if httpResponse.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	var response models.ModelStringResponse
	if err := httpResponse.Unmarshal(response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal last collected at response: %w", err)
	}

	return &response, nil
}
