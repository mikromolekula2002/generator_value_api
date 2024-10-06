package service

import (
	"fmt"
	"time"

	"github.com/mikromolekula2002/key_generate_api/internal/generator"
	"github.com/mikromolekula2002/key_generate_api/internal/models"
	"github.com/mikromolekula2002/key_generate_api/internal/repo"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repo repo.Repo
	Log  *logrus.Logger
	gen  generator.Generate
}

func InitService(logger *logrus.Logger, repo repo.Repo, gen generator.Generate) *Service {
	return &Service{
		Log:  logger,
		repo: repo,
		gen:  gen,
	}
}

func (s *Service) TransportGenValue(requestId string, length string, typeGenValue string, booler bool) (ReqId string, Genvalue string, Err error) {
	op := "Service.TransportGenValue"

	value := &models.GeneratedValues{}

	switch booler {
	case true:
		var err error
		value, err = s.repo.GetGenValueData(requestId)
		if err != nil {
			return "", "", fmt.Errorf("%s - Error: \n%v", op, err)
		}

		return requestId, value.RandomValue, nil

	case false:
		GenValue, leng, err := s.GenerateValue(requestId, length, typeGenValue)
		if err != nil {
			return "", "", fmt.Errorf("%s - Error: \n%v", op, err)
		}

		value = &models.GeneratedValues{
			RequestId:   requestId,
			RandomValue: GenValue,
			ValueType:   typeGenValue,
			Length:      *leng,
			CreatedAt:   time.Now(),
		}

		err = s.repo.SaveGenValueData(value)
		if err != nil {
			return "", "", fmt.Errorf("%s - Error: \n%v", op, err)
		}

		return requestId, GenValue, nil
	}

	return requestId, value.RandomValue, nil
}

func (s *Service) GenValueWithoutID(length string, typeGenValue string) (requestId string, value string, Err error) {
	op := "Service.GetValueWithoutID"

	typeGenValue = s.CheckType(typeGenValue)

	requestId, err := s.GenUniqueRequestID()
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return "", "", fmt.Errorf("%s - Error: \n%v", op, err)
	}

	requestId, value, err = s.TransportGenValue(requestId, length, typeGenValue, false)
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return "", "", fmt.Errorf("%s - Error: \n%v", op, err)
	}

	return requestId, value, nil
}

func (s *Service) GenValueWithID(requestId string, length string, typeGenValue string) (value string, Err error) {
	op := "Service.GetValueWithID"

	typeGenValue = s.CheckType(typeGenValue)

	booler, err := s.repo.CheckExistsData(requestId)
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return "", fmt.Errorf("%s - Error: \n%v", op, err)
	}

	_, value, err = s.TransportGenValue(requestId, length, typeGenValue, booler)
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return "", fmt.Errorf("%s - Error: \n%v", op, err)
	}

	return value, nil
}

func (s *Service) GetValue(requestId string) (*models.GeneratedValues, error) {
	op := "Service.GetValue"

	genValue, err := s.repo.GetGenValueData(requestId)
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return nil, fmt.Errorf("%s - Error: \n%v", op, err)
	}

	return genValue, nil
}
