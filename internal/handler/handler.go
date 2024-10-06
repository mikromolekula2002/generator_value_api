package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	myError "github.com/mikromolekula2002/key_generate_api/internal/errors"
	"github.com/mikromolekula2002/key_generate_api/internal/service"
	"github.com/mikromolekula2002/key_generate_api/internal/utils"
)

const (
	jsonGenValue = "/api/retrieve/?request_id="
)

// Структура сервера с фреймворком и сервисным слоем
type Handler struct {
	Gin *gin.Engine
	Svc *service.Service
}

// Инициализация сервера
func Init(Svc *service.Service) *Handler {
	router := gin.New()

	handler := &Handler{
		Gin: router,
		Svc: Svc,
	}

	handler.Routes()

	return handler
}

// Маршруты сервера
func (h *Handler) Routes() {
	// Маршрут для проверки работы сервера
	h.Gin.POST("/api/generate/", h.GenerateKey)
	h.Gin.GET("/api/retrieve/", h.RetrieveKey)
}

func (h *Handler) GenerateKey(c *gin.Context) {
	method, userAgent, url := utils.RetrieveUserRequest(c)
	var value string

	typeGenKey := c.Query("type")
	if typeGenKey == "" {
		typeGenKey = "string"
	}

	length := c.Query("length")
	if length == "" {
		length = "32"
	}

	requestId := c.Query("request_id")
	if requestId == "" {
		var err error

		requestId, value, err = h.Svc.GenValueWithoutID(length, typeGenKey)
		if err != nil {
			utils.JsonResponse(c, myError.ErrServer.Error(), requestId, "nil", typeGenKey, length)
		}
	} else {
		var err error

		value, err = h.Svc.GenValueWithID(requestId, length, typeGenKey)
		if err != nil {
			utils.JsonResponse(c, myError.ErrServer.Error(), requestId, "nil", typeGenKey, length)
		}
	}

	genRequestUrl := jsonGenValue + requestId
	utils.JsonResponse(c, "", requestId, genRequestUrl, typeGenKey, length)

	err := h.Svc.SaveUserRequest(userAgent, method, url, requestId, value)
	if err != nil {
		h.Svc.Log.Error("handler.RetrieveKey - ошибка сохранения логов")
	}

}

func (h *Handler) RetrieveKey(c *gin.Context) {
	method, userAgent, url := utils.RetrieveUserRequest(c)

	requestId := c.Query("request_id")
	if requestId == "" {
		h.Svc.Log.Error("handler.RetrieveKey - Error: Invalid request Id")
		utils.JsonResponse(c, myError.ErrMissingRequestID.Error(), requestId, "nil", "nil", "nil")
		return
	}

	Value, err := h.Svc.GetValue(requestId)
	if err != nil {
		utils.JsonResponse(c, myError.ErrServer.Error(), requestId, "nil", "nil", "nil")
		return
	}

	length := strconv.Itoa(Value.Length)
	utils.JsonResponse(c, "", requestId, Value.RandomValue, Value.ValueType, length)

	err = h.Svc.SaveUserRequest(userAgent, method, url, requestId, Value.RandomValue)
	if err != nil {
		h.Svc.Log.Error("handler.RetrieveKey - ошибка сохранения логов")
	}

}
