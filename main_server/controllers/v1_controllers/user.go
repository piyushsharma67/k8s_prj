package v1_controller

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

func (c *V1Controller) SignupUser(w http.ResponseWriter, r *http.Request) {
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

	authResponse, err := c.AuthService.Signup(ctx, &proto.SignupRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		utils.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("elapsed",time.Since(timestart))
	utils.SuccessResponse(w, &authResponse)

	defer recover()
	
	return
}
