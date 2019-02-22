package ext

import (
	"fmt"
	"html/template"
	"os"
)

// Louvre returns a URL to a Louvre-hosted image with the ID and parameters specified.
func (ns *Namespace) Louvre(id string, params ...interface{}) (template.URL, error) {
	if len(params)%2 == 1 {
		return "", fmt.Errorf("louvre should receive an even number of params")
	}

	m := map[string]interface{}{}
	for i := 0; i < len(params); i += 2 {
		k, ok := params[i].(string)
		if !ok {
			return "", fmt.Errorf("param %d should be a string (as a key)", i+1)
		}

		m[k] = params[i+1]
	}

	u := ns.deps.Cfg.GetString("params.louvre_url")
	if u == "" {
		u = os.Getenv("HUGO_LOUVRE_URL")
	}

	mw := getStr(m, "mw")
	mh := getStr(m, "mh")
	ext := getStr(m, "ext")

	if mw == "" && mh == "" && ext == "" {
		return template.URL(fmt.Sprintf("%s/%s", u, id)), nil
	}

	if ext == "" {
		return template.URL(fmt.Sprintf("%s/%s/%sx%s", u, id, mw, mh)), nil
	}

	return template.URL(fmt.Sprintf("%s/%s/%sx%s.%s", u, id, mw, mh, ext)), nil
}

func getStr(m map[string]interface{}, k string) string {
	if v, ok := m[k]; ok {
		return fmt.Sprintf("%v", v)
	}

	return ""
}
