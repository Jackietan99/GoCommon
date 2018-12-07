/**
* redis相关操作的处理
* @author:james
* @date:2015-09-09修改
* 包含功能
	字符串的读取和写入
	整形的读取和写入
	浮点的读取和写入
	布尔的读取和写入
	hash读取和写入（golang处理是map）
	set集合的读取和写入
用例：
	redisClient.InitRedis("gc", "tcp", "192.168.1.211", "6379", true)
	d := map[string]string{}
	d["a"] = "1"
	d["b"] = "2"
	d["c"] = "3"
	//map写入hash
	redisClient.HashWrite("myhash", d, 0)
	//读取hash中的field
	s := redisClient.HashReadField("myhash", "b")
	fmt.Println("s=>", s)
	//删除一个field
	redisClient.HashDelField("myhash", "c")
	//读取hash到map
	d2 := redisClient.HashReadAllMap("myhash")
	fmt.Println("d2", d2)

*/
package redisClient

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

// 连接池结构体
type RedisPool struct {
	pool      *redis.Pool
	prefix    string
	redis_pwd string
	redis_db  string
}

//redis调试开关
var debug = false

//被外部调用的结构体对象
var Redis *RedisPool

//加载包就会执行
func init() {

	//初始化redis
	redis_prefix := beego.AppConfig.String("redis_prefix")
	redis_network := beego.AppConfig.String("redis_network")
	redis_addr := beego.AppConfig.String("redis_addr")
	redis_port := beego.AppConfig.String("redis_port")
	redis_pwd := beego.AppConfig.String("redis_pwd")
	redis_db := beego.AppConfig.String("redis_db")
	redis_bug := beego.AppConfig.String("redis_debug")
	if redis_bug == "on" {
		debug = true
	}

	Redis = InitRedis(redis_prefix, redis_network, redis_addr, redis_port, redis_pwd, redis_db)
}

/**
*检查redis连接是否正常
* 正常的话，将返回nil
* @param k string  redis中主健使用的key
* @param n string	定义redis连接类型，一般为tcp
* @param ip string  redis服务器的ip地址
* @param p string redis服务器的端口
* @param w string redis认证密码
 */
func InitRedis(k, n, ip, p, w, db string) *RedisPool {

	pool := &redis.Pool{
		MaxIdle:     80,
		MaxActive:   12000, // max number of connections
		IdleTimeout: 5 * time.Second,
		Dial: func() (redis.Conn, error) {
			host := ip + ":" + p
			c, err := redis.Dial(n, host)
			if err != nil {
				if debug {
					fmt.Println("不能连接到redis->", n, ip, p, err.Error())
				}
			} else {
				if len(w) > 1 {
					//设置访问权限
					c.Send("AUTH", w)
				}
				if debug {
					fmt.Println("成功连接到redis->", n, ip, p)
				}
			}
			return c, err
		},
	}
	return &RedisPool{pool, k, w, db}
}

/**
* 设置调试模式
 */
func SetRedisDebug(debug bool) {
	debug = debug
}

/**
* 返回一个可以操作的链接
 */
func (this *RedisPool) getCon() redis.Conn {
	c := this.pool.Get()
	if len(this.redis_pwd) > 1 {
		c.Send("AUTH", this.redis_pwd)
	}
	if len(this.redis_db) > 0 {
		c.Send("SELECT", this.redis_db)
	}
	if debug {
		fmt.Println("开启连接数->", this.pool.ActiveCount())
	}
	return c
}
