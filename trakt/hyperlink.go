//Package trakt ...
package trakt

import (
	"net/url"

	"github.com/jingweno/go-sawyer/hypermedia"
)
//M ...
type M map[string]interface{}
//Hyperlink ...
type Hyperlink string
//Expand ...
func (l Hyperlink) Expand(m M) (u *url.URL, err error) {
	sawyerHyperlink := hypermedia.Hyperlink(string(l))
	u, err = sawyerHyperlink.Expand(hypermedia.M(m))
	return
}
