package nodeeditor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NodeView struct {
	grid  *fyne.Container
	title *widget.Label
	obj   fyne.CanvasObject
}

func NewNodeView() *NodeView {
	instance := new(NodeView)
	instance.obj = new(fyne.Container)
	instance.grid = container.New(layout.NewVBoxLayout())
	instance.title = widget.NewLabel("")

	instance.grid.Add(instance.title)
	return instance
}
