package models

import (
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

type UserDetail struct {
	Id        int    `db:"id"`
	Age       int    `db:"age"`
	UserName  string `db:"userName"`
	UserEmail string `db:"userEmail"`
}

func GetAllUsers() map[string]*User {
	db := mymysql.Conn()
	rows, err := db.Query("select id,age,userName,userEmail from `user`")
	///获取完毕释放rows，阻止更多的列举
 	defer rows.Close()
	if err != nil {
		beego.Error("操作出错", err)
		return nil
	}
	/**获取表头开始***/
	cols, _ := rows.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	/**获取表头结束***/
	/***遍历列表中的数据start***/
	/// 通过切片存储
	users := make([]UserDetail, 0)
	for rows.Next() {
		var user UserDetail
		rows.Scan(&user.Id, &user.UserName, &user.Age, &user.UserEmail)
		fmt.Println(user)
		users = append(users, user)
	}
	/***遍历列表中的数据end***/
	fmt.Println("users:", users)

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
	/**type UserInfo struct {
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
	return UserList**/
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
