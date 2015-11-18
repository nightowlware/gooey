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

//------------------------------------------------------
// COLOR
//------------------------------------------------------
type Color struct {
    R float32
    G float32
    B float32
    A float32
}


//------------------------------------------------------
// ELEMENT
//------------------------------------------------------
type Element struct {
    name string
    x float32
    y float32
    width float32
    height float32
    color Color
}

func (e *Element) String() string {
    return fmt.Sprintf("Element [%s] of size (%f,%f)", e.name, e.width, e.height)
}

func (e *Element) Name() string {
    return e.name
}

func (e *Element) SetName(name string) {
    e.name = name
}

func (e *Element) SetPosition(x float32, y float32) {
    e.x = x
    e.y = y
}

func (e *Element) SetColor(color Color) {
    e.color = color
}

func (e *Element) Render() {

    gl.Color4f(e.color.R, e.color.G, e.color.B, e.color.A)
    gl.Begin(gl.QUADS)

    gl.Vertex3f(e.x, e.y, 0)
    gl.Vertex3f(e.x, e.y + e.height, 0)
    gl.Vertex3f(e.x + e.width, e.y + e.height, 0)
    gl.Vertex3f(e.x + e.width, e.y, 0)

    gl.End()
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







//------------------------------------------------------
// GENERAL
//------------------------------------------------------
type Window struct {
    handle *glfw.Window

    width int
    height int

    elements []*Element;
}

func (w *Window) NewElement() *Element {
    e := &Element{}
    e.width = 50
    e.height = 50

    e.color = Color{0.0, 0.0, 0.0, 1.0}

    w.elements = append(w.elements, e)
    return e
}

func (w *Window) Render() {
    for i := 0; i < len(w.elements); i++ {
        e := w.elements[i]
        e.Render()
    }
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
        window.Render()
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




}
