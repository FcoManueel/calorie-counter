package ccontext

import (
	"golang.org/x/net/context"
	"gopkg.in/stretchr/testify.v1/assert"
	"testing"
)

func TestSetGetDebug(t *testing.T) {
	ctx := context.Background()
	assert.False(t, IsDebug(ctx), "isDebug should be false if unset")

	ctx = SetDebug(ctx, false)
	assert.False(t, IsDebug(ctx), "isDebug should be false")

	ctx = SetDebug(ctx, true)
	assert.True(t, IsDebug(ctx), "isDebug should be true")
}

func TestSetGetUserID(t *testing.T) {
	ctx := context.Background()
	assert.Empty(t, GetUserID(ctx))

	ctx = SetUserID(ctx, "prisoner 24601")
	assert.Equal(t, "prisoner 24601", GetUserID(ctx))
}

func TestSetGetRole(t *testing.T) {
	ctx := context.Background()
	assert.Empty(t, GetRole(ctx))

	ctx = SetRole(ctx, "hello")
	assert.Equal(t, "hello", GetRole(ctx))
}
