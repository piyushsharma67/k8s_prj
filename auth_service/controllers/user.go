package controllers

import (
	"auth_service/models"
	"auth_service/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (c *ControllerStruct) SignupUser(w http.ResponseWriter, r *http.Request) {
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

	db_user, err := c.service.InsertUserInDB(ctx, &user)

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, db_user)
	return
}

func (c *ControllerStruct)LoginUser(w http.ResponseWriter,r *http.Request){
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

	db_user, err := c.service.GetUserByEmail(ctx,&user)

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

func (c *ControllerStruct) SaveUserFcmToken(w http.ResponseWriter, r *http.Request) {

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

	ctx = context.WithValue(ctx, "userId", userId)

	if err := c.service.InsertUserFCMInDB(ctx, fcm_token.FcmToken); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w, "Saved Successfully!!")
	return
}
