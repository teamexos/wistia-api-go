// +build unit

package wistia_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamexos/wistia-api-go"
	"github.com/teamexos/wistia-api-go/testdata/fixtures"
	"github.com/teamexos/wistia-api-go/testdata/mocks"
)

var (
	restClient *mocks.MockHTTPClient
	ctx        context.Context
)

func init() {
	restClient = &mocks.MockHTTPClient{}
	ctx = context.Background()
}

func TestMediasShow(t *testing.T) {

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixtures.MediasShow(),
		}, nil
	}

	c := wistia.NewClient(restClient, "access_token")
	medias, err := c.MediasShow(ctx, "fakeID", nil)

	assert.Nil(t, err)
	assert.NotNil(t, medias)
	assert.EqualValues(t, 4489021, medias.ID)
	assert.EqualValues(t, 464427, medias.Project.ID)
	assert.EqualValues(t, "How They Work", medias.Project.Name)
	assert.EqualValues(t, "ln2k6qwi9k", medias.Project.HashedID)
	assert.EqualValues(t, 4, len(medias.Assets))
}
