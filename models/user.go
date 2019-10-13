package models

//User ...
type User struct {
	UserName string `form:"username"`
	Pwd      string `form:"pwd"`
}
