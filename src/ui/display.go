package ui

import (
	"app"
	"fmt"
	"github.com/gizak/termui"
	"time"
)

func Display(ch chan *app.AppHealth) {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	//termui.UseTheme("helloworld")
	for elem := range ch {
		fmt.Print("&&&&&&")
		if elem.Err == "Close" {
			break
			termui.StopLoop()
			close(ch)
			fmt.Println("Channel closed")
		}
		termui.Handle("/sys/kbd/q", func(termui.Event) {
			fmt.Println("Channel closed q captured")
			ch <- &app.AppHealth{"", "Close", time.Now()}
			termui.StopLoop()
		})
		par2 := termui.NewPar(elem.Name + " health check.\nChecked at " + elem.Timestamp.String())
		par2.Height = 5
		par2.Width = 37
		par2.Y = 4
		par2.BorderLabel = "Health Check"
		par2.BorderFg = termui.ColorYellow
		termui.Render(par2)
	}

	//termui.Loop()

	//par0 := termui.NewPar("Borderless Text")
	//par0.Height = 1
	//par0.Width = 20
	//par0.Y = 1
	//par0.Border = false
	//
	//par1 := termui.NewPar("你好，世界。")
	//par1.Height = 3
	//par1.Width = 17
	//par1.X = 20
	//par1.BorderLabel = "标签"
	//
	//par2 := termui.NewPar("Simple colored text\nwith label. It [can be](fg-red) multilined with \\n or [break automatically](fg-red,fg-bold)")
	//par2.Height = 5
	//par2.Width = 37
	//par2.Y = 4
	//par2.BorderLabel = "Multiline"
	//par2.BorderFg = termui.ColorYellow
	//
	//par3 := termui.NewPar("Long text with label and it is auto trimmed.")
	//par3.Height = 3
	//par3.Width = 37
	//par3.Y = 9
	//par3.BorderLabel = "Auto Trim"
	//
	//termui.Render(par0, par1, par2, par3)

}
