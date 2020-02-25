package model

import (
	"chat/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 操作Redis

type UserDao struct {
	pool *redis.Pool //redis 连接池
}

//服务器启动后创建一个 userdao实例|全局的变量，需要和redis操作是，直接使用
var (
	MyUserDao *UserDao
)

//使用工厂模式 创建一个userdao
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool: pool}
	return
}

// 注册用户
func (this *UserDao) Register(user *message.User) (err error) {

	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserID)
	if err == nil {
		//用户存在了
		err = ERROR_USER_EXISTS
		return
	}

	//这时 redis中还不存在，可以注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("hset", "users", user.UserID, string(data))
	if err != nil {
		fmt.Println("保存注册用户 错误 err", err)
		return
	}
	return
}

// 根据id 返回一个user
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定id去redis这个用户

	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		fmt.Println("根据ID查询redis错误")
		if err == redis.ErrNil {
			fmt.Println("没有对应ID的用户")
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{} //这里是不是多余的呢
	//需要把res序列化成user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json unmarshal err ", err)
		return
	}
	return
}

// 登录
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	if userPwd != user.UserPwd {
		//密码错误
		err = ERROR_USER_PWD
		return
	}
	return
}
