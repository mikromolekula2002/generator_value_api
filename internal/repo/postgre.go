package repo

import (
	"fmt"

	"github.com/mikromolekula2002/key_generate_api/internal/config"
	"github.com/mikromolekula2002/key_generate_api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Repo interface {
	CheckExistsData(requestID string) (bool, error)
	SaveGenValueData(GeneratedValue *models.GeneratedValues) error
	GetGenValueData(requestID string) (*models.GeneratedValues, error)
	SaveLogRequest(userRequest *models.UsersRequests) error
	GetLogRequest(requestId string, method string) (*models.UsersRequests, bool, error)
	CreateLogRequest(userRequest *models.UsersRequests) error
}

func InitDB(config *config.Config) (*Repository, error) {
	// Конфигурация для подключения к базе данных PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Port,
		config.Database.Sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.GeneratedValues{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.UsersRequests{})
	if err != nil {
		return nil, err
	}

	database := Repository{
		DB: db,
	}
	return &database, nil
}

func (r *Repository) CheckExistsData(requestID string) (bool, error) {
	op := "repo.IsRequestIdExists"

	var existingValue models.GeneratedValues
	err := r.DB.Where("request_id = ?", requestID).First(&existingValue).Error
	if err == nil {
		// Запись найдена
		return true, nil
	} else if err != gorm.ErrRecordNotFound {
		// Возникла ошибка при проверке (не "запись не найдена")
		return false, fmt.Errorf("%s - Ошибка проверки существующей записи: \n%v", op, err)
	}

	// Запись не найдена
	return false, nil
}

// Сохранение хеша рефреш токена
func (r *Repository) SaveGenValueData(GeneratedValue *models.GeneratedValues) error {
	op := "repo.SaveGenValueData"

	result := r.DB.Create(GeneratedValue)
	if result.Error != nil {
		return fmt.Errorf("%s - Ошибка сохранения GeneratedValue: \n%v", op, result.Error)
	}
	return nil
}

// Получение рефреш токена
func (r *Repository) GetGenValueData(requestID string) (*models.GeneratedValues, error) {
	op := "repo.GetGenValueData"
	var refreshToken models.GeneratedValues

	result := r.DB.Where("request_id = ?", requestID).First(&refreshToken)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("%s - GeneratedValue не найден: \n%v", op, result.Error)
		}
		return nil, fmt.Errorf("%s - Ошибка получения GeneratedValue: \n%v", op, result.Error)
	}

	return &refreshToken, nil
}
