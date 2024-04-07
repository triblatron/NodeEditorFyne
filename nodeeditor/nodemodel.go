package nodeeditor

const (
	CAT_NONE = iota
	CAT_SOURCE
	CAT_SINK
	CAT_CONDITION
	CAT_ACTION
	CAT_GROUP
	CAT_UNKNOWN
)

type NodeModel struct {
	Id       int64
	Name     string
	Category int
	Ports    []*PortModel
}

func NewNodeModel(id int64, name string, category int) *NodeModel {
	instance := new(NodeModel)
	instance.Id = id
	instance.Name = name
	instance.Category = category
	instance.Ports = make([]*PortModel, 0)
	return instance
}

func (self *NodeModel) AddPort(port *PortModel) {
	if port != nil {
		self.Ports = append(self.Ports, port)
	}
}
