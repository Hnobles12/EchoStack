package echostack

type EchoStack struct {
	Models []interface{}
}

func NewEchoStack(models ...interface{}) *EchoStack {
	return &EchoStack{Models: models}
}
