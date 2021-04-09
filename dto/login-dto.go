package dto

//RegisterDTO is use to to receive data on registration
type LoginDTO struct {
	Email		string		`json:"email" form:email binding:required`
	Password	string		`json:"password" form:password binding:required`

}