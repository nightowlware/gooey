// Gooey is a GUI system for Go programs
package gooey

import (
    "fmt"
    "runtime"
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
    // GLFW examples do this because some glfw calls (and probably all gl calls) need
    // to run on the main thread.
    runtime.LockOSThread()
}

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

    width int
    height int
}

func (w *Window) Handle() *glfw.Window {
    return w.handle
}

//TODO: how to handle multiple windows? Do we even want to allow that?
func CreateWindow(width int, height int, title string) *Window {
    handle, err := glfw.CreateWindow(width, height, title, nil, nil)
    if err != nil {
        panic(err)
    }

    w := new(Window)
    w.handle = handle
    w.width = width
    w.height = height

    handle.MakeContextCurrent()

    err = gl.Init();
    if err != nil {
        panic(err)
    }

    return w
}

func Run(window *Window) {
    setupScene(window.width, window.height)
    for !window.handle.ShouldClose() {
        drawScene()
        window.handle.SwapBuffers()
        glfw.PollEvents()
    }
}

func setupScene(width int, height int) {
    gl.Disable(gl.DEPTH_TEST)
    gl.Disable(gl.LIGHTING)

    gl.ClearColor(0.5, 0.5, 0.5, 0.0)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.Ortho(0, float64(width), 0, float64(height), -1, 1)
    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()
}

func drawScene() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    gl.LineWidth(1)

    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()
    gl.Translatef(0, 0, 0.0)
    //gl.Rotatef(rotationX, 1, 0, 0)
    //gl.Rotatef(rotationY, 0, 1, 0)

    //rotationX += 0.5
    //rotationY += 0.5

    //gl.BindTexture(gl.TEXTURE_2D, texture)

    gl.Color4f(1, 0, 0, 1)

    gl.Begin(gl.QUADS)

    gl.Vertex3f(50, 50, 0)
    gl.Vertex3f(50, 100, 0)
    gl.Vertex3f(100, 100, 0)
    gl.Vertex3f(100, 50, 0)

    gl.End()


}
