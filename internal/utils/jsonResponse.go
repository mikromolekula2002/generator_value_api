package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikromolekula2002/key_generate_api/internal/models"
)

// if error exists insert your error into field "ServErr", also if error not exists keep field "ServErr" empty
func JsonResponse(c *gin.Context, ServErr string, requestId string, generatedValue string, typeGenKey string, length string) {
	switch ServErr {
	default:
		c.JSON(http.StatusSeeOther, models.JsonResponse{
			HttpCode:       http.StatusBadRequest,
			RequestId:      requestId,
			GeneratedValue: "nil",
			TypeGenValue:   "nil",
			LengthGenValue: "nil",
			ErrorMsg:       ServErr,
		})

		return
	case "":
		c.JSON(http.StatusSeeOther, models.JsonResponse{
			HttpCode:       http.StatusOK,
			RequestId:      requestId,
			GeneratedValue: generatedValue,
			TypeGenValue:   typeGenKey,
			LengthGenValue: length,
			ErrorMsg:       "nil",
		})
	}
}
