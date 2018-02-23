// This package parse a http request from datatables (jQuery plugin) to a friendly structure
// More details in https://github.com/saulortega/datatables
// import "github.com/saulortega/datatables"

package datatables

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Column struct {
	Data        string
	Name        string
	Index       int
	Orderable   bool
	Searchable  bool
	SearchValue string
	SearchRegex bool
}

type Order struct {
	Column Column
	Dir    string
}

type Filter struct {
	Draw        int
	Start       int
	Length      int
	Order       []Order
	Columns     []Column
	SearchValue string
	SearchRegex bool
}

func Parse(r *http.Request) (Filter, error) {
	return parse(r)
}

func MustParse(r *http.Request) Filter {
	f, err := parse(r)
	if err != nil {
		panic(err)
	}
	return f
}

func parse(r *http.Request) (Filter, error) {
	var E error
	var F = Filter{
		Order:   []Order{},
		Columns: []Column{},
	}

	var errores = []string{}
	var mapColsTmp = make(map[int]map[string]string)
	var mapOrdsTmp = make(map[int]map[string]string)

	for ll, v := range r.Form {
		if regexp.MustCompile(`^columns\[`).MatchString(ll) {
			a := regexp.MustCompile(`^columns\[([0-9]+)\]\[`).ReplaceAllString(ll, "$1-")
			a = regexp.MustCompile(`search\]\[(value|regex)`).ReplaceAllString(a, "search$1")
			a = regexp.MustCompile(`\]$`).ReplaceAllString(a, "")
			p := strings.Split(a, "-")
			i, er := strconv.Atoi(p[0])
			if er != nil {
				errores = append(errores, er.Error())
				continue
			}

			if _, ok := mapColsTmp[i]; !ok {
				mapColsTmp[i] = make(map[string]string)
			}
			mapColsTmp[i][p[1]] = v[0]
		} else if regexp.MustCompile(`^order\[`).MatchString(ll) {
			a := regexp.MustCompile(`^order\[([0-9]+)\]\[`).ReplaceAllString(ll, "$1-")
			a = regexp.MustCompile(`\]$`).ReplaceAllString(a, "")
			p := strings.Split(a, "-")
			i, er := strconv.Atoi(p[0])
			if er != nil {
				errores = append(errores, er.Error())
				continue
			}

			if _, ok := mapOrdsTmp[i]; !ok {
				mapOrdsTmp[i] = make(map[string]string)
			}
			mapOrdsTmp[i][p[1]] = v[0]
		} else if ll == "search[value]" {
			F.SearchValue = strings.TrimSpace(v[0])
		} else if ll == "search[regex]" {
			F.SearchRegex, E = strconv.ParseBool(v[0])
			if E != nil {
				errores = append(errores, E.Error())
			}
		} else if ll == "start" {
			F.Start, E = strconv.Atoi(v[0])
			if E != nil {
				errores = append(errores, E.Error())
			}
		} else if ll == "length" {
			F.Length, E = strconv.Atoi(v[0])
			if E != nil {
				errores = append(errores, E.Error())
			}
		} else if ll == "draw" {
			F.Draw, E = strconv.Atoi(v[0])
			if E != nil {
				errores = append(errores, E.Error())
			}
		}
	}

	for i := 0; i < len(mapColsTmp); i++ {
		c := Column{Index: i}
		c.Data = strings.TrimSpace(mapColsTmp[i]["data"])
		c.Name = strings.TrimSpace(mapColsTmp[i]["name"])
		c.Orderable, E = strconv.ParseBool(mapColsTmp[i]["orderable"])
		if E != nil {
			errores = append(errores, E.Error())
		}
		c.Searchable, E = strconv.ParseBool(mapColsTmp[i]["searchable"])
		if E != nil {
			errores = append(errores, E.Error())
		}
		c.SearchValue = strings.TrimSpace(mapColsTmp[i]["searchvalue"])
		c.SearchRegex, E = strconv.ParseBool(mapColsTmp[i]["searchregex"])
		if E != nil {
			errores = append(errores, E.Error())
		}
		F.Columns = append(F.Columns, c)
	}

	for i := 0; i < len(mapOrdsTmp); i++ {
		dir := strings.ToUpper(mapOrdsTmp[i]["dir"])
		ic, er := strconv.Atoi(mapOrdsTmp[i]["column"])
		if er != nil || (dir != "ASC" && dir != "DESC") {
			if er != nil {
				errores = append(errores, er.Error())
			} else {
				errores = append(errores, "dir invalid: «"+dir+"»")
			}
			continue
		}

		o := Order{Dir: dir}
		for _, c := range F.Columns {
			if c.Index == ic && c.Orderable {
				o.Column = c
				break
			}
		}
		if o.Column == (Column{}) {
			continue
		}

		F.Order = append(F.Order, o)
	}

	if len(errores) > 0 {
		E = errors.New(strings.Join(errores, "; "))
	}

	return F, E
}
