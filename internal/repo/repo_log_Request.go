package repo

import (
	"fmt"

	"github.com/mikromolekula2002/key_generate_api/internal/models"
	"gorm.io/gorm"
)

func (r *Repository) SaveLogRequest(userRequest *models.UsersRequests) error {
	op := "repo.SaveLogRequest"

	// Сохраняем новый запрос
	result := r.DB.Save(userRequest)
	if result.Error != nil {
		return fmt.Errorf("%s - Ошибка сохранения UserRequest: \n%v", op, result.Error)
	}

	return nil
}

func (r *Repository) CreateLogRequest(userRequest *models.UsersRequests) error {
	op := "repo.CreateLogRequest"

	// Сохраняем новый запрос
	result := r.DB.Create(userRequest)
	if result.Error != nil {
		return fmt.Errorf("%s - Ошибка сохранения UserRequest: \n%v", op, result.Error)
	}

	return nil
}

func (r *Repository) GetLogRequest(requestId string, method string) (*models.UsersRequests, bool, error) {
	op := "repo.GetLogRequest"
	var userRequest models.UsersRequests

	result := r.DB.Where("request_id = ?", requestId).Where("method = ?", method).First(&userRequest)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Если запись не найдена, возвращаем nil и сообщение об этом
			return nil, false, nil
		}
		// Возникла ошибка при выполнении запроса
		return nil, false, fmt.Errorf("%s - Ошибка получения UserRequests: \n%v", op, result.Error)
	}

	return &userRequest, true, nil
}
