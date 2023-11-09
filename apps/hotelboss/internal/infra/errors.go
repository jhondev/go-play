package infra

type Errors struct {
	logger *Logger
}

func NewErrors(l *Logger) *Errors {
	return &Errors{logger: l}
}

func (e *Errors) Error(code string, err string) map[string]string {
	return map[string]string{code: err}
}
