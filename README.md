# datatables
Parser for Datatables

# Usage

filter, err = datatables.Parse(r)
if err != nil {
  //Handle error
}

# Struct

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
