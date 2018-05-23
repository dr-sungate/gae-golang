package dao

import (
	"net/url"
	"reflect"
	"strings"
	"gaeapp/gae-golang/api/client"

)
const (
	QUERY_LIMIT = "limit"
	QUERY_OFFSET = "offset"
	QUERY_ORDER = "order"
)
type BaseDAO struct {
	datastore *client.GaeDatastoreInstance
}

func (dao BaseDAO) GetAll(entity string, query url.Values, limit, offset int, order string, inputstruct interface{}, responsestruct interface{}) error {
	if _, err := dao.datastore.GetAll(entity, dao.GetQueryMap(query, inputstruct), limit, offset, order, responsestruct); err != nil {
		return err
	}
	return nil
}


func (dao BaseDAO) Get(inputstruct interface{}) error{
	if err := dao.datastore.GetOne(inputstruct); err != nil {
		return err
	}
	return nil
}

func (dao BaseDAO) Create(inputstruct interface{}) error{
	if err := dao.datastore.Put(inputstruct); err != nil {
		return err
	}
	return nil
}

func (dao BaseDAO) Delete(inputstruct interface{}) error{
	key := dao.datastore.GetKey(inputstruct)
	if err := dao.datastore.Delete(key); err != nil {
		return err
	}
	return nil
}

func (dao *BaseDAO) GetQueryMap(query url.Values, inputstruct interface{}) map[string]string {
	returnmap := make(map[string]string, 0)
	gaeusertype := reflect.TypeOf(inputstruct)
	for i :=0; i < gaeusertype.NumField(); i++ {
		field := gaeusertype.Field(i)
		if query.Get(field.Name) != "" {
			returnmap[field.Tag.Get("datastore") + " = "] = query.Get(field.Name)
		}
		if query.Get(strings.ToLower(field.Name)) != ""  {
			returnmap[field.Tag.Get("datastore") + " = "] = query.Get(strings.ToLower(field.Name))
		}
	}
	return returnmap
}