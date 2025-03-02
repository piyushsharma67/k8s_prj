package http_controller

import (
	"auth_service/models"
	"auth_service/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

func (c *HTTPController) SaveUserFcmHttp(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	userId, _ := r.Context().Value("userid").(int32)

	var fcm_token models.UserFcm

	if err := json.NewDecoder(r.Body).Decode(&fcm_token); err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	validate := validator.New()

	if err := validate.Struct(fcm_token); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)
	defer cancel()

	details, err := c.Service.SaveUserFcm(ctx, &models.CreateUserFcm{
		UserId:   userId,
		FcmToken: fcm_token.FcmToken,
	}); 
	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, details)
	return
}
