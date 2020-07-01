package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"myAppApi/models/mymysql"
	"myAppApi/models/myredis"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	UserList map[string]*User
)
var (
	UserListTable map[string]*UserListTableDetail
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111---------", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["list"] = &u
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

type UserListTableDetail struct {
	id        string `db:"id"`
	userName  string `db:"userName"`
	userEmail string `db:"userEmail"`
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) string {
	// if u, ok := UserList[uid]; ok {
	// 	return u, nil
	// }
	// return nil, errors.New("User not exists")
	return uid
}

type UserLists struct {
	Id        int
	Username  string
	Password  string
	userEmail string
}

func GetAllUsers() map[string]*User {
	db := mymysql.Conn()
	rows, err := db.Query("select * from user")
	if err != nil {
		beego.Error("操作出错", err)
		return nil
	}
	columns, err_1 := rows.Columns()
	if err_1 != nil {
		panic(err_1.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	fmt.Println("values------------------------------------------------", values)
	conn := myredis.Conn()
	_, err1 := conn.Do("SET", "ceshi", "123456789")
	// 设置redis
	if err1 != nil {
		beego.Error("MULTI HINCRBY for new user registeration:", err1)
		return nil
	}
	// 获取redis中的值
	C, err2 := redis.String(conn.Do("GET", "AUTH"))
	if err2 != nil {
		beego.Error("MULTI HINCRBY for new user registeration:", err2)
		return nil
	}
	///用完后将连接放回连接池
	defer conn.Close()
	fmt.Println("redis------------------------------------------------", C)
	type UserInfo struct {
		id        int    `db:"id"`
		userName  string `db:"userName"`
		userEmail string `db:"userEmail"`
	}
	var u UserInfo
	fmt.Println("rows--------", rows)
	for rows.Next() {
		err = rows.Scan(&u.id, &u.userName, &u.userEmail)
		// fmt.Println("ceshi", u)
	}
	// 更新数据
	_, err = db.Exec("update user set userName='pd' where id=17")
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
