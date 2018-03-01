# datatables
Simple parser for [DataTables](https://datatables.net/) server-side processing.

# Install

```
go get -u github.com/saulortega/datatables
```

# Usage

```go
import "github.com/saulortega/datatables"

//Parse receive *http.Request and returns a Filter struct
filter, err = datatables.Parse(r)
if err != nil {
	//Handle error
}

//Get data from DB

response := filter.PrepareResponse()
response.RecordsTotal = 629635
response.RecordsFiltered = 50
response.Data = rows

//WriteResponse receive http.ResponseWriter. It send the response even if there are any error.
//Use WriteResponseOnSuccess(w) if you do not want to send the response when there is an error.
err := response.WriteResponse(w)
if err != nil {
	//Handle error
}

```

# Struct

```go
type Filter struct {
	Draw        int
	Start       int
	Length      int
	Order       []Order
	Columns     []Column
	SearchValue string
	SearchRegex bool
}

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
```
