package request

type Signin struct {
	Username string `db:"username,use_zero" json:"username" form:"username" query:"username" validate:"required"`
	Passwd   string `db:"passwd,use_zero" json:"passwd" form:"passwd" query:"passwd" validate:"required,lte=200"`
}

type ChangePassword struct {
	CurrentPasswd string `json:"currentPasswd" form:"currentPasswd" validate:"required,lte=200"`
	Passwd        string `json:"passwd" form:"passwd" validate:"required,lte=200"`
	ConfirmPasswd string `json:"confirmPasswd" form:"confirmPasswd" validate:"required,lte=200,eqfield=Passwd"`
}

type CreateUser struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required,lte=80"`
	Email    string `json:"email" form:"email" validate:"required,lte=200,email,notexists=email email"`
	NoHp     string `json:"noHp" form:"noHp" validate:"required,lte=20,notexists=no_hp noHp"`
	Username string `json:"username" form:"username" validate:"required,lte=20,lowercase,notexists=username username"`
	Passwd   string `json:"passwd" form:"passwd" validate:"required,lte=200"`
}

type UpdateUser struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required,lte=80"`
	Email    string `json:"email" form:"email" validate:"required,lte=200,email,notexists=email email"`
	NoHp     string `json:"noHp" form:"noHp" validate:"required,lte=20,notexists=no_hp noHp"`
	Username string `json:"username" form:"username" validate:"required,lte=20,lowercase,notexists=username username"`
}

type PageUser struct {
	Paging
	Email    string `json:"email" form:"email" query:"email"`
	NoHp     string `json:"noHp" form:"noHp" query:"noHp"`
	Username string `json:"username" form:"username" query:"username"`
}
