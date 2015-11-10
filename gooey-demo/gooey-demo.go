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

    e := new(gooey.Element)
    fmt.Println(e)

    gooey.Run(window)

    fmt.Println("Goodbye!")
}
