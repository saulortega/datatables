// This package parse a http request from datatables (jQuery plugin) to a friendly structure
// More details in https://github.com/saulortega/datatables
// import "github.com/saulortega/datatables"

package datatables

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

//Parameter to string
func ptos(r *http.Request, p string) (string, error) {
	var s = strings.TrimSpace(r.FormValue(p))
	var e error

	if _, exte := r.Form[p]; !exte {
		e = errors.New("«" + p + "» parameter not received")
	}

	return s, e
}

//Parameter to string; error on empty
func ptosNoEmpty(r *http.Request, p string) (string, error) {
	s, e := ptos(r, p)
	if s == "" && e == nil {
		e = errors.New("«" + p + "» parameter empty")
	}

	return s, e
}

//Parameter to int
func ptoi(r *http.Request, p string) (int, error) {
	var i int
	s, e := ptosNoEmpty(r, p)
	if e != nil {
		return i, e
	}

	i, e = strconv.Atoi(s)

	return i, e
}

//Parameter to bool
func ptob(r *http.Request, p string) (bool, error) {
	var b bool
	s, e := ptosNoEmpty(r, p)
	if e != nil {
		return b, e
	}

	b, e = strconv.ParseBool(s)

	return b, e
}
