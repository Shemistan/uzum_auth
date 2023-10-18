package docs

import "github.com/mvrilo/go-redoc"

func Initialize() redoc.Redoc {
	doc := redoc.Redoc{
		Title:       "Documentation of AuthSystem",
		Description: "Documentation describes working procedures of AuthSystem like structs, handlers, etc.",
		SpecFile:    "./docs/api.swagger.json",
		SpecPath:    "/docs/api.swagger.json",
		DocsPath:    "/docs",
	}

	return doc
}
