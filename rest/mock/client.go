package mock

import (
	"context"

	rest "github.com/SKF/go-measurements-client/rest"
	"github.com/SKF/go-measurements-client/rest/models"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/stretchr/testify/mock"
)

type MeasurementsClientMock struct {
	*mock.Mock
}

func NewMeasurementsClient() *MeasurementsClientMock {
	client := &MeasurementsClientMock{
		Mock: &mock.Mock{},
	}

	// Ensure the returned mock implements the MeasurementsClient interface
	var _ rest.MeasurementsClient = client

	return client
}

func (c *MeasurementsClientMock) GetNodeDataRecent(ctx context.Context, nodeID uuid.UUID, contentTypes []string) (models.ModelNodeDataResponse, error) {
	args := c.Called(ctx, nodeID, contentTypes)
	return args.Get(0).(models.ModelNodeDataResponse), args.Error(1)
}
