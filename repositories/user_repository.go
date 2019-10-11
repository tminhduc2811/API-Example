package repositories

import (
	"exampleAPI/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

type UserRepository struct {
	session			*mgo.Session
	dbName			string
	dbCollection	string
}

func NewUserRepository(session *mgo.Session) *UserRepository {
	return &UserRepository{
		session: session,
		dbName: "user",
		dbCollection: "userInfo",
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while encoding password")
	}
	user.HashPassword = hpass
	user.Password = ""
	session := r.getSession()
	defer session.Close()
	err = session.DB(r.dbName).C(r.dbCollection).Insert(&user)
	return err
}

func (r *UserRepository) getSession() *mgo.Session {
	return r.session.Copy()
}