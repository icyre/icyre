package models

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleArtist  Role = "ARTIST"
	RoleManager Role = "MANAGER"
	RoleUser    Role = "USER"
)
