package ccontext

import (
	"golang.org/x/net/context"
	"log"
)

type contextKey int

const (
	userIDKey contextKey = iota + 1
	roleKey
	debugKey
)

//SetUserID  stores the userID into the context
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

//GetUserID gets the userID from the context
func GetUserID(ctx context.Context) string {
	return getContextStringValue(ctx, userIDKey)
}

// SetRole  stores the role into the context
func SetRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

// GetRole gets the role from the context
func GetRole(ctx context.Context) string {
	return getContextStringValue(ctx, roleKey)
}

//SetDebug  stores the debug flag into the context
func SetDebug(ctx context.Context, debug bool) context.Context {
	return context.WithValue(ctx, debugKey, debug)
}

//IsDebug checks if in more descriptive logging mode
func IsDebug(ctx context.Context) bool {
	return getContextBooleanValue(ctx, debugKey)
}

func getContextBooleanValue(ctx context.Context, key contextKey) bool {
	value, ok := ctx.Value(key).(bool)
	if ok {
		return value
	}
	return false
}

func getContextStringValue(ctx context.Context, key contextKey) string {
	value, ok := ctx.Value(key).(string)
	if ok {
		return value
	}
	log.Printf("Context has no value for key %s returning empty string", contextKeyToString(key))
	return ""
}

func contextKeyToString(key contextKey) string {
	switch key {
	case userIDKey:
		return "userID"
	case roleKey:
		return "role"
	case debugKey:
		return "debug"
	}
	return "Unknown key"
}
