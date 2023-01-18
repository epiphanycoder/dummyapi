package model

type User struct {
	ID           string   `json:"id" gorm:"index"`
	FirstName    string   `json:"firstName" gorm:"-"`
	LastName     string   `json:"lastName" gorm:"-"`
	Email        string   `json:"email" gorm:"primaryKey;index;size:256"`
	Title        string   `json:"title" gorm:"-"`
	Password     string   `json:"-" gorm:"size:128"`
	Status       int      `json:"-"`
	DateOfBirth  string   `json:"dateOfBirth,omitempty" gorm:"-"`
	RegisterDate string   `json:"registerDate,omitempty" gorm:"-"`
	Gender       string   `json:"gender,omitempty" gorm:"-"`
	Phone        string   `json:"phone,omitempty" gorm:"-"`
	Picture      string   `json:"picture,omitempty" gorm:"-"`
	Location     Location `json:"location,omitempty" gorm:"-"`
}

type Location struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	TimeZone string `json:"timezone"`
}

func (u *User) String() string {
	return "User"
}
