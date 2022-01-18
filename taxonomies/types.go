package taxonomies

type valueType int

const (
	typeInt valueType = iota
	typeFloat
	typeString
	typeEmpty
)

func (t valueType) String() string {
	return [...]string{"int", "float", "string", "empty"}[t]
}
