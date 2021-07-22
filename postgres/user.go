package postgres

import (
	"github.com/Thiti-Dev/scraperor/models"
	"github.com/go-pg/pg/v10"
)


type UsersRepository struct {
	DB *pg.DB
}

func (u *UsersRepository) GetUserByID(id string) (*models.User, error){
	user := new(models.User)
	err := u.DB.Model(user).Where("id = ?",id).First()
	if err != nil{
		return nil, err
	}
	return user,nil
}

func (u *UsersRepository) GetUsers() ([]*models.User,error){
	var users []*models.User
	err := u.DB.Model(&users).Select()
	if err != nil{
		return nil, err
	}
	return users,nil
}