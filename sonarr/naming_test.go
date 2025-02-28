package sonarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/craigjmidwinter/starr"
	"github.com/craigjmidwinter/starr/sonarr"
)

const namingBody = `{
	"renameEpisodes": false,
	"replaceIllegalCharacters": true,
	"multiEpisodeStyle": 0,
	"standardEpisodeFormat": "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
	"dailyEpisodeFormat": "{Series Title} - {Air-Date} - {Episode Title} {Quality Full}",
	"animeEpisodeFormat": "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
	"seriesFolderFormat": "{Series Title}",
	"seasonFolderFormat": "Season {season}",
	"specialsFolderFormat": "Specials",
	"includeSeriesTitle": true,
	"includeEpisodeTitle": false,
	"includeQuality": false,
	"replaceSpaces": true,
	"separator": " - ",
	"numberStyle": "S{season:00}E{episode:00}",
	"id": 1
}`

func TestGetNaming(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   namingBody,
			WithResponse: &sonarr.Naming{
				ID:                       1,
				RenameEpisodes:           false,
				ReplaceIllegalCharacters: true,
				MultiEpisodeStyle:        0,
				StandardEpisodeFormat:    "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				DailyEpisodeFormat:       "{Series Title} - {Air-Date} - {Episode Title} {Quality Full}",
				AnimeEpisodeFormat:       "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				SeriesFolderFormat:       "{Series Title}",
				SeasonFolderFormat:       "Season {season}",
				SpecialsFolderFormat:     "Specials",
				IncludeSeriesTitle:       true,
				IncludeEpisodeTitle:      false,
				IncludeQuality:           false,
				ReplaceSpaces:            true,
				Separator:                " - ",
				NumberStyle:              "S{season:00}E{episode:00}",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*sonarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNaming()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNaming(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &sonarr.Naming{
				ReplaceIllegalCharacters: true,
			},
			ExpectedRequest: `{"replaceIllegalCharacters":true}` + "\n",
			ResponseBody:    namingBody,
			WithResponse: &sonarr.Naming{
				ID:                       1,
				RenameEpisodes:           false,
				ReplaceIllegalCharacters: true,
				MultiEpisodeStyle:        0,
				StandardEpisodeFormat:    "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				DailyEpisodeFormat:       "{Series Title} - {Air-Date} - {Episode Title} {Quality Full}",
				AnimeEpisodeFormat:       "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				SeriesFolderFormat:       "{Series Title}",
				SeasonFolderFormat:       "Season {season}",
				SpecialsFolderFormat:     "Specials",
				IncludeSeriesTitle:       true,
				IncludeEpisodeTitle:      false,
				IncludeQuality:           false,
				ReplaceSpaces:            true,
				Separator:                " - ",
				NumberStyle:              "S{season:00}E{episode:00}",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "PUT",
			WithRequest: &sonarr.Naming{
				ReplaceIllegalCharacters: true,
			},
			ExpectedRequest: `{"replaceIllegalCharacters":true}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*sonarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNaming(test.WithRequest.(*sonarr.Naming))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
