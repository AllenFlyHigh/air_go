package models

type User struct {
	//数据表名是字段名的蛇形小写，例如 CreateTime  create_time
	Id int64
	Username string `gorm:"type:varchar(50)"`
	Realname string `gorm:"type:varchar(50)"`
	Passwd string `gorm:"type:varchar(50)"`
	Email string`gorm:"type:varchar(50)"`
	Phone string `gorm:"type:varchar(20)"`
	Sex int8
	RoleId int8 `gorm:"not null"`
}

func CheckUser(username, password string) bool {
	var user User
	db.Select("id").Where(&User{Username: username, Passwd: password}).First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

func AddUser (data map[string] interface{}) bool {
	db.Create(&User{
		Id: int64(data["id"].(int)),
		Username: data["username"].(string),
		Realname: data["realname"].(string),
		Passwd: data["password"].(string),
		Email: data["email"].(string),
		Phone: data["phone"].(string),
		Sex: int8(data["sex"].(int)),
		RoleId: int8(data["roleId"].(int)),
	})
	return true
}

func ListUsers () [] User{
	var users [] User
	db.Find(&users)
	return users
}