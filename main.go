package main

import (
	"fmt"
	"gopkg.in/qml.v1"
)

type File struct {
	Content string
}

type Html struct {
	html string
	baseUrl string
	unreachableUrl string
}

func (p *File) SetContent(content string) {
	fmt.Println("Old content is", p.Content)
	p.Content = content
	fmt.Println("New content is", p.Content)
}

func main() {
	err := qml.Run(run)
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	engine := qml.NewEngine()
	setValues(engine)
	component, err := engine.LoadFile("components/base.qml")
	if err != nil {
		return err
	}
	win := component.CreateWindow(nil)

	// source := win.ObjectByName("source")
	output := win.ObjectByName("output")

	output.Call("loadHtml", "<strong>hello!</strong>")

	win.Show()
	win.Wait()
	return nil
}

func setValues(engine *qml.Engine) {
	context := engine.Context()
	context.SetVar("file", &File{Content: "Hello world!"})
}
