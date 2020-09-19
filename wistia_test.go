// +build unit

package wistia_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamexos/wistia-api-go"
	"github.com/teamexos/wistia-api-go/testdata/mocks"
)

var restClient *mocks.MockHTTPClient

func init() {
	restClient = &mocks.MockHTTPClient{}
}

func TestMediasShow(t *testing.T) {
	json := `{"id":4489021,"name":"How They Work - Zappos","type":"Video","created":"2013-09-19T15:30:49+00:00","updated":"2013-10-28T20:53:12+00:00","duration":167,"hashed_id":"v80gyfkt28","description":"<p>\n\nWistia goes to Nevada to visit with Zappos to hear what they have to say about their company culture.&nbsp;</p>\n<p>\n\n&nbsp;</p>\n<p>\n\nFor more How They Work videos, check out:</p>\n<p>\n\n<a href=\"http://jeff.wistia.com/projects/ln2k6qwi9k\">http://jeff.wistia.com/projects/ln2k6qwi9k</a></p>\n","progress":1,"status":"ready","thumbnail":{"url":"http://embed.wistia.com/deliveries/7fbf9c2fe9c6585f9aa032f43f0aecc3f287e86b.jpg?image_crop_resized=100x60","width":100,"height":60},"project":{"id":464427,"name":"How They Work","hashed_id":"ln2k6qwi9k"},"assets":[{"url":"http://embed.wistia.com/deliveries/856970d9a4bcb9aab381a0bd9ab714f19d72c62f.bin","width":960,"height":540,"fileSize":23695556,"contentType":"video/mp4","type":"OriginalFile"},{"url":"http://embed.wistia.com/deliveries/c16c2ef4a87dc8147305637cc302f2e9f9c78977.bin","width":960,"height":540,"fileSize":17493009,"contentType":"video/x-flv","type":"FlashVideoFile"},{"url":"http://embed.wistia.com/deliveries/9e5ead0ef514bef19e3bad9062038c7dad60e10a.bin","width":640,"height":360,"fileSize":19542684,"contentType":"video/mp4","type":"IphoneVideoFile"},{"url":"http://embed.wistia.com/deliveries/7fbf9c2fe9c6585f9aa032f43f0aecc3f287e86b.bin","width":960,"height":540,"fileSize":105070,"contentType":"image/jpeg","type":"StillImageFile"}]}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}

	c := wistia.NewClient(restClient, "access_token")
	medias, err := c.MediasShow(context.Background(), "fakeID", nil)

	assert.Nil(t, err)
	assert.NotNil(t, medias)
	assert.EqualValues(t, 4489021, medias.ID)
	assert.EqualValues(t, 464427, medias.Project.ID)
	assert.EqualValues(t, "How They Work", medias.Project.Name)
	assert.EqualValues(t, "ln2k6qwi9k", medias.Project.HashedID)
	assert.EqualValues(t, 4, len(medias.Assets))
}
