package main

import (
	"fmt"
	"gopkg.in/qml.v1"
	"github.com/russross/blackfriday"
)

func main() {
	err := qml.Run(runApp)
	if err != nil {
		fmt.Println(err)
	}
}

func runApp() error {
	engine := qml.NewEngine()
	winComponent, err := engine.LoadFile("components/base.qml")

	if err != nil {
		return err
	}

	win := winComponent.CreateWindow(nil)
	source := win.ObjectByName("source")
	target := win.ObjectByName("target")

	compile(source, target)
	watch(source, target)

	win.Show()
	win.Wait()

	return nil
}

func watch(source qml.Object, target qml.Object) {
	source.On("textChanged", func() {
		compile(source, target)
    })
}

func compile(source qml.Object, target qml.Object) {
	input := source.String("text")
	output := blackfriday.MarkdownBasic([]byte(input))
	target.Call("loadHtml", string(output))
}