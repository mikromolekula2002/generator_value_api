package service

import (
	"fmt"

	"github.com/mikromolekula2002/key_generate_api/internal/models"
)

func (s *Service) SaveUserRequest(userAgent, method, url, requestId, genValue string) error {
	op := "Service.SaveUserRequest"
	var saveRequest *models.UsersRequests

	userRequest, Exists, err := s.repo.GetLogRequest(requestId, method)
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return fmt.Errorf("%s - Error: \n%v", op, err)
	}
	if Exists {

		userRequest.RequestCount = userRequest.RequestCount + 1
		saveRequest = userRequest

		err = s.repo.SaveLogRequest(saveRequest)
		if err != nil {
			s.Log.Errorf("%s - Error: \n%v", op, err)
			return fmt.Errorf("%s - Error: \n%v", op, err)
		}
	} else {

		saveRequest = &models.UsersRequests{
			RequestId:    requestId,
			RandomValue:  genValue,
			UserAgent:    userAgent,
			Url:          url,
			Method:       method,
			RequestCount: 1,
		}

		err = s.repo.CreateLogRequest(saveRequest)
		if err != nil {
			s.Log.Errorf("%s - Error: \n%v", op, err)
			return fmt.Errorf("%s - Error: \n%v", op, err)
		}

	}

	return nil
}

func (s *Service) GetUserRequest(requestId, method string) (*models.UsersRequests, error) {
	op := "Service.SaveUserRequest"

	userRequest, Exists, err := s.repo.GetLogRequest(requestId, method)
	if err != nil {
		s.Log.Errorf("%s - Error: \n%v", op, err)
		return nil, fmt.Errorf("%s - Error: \n%v", op, err)
	}
	if Exists {
		fmt.Println("Лог запроса юзера:", userRequest)
		return userRequest, nil
	} else {
		fmt.Println("Лога запроса юзера не найдено.")
		return nil, nil
	}

}
