package wistia

// Projects are the main organizational objects in Wistia
type Projects []Project

// Project represents the main organizational object in Wistia
type Project struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	MediaCount           int     `json:"mediaCount"`
	Medias               []Media `json:"medias,omitempty"`
	Created              string  `json:"created"`
	Updated              string  `json:"updated"`
	HashedID             string  `json:"hashedId"`
	AnonymousCanUpload   bool    `json:"anonymousCanUpload"`
	AnonymousCanDownload bool    `json:"anonymousCanDownload"`
	Public               bool    `json:"public"`
	PublicID             string  `json:"publicId"`
}

// PaginationOptions defines the pagination
// options supported by Wistia endpoints
type PaginationOptions struct {
	Page          int    `json:"page"`
	PerPage       int    `json:"per_page"`
	SortBy        string `json:"sort_by"`
	SortDirection int    `json:"sort_direction"`
}
