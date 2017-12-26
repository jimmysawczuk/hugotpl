package ext

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Curator returns a CDN-prefixed URL for the image with the ID and max dimensions specified.
func (ns *Namespace) Curator(id string, maxWidth int, maxHeight int, args ...interface{}) (string, error) {
	if len(args)%2 != 0 {
		return "", errors.New("curator: extra arguments to curator must be multiple of 2")
	}

	extra := map[string]interface{}{}
	for i := 0; i < len(args); i = i + 2 {
		k := args[i].(string)
		v := args[i+1]

		extra[k] = v
	}

	extraHash := ""
	if len(extra) > 0 {
		h := md5.New()
		for k, v := range extra {
			io.WriteString(h, fmt.Sprintf("%s=%v;", k, v))
		}
		extraHash = fmt.Sprintf("-%x", h.Sum(nil))
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "curator: getwd")
	}

	err = os.MkdirAll(filepath.Join(wd, ".curator-build"), 0755)
	if err != nil {
		return "", errors.Wrap(err, "curator: mkdirall")
	}

	filename := fmt.Sprintf("%s-%d-%d%s.curator", id, maxWidth, maxHeight, extraHash)
	body := fmt.Sprintf("id = %q\nwidth = %d\nheight = %d", id, maxWidth, maxHeight)
	for k, v := range extra {
		body = body + fmt.Sprintf("\n%s = %v", k, v)
	}
	body = body + "\n"

	err = ioutil.WriteFile(filepath.Join(wd, ".curator-build", filename), []byte(body), 0644)
	if err != nil {
		return "", errors.Wrap(err, "curator: writefile")
	}

	cdnURL := os.Getenv("CURATOR_CDN_URL")

	return fmt.Sprintf("%s/%s/%dx%d", cdnURL, id, maxWidth, maxHeight), nil
}
