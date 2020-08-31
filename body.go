package templates

import (
	"bytes"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type body struct {
	body  []byte
	cache *cache
}

func (a *body) Bytes() []byte {
	return a.body
}

func (a *body) String() string {
	return string(a.body)
}

func (a *body) Replace(k, v string) *body {
	if a == nil {
		return a
	}
	a.body = bytes.ReplaceAll(a.body, []byte(k), []byte(v))
	return a
}

func (a *body) ReplaceJson(k string, v interface{}) *body {
	if a == nil {
		return a
	}
	b, _ := ffjson.Marshal(v)
	a.body = bytes.ReplaceAll(a.body, []byte(k), b)
	return a
}

func (a *body) ReplaceBytes(k string, v []byte) *body {
	if a == nil {
		return a
	}
	a.body = bytes.ReplaceAll(a.body, []byte(k), v)
	return a
}

/* set cache */
func (a *body) Cache(path string, t ...time.Duration) (body []byte) {
	a.cache.Add(path, a.body, t...)
	return a.body
}
