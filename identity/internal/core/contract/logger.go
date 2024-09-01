package contract

type Logger interface {
	Info(context Context, msg string, fields *[]Field)
	Error(context Context, msg string, fields *[]Field)
	Close()
	NewMessage(context string, msg string, fields *[]Field) *Message
}

type Field struct {
	Key   string
	Value any
}

type Message struct {
	Project string
	Name    string
	ID      string
	ENV     string
	TZ      string
	Context Context
	Message string
	Fields  *[]Field `json:"fields,omitempty"`
}
