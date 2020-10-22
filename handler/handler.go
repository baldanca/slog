package handler

type (
	// Func type
	Func func(values ...interface{}) []interface{}
	// Handlers store type
	Handlers []Func
)

// NewHandlers factory
func NewHandlers() Handlers {
	return Handlers{}
}

// Add handler function
func (hs Handlers) Add(f Func) {
	hs = append(hs, f)
}

// Run handlers function
func (hs Handlers) Run(values ...interface{}) []interface{} {
	for _, handler := range hs {
		values = handler(values...)
	}
	return values
}
