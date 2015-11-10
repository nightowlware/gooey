// GooeyDemo is a demo of Gooey, a GUI system for Go.
package main

import (
    "fmt"
    "runtime"
    "github.com/nightowlware/gooey"
)

// Needed for GLFW: TODO: can I move this into gooey package? not sure how...
func init() {
    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

func main() {
    fmt.Println("Gooey Demo")

    gooey.Init()
    defer gooey.Terminate()

    window := gooey.CreateWindow(640, 480, "Gooey Demo v0.1")

    e := new(gooey.Element)
    fmt.Println(e)

    gooey.Run(window)

    fmt.Println("Goodbye!")
}
