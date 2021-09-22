package Users

type Login struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}
