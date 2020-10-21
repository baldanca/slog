package slog

/* type (
	// StackService contract
	StackService interface {
		GetAll()
	}
	// data array type
	data [][]byte
	// MultiWriter array writer type
	MultiWriter []io.Writer
	// Stack struct type
	Stack struct {
		data      *data
		logger    *Logger
		customOut *MultiWriter
		debugOut  *MultiWriter
		errOut    *MultiWriter
		fatalOut  *MultiWriter
		infoOut   *MultiWriter
		panicOut  *MultiWriter
		warnOut   *MultiWriter
	}
)

// NewStack function
func NewStack(l *Logger) *Stack {
	data := new(data)
	s := &Stack{
		data:      data,
		logger:    l,
		customOut: NewMultiWriter(l.custom.Writer(), data),
		debugOut:  NewMultiWriter(l.debug.Writer(), data),
		errOut:    NewMultiWriter(l.err.Writer(), data),
		fatalOut:  NewMultiWriter(l.fatal.Writer(), data),
		infoOut:   NewMultiWriter(l.info.Writer(), data),
		panicOut:  NewMultiWriter(l.panic.Writer(), data),
		warnOut:   NewMultiWriter(l.warn.Writer(), data),
	}
	l.custom.SetOutput(s.customOut)
	l.debug.SetOutput(s.debugOut)
	l.err.SetOutput(s.errOut)
	l.fatal.SetOutput(s.fatalOut)
	l.info.SetOutput(s.infoOut)
	l.panic.SetOutput(s.panicOut)
	l.warn.SetOutput(s.warnOut)
	return s
}

// Close stack
func (s *Stack) Close() {
	s.logger.custom.SetOutput((*s.customOut)[0])
	s.logger.debug.SetOutput((*s.debugOut)[0])
	s.logger.err.SetOutput((*s.errOut)[0])
	s.logger.fatal.SetOutput((*s.fatalOut)[0])
	s.logger.info.SetOutput((*s.infoOut)[0])
	s.logger.panic.SetOutput((*s.panicOut)[0])
	s.logger.warn.SetOutput((*s.warnOut)[0])
}

// GetAll stack function
func (s *Stack) GetAll() [][]byte {
	return *s.data
}

// GetAllString stack function
func (s *Stack) GetAllString() []string {
	stackString := []string{}
	for _, value := range *s.data {
		stackString = append(stackString, string(value))
	}
	return stackString
}

// PrintAll stack function
func (s *Stack) PrintAll() {
	for _, value := range *s.data {
		fmt.Printf(string(value))
	}
}

// Save stack
func (s *Stack) Save(w io.Writer) {
	bytes, err := jsoniter.Marshal(s.data)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

// Write a log to stack
func (d *data) Write(p []byte) (int, error) {
	*d = append(*d, p)
	return len(p), nil
}

// NewMultiWriter function
func NewMultiWriter(w ...io.Writer) *MultiWriter {
	t := new(MultiWriter)
	*t = w
	return t
}

// Write on multiWriter
func (t *MultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range *t {
		n, err = w.Write(p)
		if err != nil {
			return
		}
	}
	return len(p), nil
}
*/
