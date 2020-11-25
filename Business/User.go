package Business

import (
	"Book/Database"
	"Book/Database/DAL"
	"Book/Database/Models"
	"errors"
	guuid "github.com/google/uuid"
	"log"
	"strings"
)

type userBusiness struct {
	userRepository Models.UserRepository
}

func (u userBusiness) UpdateStatus(id string, status bool) error {
	err := u.userRepository.UpdateStatus(id, status)
	if err != nil {
		Database.ErrLog.Print(err)
		return  err
	}
	return err
}

func (u userBusiness) Create(user Models.User) (*Models.User, error) {
	code, err := u.GenCodeOrder()
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	user.Code = code
	result, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (u userBusiness) GetAll() ([]*Models.User, error) {
	result, err := u.userRepository.Gets()
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return result, err
}

func (u userBusiness) Get(id string) (*Models.User, error) {
	result, err := u.userRepository.Get(id)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return result, err
}

func (u userBusiness) Update(id string, user Models.User) (*Models.User, error) {
	result, err := u.userRepository.Update(id,user)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return result, err
}

type UserBusiness interface {
	Create(user Models.User) (*Models.User, error)
	GetAll() ([]*Models.User, error)
	Get(id string)( *Models.User, error)
	Update(id string, user Models.User) (*Models.User, error)
	UpdateStatus(id string, status bool) error
	GenCodeOrder() (code string, err error)
}

func (u userBusiness) GenCodeOrder() (code string, err error) {
	id := guuid.New().String()[0:8]
	code = strings.ToUpper(id)
	log.Print(code)
	if code == "" {
		return "", errors.New("dont generate code")
	}
	return code, err
}

func NewUserBusiness() UserBusiness {
	userRepository := DAL.NewUserRepository(Database.Db)
	return userBusiness{userRepository: userRepository}
}
