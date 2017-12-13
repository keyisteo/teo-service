package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id            int       `form:"userID" json:"userid"`
	Username      string    `form:"username" json:"username"`
	Email         string    `form:"email" json:"email"`
	Password      string    `form:"password" json:"-"`
	Type          string    `form:"-" json:"type"`
	RememberToken string    `form:"-" json:"-"`
	CreatedAt     time.Time `form:"-" json:"join_date"`
	UpdatedAt     time.Time `form:"-" json:"-"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

func InsertUser() {
	//o := orm.NewOrm()
}

func RetrieveOneUser(usernm string, pwd string) User {
	o := orm.NewOrm()
	user := User{Username: usernm, Password: pwd}
	err := o.Raw("SELECT id, username, email, type FROM users WHERE username = ? AND password=?", usernm, pwd).QueryRow(&user)
	if err == orm.ErrNoRows {
		user.Id = 0
		fmt.Println("No result found.")
		return user
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
		user.Id = 0
		return user
	} else {
		return user
	}
}
