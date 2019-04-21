package auth

import (
	. "tinyUrl/context"
)

func Authenticate(ctx *Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		SendForbidden(ctx)
		return
	}
	ctx.Next()
}
