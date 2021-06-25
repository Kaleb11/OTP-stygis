package user

//go:generate mockgen -destination=../../../mocks/user/usecase_mock.go -package=user_mock -source=initiator.go

import (
	"Auth/internal/constant/model"
	"net/http"

	"Auth/internal/storage/persistence"

	"github.com/julienschmidt/httprouter"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	Adduser(w http.ResponseWriter, req *http.Request, p httprouter.Params, user *model.Otpuser) error
	Validate(w http.ResponseWriter, req *http.Request, p httprouter.Params, user model.Otpuser) (bool, error)
}

type service struct {
	userPersist persistence.UserPersistence
}

// Initialize takes all necessary service for domain user to run the business logic of domain user
func Initialize(

	userPersist persistence.UserPersistence,

) Usecase {
	return &service{
		userPersist,
	}
}
