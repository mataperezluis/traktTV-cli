// Package trakt ...
package trakt  // import trakt

import "fmt"

//AuthMethod ...
type AuthMethod interface {
	fmt.Stringer
}

//TokenAuth ...
type TokenAuth struct {
	AccessToken string
}

func (t TokenAuth) String() string {
	return fmt.Sprintf("Bearer %s", t.AccessToken)
}
