package models

type Registered struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" gorm:"unique"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}
type User struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type SpamNumber struct {
	Number string `json:"number" gorm:"unique"`
}

type SearchByNameResult struct {
	GlobalUser User
	MarkAsSpam bool
}
