// GooeyDemo is a demo of Gooey, a GUI system for Go.
package main

import (
    "fmt"
    "runtime"

    "github.com/nightowlware/gooey"

    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.1/glfw"
)

var screenwidth int = 640
var screenheight int = 480

// Needed for GLFW:
func init() {
    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

func main() {
    fmt.Println("Gooey Demo")

    gooey.Init()
    defer gooey.Terminate()

    w := gooey.CreateWindow(screenwidth, screenheight, "Gooey Demo v0.1")

    window := w.Handle()

    window.MakeContextCurrent()

    err := gl.Init();
    if err != nil {
        panic(err)
    }


    e := new(gooey.Element)
    fmt.Println(e)




    setupScene()
    for !window.ShouldClose() {
        drawScene()
        window.SwapBuffers()
        glfw.PollEvents()
    }


    fmt.Println("Goodbye!")
}


func setupScene() {
    gl.Disable(gl.DEPTH_TEST)
    gl.Disable(gl.LIGHTING)

    gl.ClearColor(0.5, 0.5, 0.5, 0.0)

    gl.MatrixMode(gl.PROJECTION)
    gl.LoadIdentity()
    gl.Ortho(0, float64(screenwidth), 0, float64(screenheight), -1, 1)
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
