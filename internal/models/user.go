package models

type User struct {
	Id             string  `json:"id"`
	Username       string  `json:"username"`
	HashedPassword string  `json:"hashedPassword"`
	ProfilePicture *string `json:"profilePicture"`
	Role           Role    `json:"role"`
}
