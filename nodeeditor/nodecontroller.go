package nodeeditor

type NodeController struct {
	node *NodeModel
	view *NodeView
}

func NewNodeController(model *NodeModel) *NodeController {
	instance := new(NodeController)
	instance.view = NewNodeView()

	return instance
}
