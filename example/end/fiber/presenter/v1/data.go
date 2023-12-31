package presenter

type DataSuccess struct {
	Data interface{} `json:"data,omitempty"`
}

type DataArraySuccess struct {
	Data        interface{} `json:"data,omitempty"`   // Data is the array of data
	ReturnValue int         `json:"return,omitempty"` // ReturnValue is the number of items returned
	Limit       int         `json:"limit,omitempty"`  // Limit is the query limit, what was the query limited to return
	From        interface{} `json:"from,omitempty"`   // From is the first id
	Next        interface{} `json:"next,omitempty"`   // Next is the next id
	Total       int64       `json:"total,omitempty"`  // Total is the total number of items in the query
}

type DataError struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}
