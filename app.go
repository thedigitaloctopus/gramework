package gramework

// ToTLSHandler returns handler that redirects user to HTTP scheme
func (app *App) ToTLSHandler() func(*Context) {
	return func(ctx *Context) {
		ctx.ToTLS()
	}
}

// HTTP returns HTTP-only router
func (app *App) HTTP() *Router {
	return app.defaultRouter.HTTP()
}

// HTTPS returns HTTPS-only router
func (app *App) HTTPS() *Router {
	return app.defaultRouter.HTTPS()
}
