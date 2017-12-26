package ext

import (
	"github.com/gohugoio/hugo/tpl/collections"

	"html/template"
	"os"
)

// Louvre returns a URL to a Louvre-hosted image with the ID and parameters specified.
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
