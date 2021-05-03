package mock

import (
	"context"

	measurements "github.com/SKF/go-measurements-client/rest"
	"github.com/SKF/go-measurements-client/rest/models"
	"github.com/SKF/go-utility/v2/uuid"

	"github.com/stretchr/testify/mock"
)

type MeasurementClientMock struct {
	*mock.Mock
}

func CreateMeasurementClientMock() *MeasurementClientMock {
	client := &MeasurementClientMock{
		Mock: &mock.Mock{},
	}

	// Ensure the returned mock implements the MeasuremntsClient interface
	var _ measurements.MeasurementsClient = client

	return client
}

func (mock *MeasurementClientMock) GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentType []string) (models.ModelNodeDataResponse, error) {
	args := mock.Called(ctx, nodeID, contentType)
	return args.Get(0).(models.ModelNodeDataResponse), args.Error(1)
}

func (mock *MeasurementClientMock) PostNodeData(ctx context.Context, nodeData []models.ModelNodeDataRequest) error {
	args := mock.Called(ctx, nodeData)
	return args.Error(0)
}

func (mock *MeasurementClientMock) GetNodeData(ctx context.Context, measurementID uuid.UUID, excludeCoordinates bool) (models.ModelNodeData, error) {
	args := mock.Called(ctx, measurementID, excludeCoordinates)
	return args.Get(0).(models.ModelNodeData), args.Error(1)
}
