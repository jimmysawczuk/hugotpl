package louvre

import (
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/tpl/collections"

	"html/template"
	"os"
)

// New returns a new instance of the collections-namespaced template functions.
func New(deps *deps.Deps) *Namespace {
	return &Namespace{
		deps: deps,
	}
}

// Namespace provides template functions for the "collections" namespace.
type Namespace struct {
	deps *deps.Deps
}

// Apply takes a map, array, or slice and returns a new slice with the function fname applied over it.
func (ns *Namespace) Louvre(id string, params ...interface{}) (template.URL, error) {

	querify := collections.New(ns.deps).Querify

	u := ns.deps.Cfg.GetString("params.louvre_url")
	if u == "" {
		u = os.Getenv("HUGO_LOUVRE_URL")
	}

	q, err := querify(params...)
	if err != nil {
		return "", err
	}

	return template.URL(u + "/" + id + "?" + q), nil
}
