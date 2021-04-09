package entity

//User schema for the user table

type User struct {
	ID 			uint64		`gorm:"primary_key:auto_increment" json:"id"`
	Email		string		`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password	string		`gorm:"->;<-;not null" json:"password"`
	FirstName	string		`grom:"type:varchar(255)" json:"first_name"`
	LastName	string		`grom:"type:varchar(255)" json:"last_name"`
	DataOFBirth	string		`grom:"type:varchar(255)" json:"data_of_birth"`
	AcctBalance	string		`grom:"type:varchar(255);default:0" json:"acct_balance"`
	Phone		string		`grom:"type:varchar(255)" json:"phone"`

	Token    string  `gorm:"-" json:"token,omitempty"`
}