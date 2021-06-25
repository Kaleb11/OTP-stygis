package persistence

//go:generate mockgen -destination=../../../mocks/user/persistence_mock.go -package=user_mock -source=user.go

import ( //	"context"
	//	"github.com/iDevoid/cptx"
	"Auth/internal/constant/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	//"Auth/platform/gorm"
	//	"context"
	//	"github.com/iDevoid/cptx"
)

// UserPersistence contains the list of functions for database table users
type UserPersistence interface {
	InsertOtpUser(user *model.Otpuser) error
	Retdb() *redis.Client
}

type userPersistence struct {
	db *redis.Client
}

// UserInit is to init the user persistence that contains data accounts
func UserInit(db *redis.Client) UserPersistence {
	return &userPersistence{
		db,
	}
}

// NewUser retrun a pointer to a User
// func NewUser() *model.User {
// 	return new(model.User)
// }

// InsertUser is the input the data record to database table users
func (up *userPersistence) InsertOtpUser(user *model.Otpuser) error {
	json, err := json.Marshal(model.Otpuser{Phonenumber: user.Phonenumber, Otpcode: user.Otpcode})

	if err != nil {
		fmt.Println(err)
	}
	usr := up.db.Set("id1234", json, 1*time.Minute).Err()
	//user.Id = usr.
	return usr
	//return up.db.Main().QueryRowMustTx(ctx, query.UserInsert, params, &user.ID)
}

func (up *userPersistence) Retdb() *redis.Client {
	return up.db
}
