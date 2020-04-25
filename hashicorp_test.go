package tfvm

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExtractReleases(t *testing.T) {

	html := "foo <a>terraform_1.0</a> bar\n<a>terraform_2.0abc </a> stuff"

	releases, err := extractReleases(html)
	if err != nil {
		t.Errorf("Cannot extract.")
	}

	assert.Equal(t, "1.0", releases[0].version)
	assert.Equal(t, "2.0abc", releases[1].version)
}
