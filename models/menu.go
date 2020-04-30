package models

var (
	MenuList map[string]*Menu
)

func init() {
	MenuList = make(map[string]*Menu)
	u := Menu{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	MenuList["user_11111"] = &u
}

type Menu struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

func GetMenu(uid string) string {
	return "123123"
}
