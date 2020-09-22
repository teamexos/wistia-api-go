// +build unit

package wistia_test

import (
	"context"
	"errors"
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

func TestFailedRequest(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("error")
	}

	_, err := wistiaClient.MediasShow(ctx, "fakeID", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusInternalServerError)
	assert.EqualValues(t, err.Message, "failed to make request")
}

func TestUnauthorized(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       fixtures.ResponseUnauthorized(),
		}, nil
	}

	_, err := wistiaClient.MediasShow(ctx, "fakeID", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusUnauthorized)
	assert.EqualValues(t, err.Message, "Invalid credentials.")
}

func TestMediaNotFound(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       fixtures.ResponseMediaNotFound(),
		}, nil
	}

	_, err := wistiaClient.MediasShow(ctx, "fakeID", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Message, "Media with ID 123 not found.")
}

func TestProjectNotFound(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       fixtures.ResponseProjectNotFound(),
		}, nil
	}

	_, err := wistiaClient.ProjectsShow(ctx, "fakeID", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Message, "Project with ID '123' not found.")
}

func TestRouteNotFound(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       fixtures.ResponseRouteNotFound(),
		}, nil
	}

	_, err := wistiaClient.ProjectsShow(ctx, "fakeID", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Message, "Route not found")
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

	project, err := wistiaClient.ProjectsShow(ctx, "fakeID", nil)

	assert.Nil(t, err)
	assert.NotNil(t, project)
	assert.EqualValues(t, 464427, project.ID)
	assert.EqualValues(t, "ln2k6qwi9k", project.HashedID)
	assert.EqualValues(t, 3, project.MediaCount)
	assert.EqualValues(t, 3, len(project.Medias))

	for _, media := range project.Medias {
		assert.NotNil(t, media.Section)
	}
}
