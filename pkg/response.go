package pkg

import "oss.nandlabs.io/golly/rest"

func ResponseJSON(ctx rest.ServerContext, status int, data interface{}) {
	ctx.SetStatusCode(status)
	ctx.SetHeader("Content-Type", "application/json")
	ctx.WriteJSON(data)
}
