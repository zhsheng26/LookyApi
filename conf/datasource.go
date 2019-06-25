package conf

/**
* @author zhangsheng
* @date 2019/6/6
 */
import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"
	"github.com/sirupsen/logrus"
	"time"
)

type Db struct {
	Orm     *xorm.Engine
	RedisDb *redis.Client
}

var Data = &Db{}

func ConnectSQL(oracle Oracle) (db *Db) {
	dbSource := fmt.Sprintf(
		"%s/%s@%s:%s/ORCL",
		oracle.Username,
		oracle.Password,
		oracle.Host,
		oracle.Port,
	)
	orm, err := xorm.NewEngine("oci8", dbSource)
	if err != nil {
		logrus.Panicf("Could not connect oracle due to err: %s \n", err)
	}
	orm.Logger().SetLevel(core.LOG_DEBUG)
	Data.Orm = orm
	return Data
}

func ConnectRedis(conf Redis) (db *Db) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         conf.Address(),
		Password:     conf.Password,
		DB:           conf.Database,
		DialTimeout:  30 * time.Second,
		MinIdleConns: 3,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	if _, err := redisClient.Ping().Result(); err != nil {
		logrus.Panicf("Could not ping redis server due to err: %s \n", err)
		return
	}
	Data.RedisDb = redisClient
	db = Data
	return
}
