package wistia

// Media represents a media object
type Media struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Created     string  `json:"created"`
	Updated     string  `json:"updated"`
	Duration    float32 `json:"duration"`
	HashedID    string  `json:"hashed_id"`
	Description string  `json:"description"`
	Progress    float32 `json:"progress"`
	Status      string  `json:"status"`
	Section     string  `json:"section"`
	Thumbnail   struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"thumbnail"`
	Project struct {
		ID       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		HashedID string `json:"hashed_id,omitempty"`
	} `json:"project,omitempty"`
	Assets []Asset `json:"assets,omitempty"`
}
