package param

type User struct {
	Id          string `db:"id"`
	Password    string `db:"password"`
	Name        string `db:"name"`
	Academy     string `db:"academy"`
	IdNumber    string `db:"id_number"`
	PhoneNumber string `db:"phone_number"`
	Admin       int    `db:"admin"`
}
type ReqUserGetPublicKey struct {
	Id string `query:"id" validate:"required,len=10"`
}
type ResUserGetPublicKey struct {
	PublicKey string `json:"publickey"`
}
type ReqUserGetToken struct {
	Id       string `query:"id" validate:"required,len=10"`
	Password string `query:"password" validate:"required"`
}
type ResUserGetToken struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}
type ReqUserRegister struct {
	Id          string `query:"id" validate:"required,len=10"`
	Password    string `query:"password" validate:"required"`
	Name        string `query:"name" validate:"required"`
	Academy     string `query:"academy" validate:"required"`
	IdNumber    string `query:"id_number" validate:"required"`
	PhoneNumber string `query:"phone_number" validate:"required"`
}
