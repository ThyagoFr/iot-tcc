package usesCases

import (
	"context"
	"time"

	model2 "github.com/thyago/tcc/api-service/domain/model"
	"github.com/thyago/tcc/api-service/internal/device"
	"github.com/thyago/tcc/api-service/internal/measurement"
)

type Servicer interface {
	Home() ([]*model2.Measurement, error)
	Historical(deviceID int, to time.Time, from time.Time) ([]*model2.Measurement, error)
}

type Service struct {
	deviceRepository device.Repository
	measurement      measurement.Repository
}

func NewService(deviceRepository device.Repository, measurement measurement.Repository) Servicer {
	return &Service{deviceRepository: deviceRepository, measurement: measurement}
}

func (s *Service) Home() ([]*model2.Measurement, error) {
	ctx := context.Background()
	devices, _ := s.deviceRepository.GetAll(ctx)
	lastMeasurements := make([]*model2.Measurement, 0)
	for _, d := range devices {
		measure, err := s.measurement.LastMeasurementsForDevice(ctx, d.DeviceID)
		if err != nil {
			return nil, err
		}
		lastMeasurements = append(lastMeasurements, measure)
	}
	return lastMeasurements, nil
}

func (s *Service) Historical(deviceID int, to time.Time, from time.Time) ([]*model2.Measurement, error) {
	ctx := context.Background()
	return s.measurement.Historical(ctx, deviceID, to, from)
}
