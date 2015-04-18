package interfaces

type Adder func(e Entity)

type Entity interface {
	Declaration() string
	Definition() string
}

type EntityLoader interface {
	Load(a EntityAdder, root string, filelist []string) error
}

type EntityAdder interface {
	Add(e Entity)
}