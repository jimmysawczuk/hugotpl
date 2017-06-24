# hugotpl

Some extra hugo template functions.

## Installation

```bash
# Get both source trees
go get github.com/gohugoio/hugo
go get github.com/jimmysawczuk/hugotpl

# Install the ext set of functions:
cd $GOPATH/gohugoio/hugo/tpl
ln -s ../../../jimmysawczuk/hugotpl/ext ext
```

Then, in `$GOPATH/gohugoio/hugo/tpl/tplimpl/template_funcs.go`, add the following import:

```go
import(
    _ "github.com/gohugoio/hugo/tpl/ext" // symlink to $GOPATH/github.com/jimmysawczuk/hugotpl/ext
)
```

Finally, rebuild Hugo:

```bash
go install github.com/gohugoio/hugo
```

