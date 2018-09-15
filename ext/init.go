package ext

import (
	"github.com/gohugoio/hugo/deps"
	hugointernal "github.com/gohugoio/hugo/tpl/internal"
)

const name = "ext"

func init() {
	f := func(d *deps.Deps) *hugointernal.TemplateFuncsNamespace {
		ctx := New(d)

		ns := &hugointernal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.Louvre,
			[]string{"louvre"},
			[][2]string{},
		)

		return ns

	}

	hugointernal.AddTemplateFuncsNamespace(f)
}

// New returns a new instance of the ext-namespaced template functions.
func New(deps *deps.Deps) *Namespace {
	return &Namespace{
		deps: deps,
	}
}

// Namespace provides the "louvre" template function.
type Namespace struct {
	deps *deps.Deps
}
