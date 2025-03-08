package models

import (
	"captcha-service/app/config/constant"
	"context"
)

type ReqIDContextKey string

func NewContext(ctx context.Context, requestID string) context.Context {
	k := ReqIDContextKey(constant.RequestID)
	return context.WithValue(ctx, k, requestID)
}

func FromContext(ctx context.Context) (value string) {
	if ctx == nil {
		return
	}

	k := ReqIDContextKey(constant.RequestID)
	if v := ctx.Value(k); v != nil {
		value = v.(string)
	}
	return
}

func GetValueFromContext(ctx context.Context, key string) (value string) {
	if ctx == nil {
		return
	}

	k := ReqIDContextKey(key)
	if v := ctx.Value(k); v != nil {
		value = v.(string)
	}
	return
}
