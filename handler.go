package slog

type (
	// Handler function type
	Handler func(v ...interface{}) []interface{}
	// Handlers store type
	Handlers []Handler
)

// NewHandlers function
func NewHandlers() *Handlers {
	return new(Handlers)
}

// Add handler
func (hs *Handlers) Add(h Handler) {
	*hs = append(*hs, h)
}

// run handlers function
func (hs *Handlers) run(v ...interface{}) []interface{} {
	for _, handler := range *hs {
		v = handler(v...)
	}
	return v
}
