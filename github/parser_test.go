package github

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	f, err := os.Open("testdata/go_daily.html")
	require.NoError(t, err)
	defer f.Close()

	repos, err := Parse(f)
	require.NoError(t, err)
	assert.NotEmpty(t, repos)

	for _, r := range repos {
		assert.NotEmpty(t, r.Owner)
		assert.NotEmpty(t, r.Name)
		assert.NotEmpty(t, r.Description)
		assert.NotEmpty(t, r.Language)
		assert.NotEmpty(t, r.LanguageColor)
		assert.NotEmpty(t, r.StarsTotal)
		assert.NotEmpty(t, r.Forks)
		assert.NotEmpty(t, r.StarsSince)
	}
}
