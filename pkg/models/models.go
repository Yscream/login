package models

type User struct {
	ID       uint   `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type Image struct {
	ID        uint   `json:"-" db:"id"`
	UserID    uint   `json:"user_id" db:"user_id"`
	ImagePath string `json:"image_path" db:"image_path"`
	ImageURL  string `json:"image_url" db:"image_url"`
}
