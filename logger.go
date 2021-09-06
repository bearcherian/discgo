package discgo

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}

var _ Logger = &noopLogger{}

type noopLogger struct{}

func (l *noopLogger) Debug(...interface{})          {}
func (l *noopLogger) Debugf(string, ...interface{}) {}
func (l *noopLogger) Info(...interface{})           {}
func (l *noopLogger) Infof(string, ...interface{})  {}
func (l *noopLogger) Error(...interface{})          {}
func (l *noopLogger) Errorf(string, ...interface{}) {}
