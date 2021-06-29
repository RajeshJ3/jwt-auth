package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User Model
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Encode password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare hash and password
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Before saving User to DB
func (u *User) PreSave() error {
	// Hash the password
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}

	// save hashed password to database
	u.Password = string(hashedPassword)
	return nil
}

// Clean data
func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	if u.Password == "" {
		return errors.New("password required")
	}
	if u.Email == "" {
		return errors.New("email required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}

// Save user to DB
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	// run PreSave
	u.PreSave()
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
