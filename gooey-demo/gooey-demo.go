// GooeyDemo is a demo of Gooey, a GUI system for Go.
package main

import (
    "fmt"
    "github.com/nightowlware/gooey"
)

func main() {
    fmt.Println("Gooey Demo")

    gooey.Init()
    defer gooey.Terminate()

    window := gooey.CreateWindow(640, 480, "Gooey Demo v0.1")

    e := window.NewElement()
    e.SetName("button1")
    e.SetPosition(100, 40)
    e.SetColor(gooey.Color{1.0, 0.0, 0.0, 1.0})

    e = window.NewElement()
    e.SetName("checkbox")
    e.SetPosition(300, 120)
    e.SetColor(gooey.Color{0.0, 1.0, 0.0, 1.0})

    e = window.NewElement()
    e.SetName("slider")
    e.SetPosition(20, 300)
    e.SetColor(gooey.Color{0.0, 0.0, 1.0, 1.0})

    gooey.Run(window)

    fmt.Println("Goodbye!")
}
