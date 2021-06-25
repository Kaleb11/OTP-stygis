package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Auth/internal/constant/model"

	"Auth/internal/module/user"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// UserHandler contains the function of handler for domain user
type UserHandler interface {
	Adduser(w http.ResponseWriter, req *http.Request, p httprouter.Params)
	Validate(w http.ResponseWriter, req *http.Request, p httprouter.Params)
}

type userHandler struct {
	userCase user.Usecase
}

// UserInit is to initialize the rest handler for domain user
func UserInit(userCase user.Usecase) UserHandler {
	return &userHandler{
		userCase,
	}
}

// Test is handler testing
// func (uh *userHandler) Test(ctx *fiber.Ctx) error {
// 	return ctx.SendString("Hello, World!")
// }

func (uh *userHandler) Adduser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	body := model.Otpuser{}

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&body)
	fmt.Println(err)

	if err != nil || body.Phonenumber == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(err)
	err = uh.userCase.Adduser(w, req, p, &body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"domain":   "user",
			"action":   "create new user",
			"usecase":  "Register",
			"username": body.Phonenumber,
			"otpcode":  body.Otpcode,
		}).Errorln(err)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	w.WriteHeader(http.StatusOK)
	return
}
func (uh *userHandler) Validate(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var user model.Otpuser
	body := req.Body
	//fmt.Println("Iam the error-1", body)
	decoder := json.NewDecoder(body)
	fmt.Println(decoder)

	err := decoder.Decode(&user)
	fmt.Println("Iam the error0", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if body.Otpcode == 0 || body.Phonenumber == "" {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	isvalid, err := uh.userCase.Validate(w, req, p, user)

	fmt.Println("Iam the error2", err)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"domain":      "user",
			"action":      "create new user",
			"usecase":     "Otpcode",
			"phonenumber": user.Phonenumber,
			"otpcode":     user.Otpcode,
		}).Errorln(err)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	if !isvalid {

		http.Error(w, "Invalid otp", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Valid otp"))
	return
}
