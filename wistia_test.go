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
	httpClient   *mocks.MockHTTPClient
	ctx          context.Context
	wistiaClient *wistia.Client
)

func init() {
	httpClient = &mocks.MockHTTPClient{}
	ctx = context.Background()
	wistiaClient = wistia.NewClient(httpClient, "access_token")
}

func TestMediasShow(t *testing.T) {

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixtures.MediasShow(),
		}, nil
	}

	medias, err := wistiaClient.MediasShow(ctx, "fakeID", nil)

	assert.Nil(t, err)
	assert.NotNil(t, medias)
	assert.EqualValues(t, 4489021, medias.ID)
	assert.EqualValues(t, 464427, medias.Project.ID)
	assert.EqualValues(t, "How They Work", medias.Project.Name)
	assert.EqualValues(t, "ln2k6qwi9k", medias.Project.HashedID)
	assert.EqualValues(t, 4, len(medias.Assets))
}

func TestProjectsList(t *testing.T) {

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixtures.ProjectsList(),
		}, nil
	}

	projects, err := wistiaClient.ProjectsList(ctx, nil)

	assert.Nil(t, err)
	assert.NotNil(t, projects)
	assert.EqualValues(t, 2, len(*projects))

	for _, project := range *projects {
		assert.NotNil(t, project.ID)
		assert.NotNil(t, project.Name)
		assert.NotNil(t, project.MediaCount)
		assert.NotNil(t, project.HashedID)
		assert.NotNil(t, project.PublicID)
	}
}

func TestProjectsShow(t *testing.T) {

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixtures.ProjectsShow(),
		}, nil
	}

	project, err := wistiaClient.ProjectShow(ctx, "fakeID", nil)

	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.EqualValues(t, 464427, project.ID)
	assert.EqualValues(t, "ln2k6qwi9k", project.HashedID)
	assert.EqualValues(t, 3, project.MediaCount)
	assert.EqualValues(t, 3, len(project.Medias))

	for _, media := range project.Medias {
		assert.Nil(t, media.Section)
	}
}
