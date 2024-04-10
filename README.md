A node editor GUI using Go and Fyne.

It uses MVC with micro-Controllers to control each NodeView.

Progress so far is initial attempts at:
* NodeModel
* PortModel
* NodeView
* NodeController

There are issues with layout, we are currently using a VBox of HBox arrangement with spacers, but the spacer on the output port pushes the circle too far to the right so that it is only half in the window.

This issue has been resolved, but setting the background colour for the title seems difficult.
The project has been discontinued in favour of using Qt from Go.
See NodeEditorGoQt
