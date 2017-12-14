package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"teo-service/utils/kripto"
	"time"
)

type User struct {
	Id            int       `form:"userID" json:"userid"`
	Username      string    `form:"username" json:"username"`
	Email         string    `form:"email" json:"email"`
	Password      string    `form:"password" json:"-"`
	Type          string    `form:"-" json:"type"`
	RememberToken string    `form:"-" json:"-"`
	CreatedAt     time.Time `form:"-" json:"join_date" orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `form:"-" json:"-" orm:"auto_now;type(datetime)`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

func InsertUser(user User) int {
	o := orm.NewOrm()
	user.Password = kripto.Encrypt(user.Password)
	res, err := o.Raw("INSERT INTO users (username, email, password, type, created_at, updated_at) VALUES(?,?,?,?,?,?)", user.Username, user.Email, user.Password, 1, time.Now(), time.Now()).Exec()
	if err == nil {
		ret, _ := res.LastInsertId()
		fmt.Println("Id yang ditambahkan: ", ret)
		return int(ret)
	}

	return 0
}

func RetriveNameBasedID(id int) string {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Raw("SELECT id, username, email, type FROM users WHERE id= ?", user.Id).QueryRow(&user)
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
		return user.Username
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
		return user.Username
	} else {
		return user.Username
	}
}

func RetrieveOneUser(usernm string, pwd string) User {
	o := orm.NewOrm()
	user := User{Username: usernm, Password: pwd, Id: 0}
	var pwdValidation string
	err := o.Raw("SELECT password FROM users WHERE username = ?", usernm).QueryRow(&pwdValidation)
	if err == orm.ErrNoRows {
		fmt.Println("User Not Found")
		return user
	}
	if kripto.Decrypt(pwdValidation) == pwd {
		err = o.Raw("SELECT id, username, email, type FROM users WHERE username = ?", usernm).QueryRow(&user)
	}
	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
		return user
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
		return user
	} else {
		return user
	}
}
func RetrieveAllUser() []User {
	o := orm.NewOrm()
	var u []User
	_, _ = o.Raw("SELECT id, username FROM users").QueryRows(&u)
	return u
}
