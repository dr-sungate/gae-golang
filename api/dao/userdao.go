package dao

import (
	"net/http"
	"time"
	"log"
	"gaeapp/gae-golang/api/client"
	"gaeapp/gae-golang/api/entity"

)

type UserDAO struct {
	User *entity.GaeUsers
	Users *[]entity.GaeUsers
	datastore *client.GaeDatastoreInstance
}

// インスタンス作成用のメソッド
func UserDAONew(r *http.Request) *UserDAO {

	I := &UserDAO{
		User: &entity.GaeUsers{},
		Users: &[]entity.GaeUsers{},
		datastore: client.NewGaeDatastore(r),
	}
	return I
}

func (dao *UserDAO) GetAll(name string) error{
	if _, err := dao.datastore.GetAll(entity.GAEUSERS_ENTITYNAME, "Name =", name, dao.Users); err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) Get(id int64) error{
	dao.User.Id = id
	if err := dao.datastore.GetOne(dao.User); err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) Create() error{
	log.Print("###########")	
	log.Print(dao.User)
	log.Print("######2#####")	

	if err := dao.datastore.Put(dao.User); err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) SetData(name, email string, newflag bool) {
	dao.User.Name = name
	dao.User.Email = email
	if newflag {
		dao.User.CreateDate = time.Now()
		dao.User.UpdateDate = time.Now()
	}
	log.Print(dao.User)
}