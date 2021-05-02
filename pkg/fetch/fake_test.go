package fetch_test

import (
	"github.com/eh-am/i3-tree-viewer/pkg/fetch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromFake(t *testing.T) {
	f := fetch.FromFake{}
	got, gotErr := f.Fetch()

	assert.Nil(t, gotErr)
	assert.NotNil(t, got)
}
