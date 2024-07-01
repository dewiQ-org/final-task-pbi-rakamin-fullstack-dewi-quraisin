package dto

import "mime/multipart"

type PostResponse struct {
	ID        int  `json:"id"`
	UserID    int  `json:"-"`
	User      User `gorm:"foreignKey:UserID" json:"user"`
	Picture   int  `json:"picture"`
	CreatedAt int  `json:"created_at"`
	UpdatedAt int  `json:"updated_at"`
}

type PostRequest struct {
	UserID  int                   `form:"user_id"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	ID       int    `form:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
