package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"main_server/models"
	"main_server/proto"
	"main_server/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func (c *ControllerStruct)SaveUserFcm(w http.ResponseWriter,r *http.Request){
	timestart:=time.Now()
	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	defer func(){
		if re := recover(); re != nil {
			// Handle the panic
			utils.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
			// Optionally log the panic
			fmt.Println("Recovered from panic:", r)
		}
	}()

	var fcm models.UserFcm

	if err := json.NewDecoder(r.Body).Decode(&fcm); err != nil {
		utils.ErrorResponse(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	validate := validator.New()

	if err := validate.Struct(fcm); err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.ApiTimeoutTime)
	defer cancel()

	userId, _ := r.Context().Value("userid").(int32)

	details,err:=c.authService.SaveFcmToken(ctx,&proto.SaveUserFcmRequest{
		FcmToken: fcm.FcmToken,
		UserId: userId,
	})

	if err!=nil{
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(w,models.ConvertToLowercaseRequest(details))

	fmt.Println("elapsed",time.Since(timestart))
}