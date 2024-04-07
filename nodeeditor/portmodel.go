package nodeeditor

const (
	TYPE_INT = iota
	TYPE_DOUBLE
	TYPE_STRING
	TYPE_BOOL
)

const (
	DIR_UNKNOWN = iota
	DIR_IN
	DIR_OUT
	DIR_INTERNAL
)

type PortModel struct {
	Id        int64
	Name      string
	Type      int
	Direction int
	Value     interface{}
}

func NewPortModel(id int64, name string, portType int, dir int) *PortModel {
	instance := new(PortModel)
	instance.Id = id
	instance.Name = name
	instance.Type = portType
	instance.Direction = dir

	return instance
}
