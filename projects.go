package wistia

// Projects are the main organizational objects in Wistia
type Projects []Project

// Project represents the main organizational object in Wistia
type Project struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	MediaCount           int    `json:"mediaCount"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
	HashedID             string `json:"hashedId"`
	AnonymousCanUpload   bool   `json:"anonymousCanUpload"`
	AnonymousCanDownload bool   `json:"anonymousCanDownload"`
	Public               bool   `json:"public"`
	PublicID             string `json:"publicId"`
}
