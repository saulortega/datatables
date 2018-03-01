// This package parse a http request from datatables (jQuery plugin) to a friendly structure
// More details in https://github.com/saulortega/datatables
// import "github.com/saulortega/datatables"

package datatables

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
	Error           error       `json:"error,omitempty"`
}

func (r *Response) Check() error {
	var err error

	if r.Data == nil || r.RecordsFiltered == 0 {
		r.Data = []string{} //Vacío, no nulo
	}

	if r.RecordsFiltered > r.RecordsTotal {
		err = errors.New("wrong filtered or total records")
	}

	return err
}

func (r *Response) WriteResponseOnSuccess(w http.ResponseWriter) error {
	err := r.Check()
	if err != nil {
		return err
	}

	RJSON, err := json.Marshal(r)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(RJSON)

	return nil
}

func (r *Response) WriteResponse(w http.ResponseWriter) error {
	err := r.WriteResponseOnSuccess(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return err
}
