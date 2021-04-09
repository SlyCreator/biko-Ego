package dto

//RegisterDTO is use to to receive data on registration
type RegisterDTO struct {
	Email		string		`json:"email" form:email binding:required`
	Password	string		`json:"password" form:password binding:required`
	FirstName	string		`json:"first_name" from:first_name binding:required`
	LastName	string		`json:"last_name" from:last_name binding:required`
}