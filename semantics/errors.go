package semantics

func withPrefix(text string) string {
	return "semantics: error: " + text
}

type NilNodeError struct{}

func (e *NilNodeError) Error() string {
	return withPrefix("nil node")
}

type UnexpectedNodeTypeError struct{}

func (e *UnexpectedNodeTypeError) Error() string {
	return withPrefix("unexpected node type")
}
