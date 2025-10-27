package entities

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	UserName string `gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Avatar   *string
}
