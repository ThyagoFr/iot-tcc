package service

import (
	"context"
	"github.com/thyago/tcc/subscriber/internal/device"
	"github.com/thyago/tcc/subscriber/internal/measurement"
	"time"
)

type Servicer interface {
	Home() ([]*measurement.Measurement,error)
	Historical(deviceID int, to time.Time, from time.Time) ([]*measurement.Measurement, error)
}

type Service struct {
	deviceRepository device.Repository
	measurement measurement.Repository
}

func NewService(deviceRepository device.Repository, measurement measurement.Repository) Servicer {
	return &Service{deviceRepository: deviceRepository, measurement: measurement}
}

func (s *Service) Home() ([]*measurement.Measurement,error){
	ctx := context.Background()
	devices, _ := s.deviceRepository.GetAll(ctx)
	lastMeasurements := make([]*measurement.Measurement,0)
	for _, d := range devices {
		measure, err := s.measurement.LastMeasurementsForDevice(ctx,d.DeviceID)
		if err != nil {
			return nil,err
		}
		lastMeasurements = append(lastMeasurements, measure)
	}
	return lastMeasurements,nil
}

func (s *Service) Historical(deviceID int, to time.Time, from time.Time) ([]*measurement.Measurement, error) {
	ctx := context.Background()
	return s.measurement.Historical(ctx,deviceID,to,from)
}
