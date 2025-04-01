package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()

	// Note : Below written code would not be necessary if we want to serve our
	// files in the folder "static" under "/static/" endpoint. In case we give "/",
	// it is served on the root endpoint, or if we give the path as say, "/website" or "/website/",
	// GoFr resolves both path to serve the files as "/website/*" with "/website/" directing automatically
	// to present index.html file.
	app.AddStaticFiles("/", "./website")

	app.Run()
}
