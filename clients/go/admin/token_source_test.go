package admin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)
import "golang.org/x/oauth2"

type DummyTestTokenSource struct {
	oauth2.TokenSource
}

func (d DummyTestTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "abc",
	}, nil
}

func TestNewTokenSource(t *testing.T) {
	tokenSource := DummyTestTokenSource{}
	flyteTokenSource := NewTokenSourcePerCallCredential(tokenSource, "test")
	metadata, err := flyteTokenSource.GetRequestMetadata(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "Bearer abc", metadata["test"])
}