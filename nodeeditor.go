package main

import (
	"NodeEditorFyne/nodeeditor"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Node Editor")

	//node := nodeeditor.NewNodeModel(0, "test", nodeeditor.CAT_NONE)
	input := nodeeditor.NewPortModel(0, "foo", nodeeditor.TYPE_DOUBLE, nodeeditor.DIR_IN)
	nodeView := nodeeditor.NewNodeView()
	nodeView.SetTitle("test")
	nodeView.AddPort(input)
	output := nodeeditor.NewPortModel(1, "bar", nodeeditor.TYPE_DOUBLE, nodeeditor.DIR_OUT)
	nodeView.AddPort(output)
	internal := nodeeditor.NewPortModel(2, "int", nodeeditor.TYPE_BOOL, nodeeditor.DIR_INTERNAL)
	nodeView.AddPort(internal)
	//data := binding.BindStringList(
	//	&[]string{"Item 1", "Item 2", "Item 3"},
	//)
	//
	//list := widget.NewListWithData(data,
	//	func() fyne.CanvasObject {
	//		return widget.NewLabel("template")
	//	},
	//	func(i binding.DataItem, o fyne.CanvasObject) {
	//		o.(*widget.Label).Bind(i.(binding.String))
	//	})
	//
	//add := widget.NewButton("Append", func() {
	//	val := fmt.Sprintf("Item %d", data.Length()+1)
	//	data.Append(val)
	//})
	myWindow.SetContent(nodeView.GetGrid())
	myWindow.ShowAndRun()
}
