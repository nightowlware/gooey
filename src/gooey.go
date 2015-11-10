// Gooey is a GUI system for Go programs
package gooey

import (
"fmt"
"github.com/go-gl/glfw/v3.1/glfw"
)

type Element struct {
    name string
    x float64
    y float64
    width float64
    height float64
}

// Define custom interfaces for Element types:
type Named interface {
    Name() string
    SetName(name string)
}

type Renderable interface {
    Render()
}

// Implement Stringer interface:
func (w *Element) String() string {
    return fmt.Sprintf("Element [%s] of size (%f,%f)", w.name, w.width, w.height)
}

// Implement Named interface:
func (w *Element) Name() string {
    return w.name
}
func (w *Element) SetName(name string) {
    w.name = name
}

// Implement Renderable interface:
func (w *Element) Render() {
    fmt.Println("RENDER: ", w)
}





func Init() {
    err := glfw.Init()
    if err != nil {
        panic(err)
    }
}

func Terminate() {
    glfw.Terminate()
}

type Window struct {
    handle *glfw.Window
}

func (w *Window) Handle() *glfw.Window {
    return w.handle
}

func CreateWindow(width int, height int, title string) *Window {
    handle, err := glfw.CreateWindow(width, height, title, nil, nil)
    if err != nil {
        panic(err)
    }

    w := new(Window)
    w.handle = handle
    return w
}


