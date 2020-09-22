package fixtures

import (
	"bytes"
	"io"
	"io/ioutil"
)

// MediasShow returns an example Medias#Show response
func MediasShow() io.ReadCloser {
	json := `{"id":4489021,"name":"How They Work - Zappos","type":"Video","created":"2013-09-19T15:30:49+00:00","updated":"2013-10-28T20:53:12+00:00","duration":167,"hashed_id":"v80gyfkt28","description":"<p>\n\nWistia goes to Nevada to visit with Zappos to hear what they have to say about their company culture.&nbsp;</p>\n<p>\n\n&nbsp;</p>\n<p>\n\nFor more How They Work videos, check out:</p>\n<p>\n\n<a href=\"http://jeff.wistia.com/projects/ln2k6qwi9k\">http://jeff.wistia.com/projects/ln2k6qwi9k</a></p>\n","progress":1,"status":"ready","thumbnail":{"url":"http://embed.wistia.com/deliveries/7fbf9c2fe9c6585f9aa032f43f0aecc3f287e86b.jpg?image_crop_resized=100x60","width":100,"height":60},"project":{"id":464427,"name":"How They Work","hashed_id":"ln2k6qwi9k"},"assets":[{"url":"http://embed.wistia.com/deliveries/856970d9a4bcb9aab381a0bd9ab714f19d72c62f.bin","width":960,"height":540,"fileSize":23695556,"contentType":"video/mp4","type":"OriginalFile"},{"url":"http://embed.wistia.com/deliveries/c16c2ef4a87dc8147305637cc302f2e9f9c78977.bin","width":960,"height":540,"fileSize":17493009,"contentType":"video/x-flv","type":"FlashVideoFile"},{"url":"http://embed.wistia.com/deliveries/9e5ead0ef514bef19e3bad9062038c7dad60e10a.bin","width":640,"height":360,"fileSize":19542684,"contentType":"video/mp4","type":"IphoneVideoFile"},{"url":"http://embed.wistia.com/deliveries/7fbf9c2fe9c6585f9aa032f43f0aecc3f287e86b.bin","width":960,"height":540,"fileSize":105070,"contentType":"image/jpeg","type":"StillImageFile"}]}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ProjectsList returns an example Projects#List response
func ProjectsList() io.ReadCloser {
	json := `[{"id":22570,"name":"My Project Title","description":"My Project Description","mediaCount":2,"created":"2010-08-13T18:47:39+00:00","updated":"2010-08-19T21:47:00+00:00","hashedId":"4d23503f70","anonymousCanUpload":false,"anonymousCanDownload":false,"public":false,"publicId":"4bD"},{"id":10495,"name":"Another Project Title","description":"Another Project Description","mediaCount":4,"created":"2010-08-13T18:47:39+00:00","updated":"2010-08-19T21:47:00+00:00","hashedId":"4d23503f70","anonymousCanUpload":false,"anonymousCanDownload":false,"public":false,"publicId":"3dF"}]`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ProjectsShow returns an example Projects#Show response
func ProjectsShow() io.ReadCloser {
	json := `{"anonymousCanDownload":false,"anonymousCanUpload":false,"created":"2013-09-19T15:05:15+00:00","description":"","hashedId":"ln2k6qwi9k","id":464427,"mediaCount":3,"medias":[{"created":"2013-09-19T15:05:17+00:00","description":"","duration":167.29,"hashed_id":"azh1x9nesb","id":4488635,"name":"How They Work: Zappos","progress":1,"status":"ready","thumbnail":{"height":60,"url":"http://embed.wistia.com/deliveries/11bb346da9e041a4ce5c8f0a37803f8ef50e2848.jpg?image_crop_resized=100x60","width":100},"type":"Video","updated":"2013-10-28T20:53:12+00:00"},{"created":"2013-09-19T15:30:49+00:00","description":"<p>\n\nWistia goes to Nevada to visit with Zappos to hear what they have to say about their company culture.&nbsp;</p>\n<p>\n\n&nbsp;</p>\n<p>\n\nFor more How They Work videos, check out:</p>\n<p>\n\n<a href=\"http://jeff.wistia.com/projects/ln2k6qwi9k\">http://jeff.wistia.com/projects/ln2k6qwi9k</a></p>\n","duration":167,"hashed_id":"v80gyfkt28","id":4489021,"name":"How They Work - Zappos","progress":1,"status":"ready","thumbnail":{"height":60,"url":"http://embed.wistia.com/deliveries/7fbf9c2fe9c6585f9aa032f43f0aecc3f287e86b.jpg?image_crop_resized=100x60","width":100},"type":"Video","updated":"2013-10-28T20:53:12+00:00"},{"created":"2013-09-19T15:44:10+00:00","description":"<p>\n\nWistia goes to Nevada to visit with Zappos to hear what they have to say about their company culture.&nbsp;</p>\n<p>\n\n&nbsp;</p>\n<p>\n\nFor more How They Work videos, check out:</p>\n<p>\n\n<a href=\"http://jeff.wistia.com/projects/ln2k6qwi9k\">http://jeff.wistia.com/projects/ln2k6qwi9k</a></p>\n","duration":167,"hashed_id":"cqvzbmyeid","id":4489159,"name":"How They Work - Zappos","progress":1,"status":"ready","thumbnail":{"height":60,"url":"http://embed.wistia.com/deliveries/1d5b471a6a28b5d54787ae01a67e8d20c3b39267.jpg?image_crop_resized=100x60","width":100},"type":"Video","updated":"2013-10-28T20:53:12+00:00"}],"name":"How They Work","public":true,"publicId":"ln2k6qwi9k","updated":"2013-10-28T20:53:12+00:00"}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseMediaNotFound returns an example media not found response
func ResponseMediaNotFound() io.ReadCloser {
	json := `{"error":"Media with ID 123 not found."}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseProjectNotFound returns an example project not found response
func ResponseProjectNotFound() io.ReadCloser {
	json := `{"error":"Project with ID '123' not found."}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseRouteNotFound returns an example Route not found response
func ResponseRouteNotFound() io.ReadCloser {
	json := `{"error":"Route not found"}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseUnauthorized returns an example Unauthorized response
func ResponseUnauthorized() io.ReadCloser {
	json := `{"code":"unauthorized_credentials","error":"Invalid credentials."}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}
