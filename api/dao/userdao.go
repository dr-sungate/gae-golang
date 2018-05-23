package dao

import (
	"net/url"
	"net/http"
	"time"
	"strconv"
	"gaeapp/gae-golang/api/client"
	"gaeapp/gae-golang/api/entity"
	log "gaeapp/gae-golang/api/logger"
)

type UserDAO struct {
	User *entity.GaeUsers
	Users *[]entity.GaeUsers
	BaseDAO
}

// インスタンス作成用のメソッド
func UserDAONew(r *http.Request) *UserDAO {

	I := &UserDAO{
		User: &entity.GaeUsers{},
		Users: &[]entity.GaeUsers{},
		BaseDAO: BaseDAO{datastore: client.NewGaeDatastore(r)},
	}
	return I
}

func (dao *UserDAO) GetAll(query url.Values) error{
	log.Info(query)
	limit, _ := strconv.Atoi(query.Get(QUERY_LIMIT))
	query.Del(QUERY_LIMIT)
	offset, _ := strconv.Atoi(query.Get(QUERY_OFFSET))
	query.Del(QUERY_OFFSET)
	order := query.Get(QUERY_ORDER)
	query.Del(QUERY_ORDER)
	return dao.BaseDAO.GetAll(entity.GAEUSERS_ENTITYNAME, query, limit, offset, order, entity.GaeUsers{}, dao.Users)
}

func (dao *UserDAO) Get(id int64) error{
	log.Info(id)
	dao.User.Id = id
	if err := dao.BaseDAO.Get(dao.User); err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) Create() error{
	if err := dao.BaseDAO.Create(dao.User); err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) Delete(id int64) error{
	dao.User.Id = id
	if err := dao.BaseDAO.Delete(dao.User); err != nil {
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
}

