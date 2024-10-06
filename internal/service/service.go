package service

import (
	"fmt"
	"strconv"
)

func (s *Service) GenerateValue(requestId string, length string, typeGenValue string) (string, *int, error) {
	op := "Service.GenerateValue"

	var CompleteValue string
	var err error

	leng, err := strconv.Atoi(length)
	if err != nil {

		return "", nil, fmt.Errorf("%s - Error: \n%v", op, err)
	}

	switch typeGenValue {
	case "numeric":
		CompleteValue = s.gen.NumericValue(leng)

	case "string":
		CompleteValue = s.gen.StringValue(leng)

	default:
		CompleteValue = s.gen.AlphaNumericValue(leng)
	}

	return CompleteValue, &leng, nil
}

func (s *Service) GenUniqueRequestID() (string, error) {
	op := "Service.GenUniqueRequestID"
	var requestId string

	for {
		requestId = s.gen.RequestID()
		booler, err := s.repo.CheckExistsData(requestId)
		if err != nil {
			return "", fmt.Errorf("%s - Error: \n%v", op, err)
		}

		if !booler {
			// Если получили true, выходим из цикла
			break
		}
	}
	return requestId, nil
}

func (s *Service) CheckType(typeGenValue string) string {
	switch typeGenValue {
	case "numeric":
	case "string":
	default:
		typeGenValue = "alpha_numeric"
	}

	return typeGenValue
}
