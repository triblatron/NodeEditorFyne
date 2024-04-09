package nodeeditor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
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
	instance.grid.Add(container.NewCenter(instance.title))

	return instance
}

func (self *NodeView) GetObj() fyne.CanvasObject {
	return self.obj
}

func (self *NodeView) GetGrid() *fyne.Container {
	return self.grid
}

func (self *NodeView) SetTitle(title string) {
	self.title.SetText(title)
	self.title.Alignment = fyne.TextAlignCenter
}

func (self *NodeView) AddPort(port *PortModel) {
	parent := container.NewHBox()
	switch port.Direction {
	case DIR_IN:
		dirChild := canvas.NewCircle(colornames.Red)
		dirChild.Resize(fyne.NewSize(10, 10))
		nameChild := widget.NewLabel(port.Name)
		nameChild.Alignment = fyne.TextAlignCenter
		noLayout := container.NewBorder(nil, nil, container.NewWithoutLayout(dirChild), nil, container.NewPadded(nameChild))
		parent.Add(noLayout)
		//		parent.Add(nameChild)
		//parent.Add(layout.NewSpacer())
	case DIR_OUT:
		dirChild := canvas.NewCircle(colornames.Green)
		dirChild.Resize(fyne.NewSize(10, 10))
		noLayout := container.NewWithoutLayout(dirChild)

		nameChild := widget.NewLabel(port.Name)
		nameChild.Alignment = fyne.TextAlignTrailing
		// Create a container with padding to ensure fixed distance from window edge
		border := container.NewBorder(nil, nil, noLayout, nil)
		parent.Add(layout.NewSpacer())
		parent.Add(nameChild)
		parent.Add(container.NewPadded(border))

	case DIR_INTERNAL:
		nameChild := widget.NewLabel(port.Name)
		parent.Add(layout.NewSpacer())
		parent.Add(nameChild)
		parent.Add(layout.NewSpacer())
	default:
		// Do nothing.
	}
	self.grid.Add(parent)
}
