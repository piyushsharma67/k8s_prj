package http_controller

import (
	"auth_service/models"
	"auth_service/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (c *HTTPController) SignupHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)

	defer cancel()

	db_user, err := c.Service.InsertUserInDB(ctx, &user)

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, db_user)
	return
}

func (c *HTTPController)LoginHttp(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	var user models.User

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)

	defer cancel()

	db_user, err := c.Service.GetUserByEmail(ctx,&user)

	if err!=nil{
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	token,err:=utils.EncodeToken(user.ID)

	if err!=nil{
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	db_user.Token=token

	utils.SuccessResponse(w,db_user)
	return
}
