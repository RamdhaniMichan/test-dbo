package models

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string                 `json:"name"`
	Email          string                 `json:"email" gorm:"unique"`
	Password       string                 `json:"password"`
	PhoneNumber    string                 `json:"phone_number"`
	Address        string                 `json:"address"`
	BirthDate      string                 `json:"birthdate"`
	Gender         string                 `json:"gender"`
	ProfilePicture string                 `json:"profile_picture"`
	Status         string                 `json:"status" gorm:"default:active"`
	Category       string                 `json:"category"`
	Preferences    map[string]interface{} `gorm:"type:jsonb"`
	Role           string                 `json:"role"`
	Orders         []Order                `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	var oldCustomer User
	if err := tx.First(&oldCustomer, u.ID).Error; err != nil {
		return err
	}

	changes := make(map[string]interface{})
	if oldCustomer.Name != u.Name {
		changes["full_name"] = oldCustomer.Name
	}
	if oldCustomer.Email != u.Email {
		changes["email"] = oldCustomer.Email
	}
	if oldCustomer.PhoneNumber != u.PhoneNumber {
		changes["phone_number"] = oldCustomer.PhoneNumber
	}
	if oldCustomer.Address != u.Address {
		changes["address"] = oldCustomer.Address
	}
	if oldCustomer.Status != u.Status {
		changes["status"] = oldCustomer.Status
	}
	if oldCustomer.Category != u.Category {
		changes["category"] = oldCustomer.Category
	}
	if fmt.Sprintf("%v", oldCustomer.Preferences) != fmt.Sprintf("%v", u.Preferences) {
		changes["preferences"] = oldCustomer.Preferences
	}

	if len(changes) > 0 {
		history := UserHistory{
			CustomerID: u.ID,
			Changes:    changes,
			UpdatedAt:  time.Now(),
		}
		if err := tx.Create(&history).Error; err != nil {
			return err
		}
	}

	return nil
}

// Validate function for User
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(3, 100)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
	)
}
