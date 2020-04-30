package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
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

func GetAllUsers() map[string]*User {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/ZHDJ?charset=utf8")
	if err != nil {
		beego.Error("连接数据库出错", err)
		return nil
	} else {
		beego.Info("连接数据库成功")
	}
	rows, err := db.Query("select id,userName from user")
	type UserInfo struct {
		id       int
		userName string
	}
	var u UserInfo
	fmt.Println("rows--------", rows)
	for rows.Next() {
		err = rows.Scan(&u.id, &u.userName)
		fmt.Println(u)
	}
	// 更新数据
	_, err = db.Exec("update user set userName='pd' where id=17")

	fmt.Println("testing...")
	defer db.Close()
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
