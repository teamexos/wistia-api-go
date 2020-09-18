package wistia

// Asset represents a media asset
type Asset struct {
	URL         string `json:"url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FileSize    int    `json:"fileSize"`
	ContentType string `json:"contentType"`
	Type        string `json:"type"`
}
