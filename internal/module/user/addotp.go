package user

import (
	"Auth/internal/constant/model"
	"Auth/platform/encryption"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Registration is basically the flow of creating a new unverified user
func (s *service) Adduser(w http.ResponseWriter, req *http.Request, p httprouter.Params, user *model.Otpuser) error {
	generaterandnum, err := encryption.GetOtpNum()
	num, err := strconv.Atoi(generaterandnum)
	user.Otpcode = num
	if err != nil {
		return fmt.Errorf("failed to save new user %s", err.Error())
	}
	errr := s.userPersist.InsertOtpUser(user)
	if errr != nil {
		return fmt.Errorf("failed to save new user %s", err.Error())
	}
	err = json.NewEncoder(w).Encode(num)
	fmt.Println(num)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	w.WriteHeader(http.StatusOK)
	return err
}
func (s *service) Validate(w http.ResponseWriter, req *http.Request, p httprouter.Params, user model.Otpuser) (bool, error) {

	var uss model.Otpuser
	val, err := s.userPersist.Retdb().Get("id1234").Result()
	err = json.Unmarshal([]byte(val), &uss)
	if err != nil {
		fmt.Println(err)
	}
	if &uss.Otpcode == &user.Otpcode {
		return true, nil
	}
	if err != nil {
		return false, fmt.Errorf("Time expired")
	}

	return reflect.DeepEqual(&uss, &user), nil
}
