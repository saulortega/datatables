# datatables
Simple parser for [DataTables](https://datatables.net/).

# Install

```
go get github.com/saulortega/datatables
```

# Usage

```go
import "github.com/saulortega/datatables"

//Parse receive *http.Request and returns a Filter struct
filter, err = datatables.Parse(r)
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
