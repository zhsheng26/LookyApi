package support

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

/**
* @author zhangsheng
* @date 2019/6/17
 */
type BaseEntity struct {
	Id         string `json:"id" xorm:"'T_ID' unique"`
	CreateDate UsTime `json:"createDate" xorm:"'CREATE_DATE' CREATED"`
	UpdateDate UsTime `json:"updateDate"xorm:"'UPDATE_DATE' UPDATED"`
	Deleted    bool   `json:"deleted" xorm:"'DELETED' bit(1)"`
}
type BaseCache struct {
	RedisDb *redis.Client
}

func (cache *BaseCache) MarshalBinary(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (cache *BaseCache) UnmarshalBinary(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

func (cache *BaseCache) Key(args ...string) string {
	if len(args) == 1 {
		return args[0]
	}
	return strings.Join(args, "_")
}

func (cache *BaseCache) GetCache(target interface{}, field string, key ...string) error {
	sds, err := cache.RedisDb.HGet(cache.Key(key...), field).Result()
	if err != nil {
		return errors.NewNotFound(err, "not find cache")
	}
	if err = cache.UnmarshalBinary([]byte(sds), target); err != nil {
		logrus.Error(err)
	}
	return err
}

func (cache *BaseCache) SetCache(target interface{}, field string, key ...string) {
	bytes, err := cache.MarshalBinary(target)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = cache.RedisDb.HSet(cache.Key(key...), field, bytes).Err()
	if err != nil {
		logrus.Error(err)
	}
}
