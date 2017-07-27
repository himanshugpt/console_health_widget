package main

import (
	"app"
	"fmt"
	"github.com/gizak/termui"
	"health_checker"
)

func main() {
	ch := make(chan *app.AppHealth, 200)
	stop :=  make(chan struct{})
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		fmt.Println("Channel closed q captured")
		close(stop)
		termui.StopLoop()
	})

	map_app := make(map[string]*app.Application)
	health_check := health.HealthCheck{map_app}
	application := app.Application{Name: "Service Name", Url: "http://localhost:5000/api/ping"}
	health_check.AddApp(&application)

	//application2 := app.Application{Name: "omega2", Url: "http://employee-management.mongodb.cc/api/ping"}
	//health_check.AddApp(&application2)

	go health_check.Run(ch)
	go Display(ch, stop)
	termui.Loop()
}

func GetPar(app *app.AppHealth) *termui.Par {
	if app.Err == nil {
		par2 := termui.NewPar("\nStatus: Healthy \n" +  "Checked at: " + app.Timestamp.String() + "\nURL: " + app.URL)
		par2.Height = 6
		par2.Width = 60
		par2.Y = 4
		par2.BorderLabel = "Application: " + app.Name + " "
		par2.BorderFg = termui.ColorYellow
		return par2
	}else {
		par2 := termui.NewPar("\nStatus: Unhealthy \n" +  "Checked at: " + app.Timestamp.String() + "\nURL: " + app.URL)
		par2.Height = 6
		par2.Width = 60
		par2.Y = 4
		par2.BorderLabel = "Application: " + app.Name + " "
		par2.BorderFg = termui.ColorRed
		par2.TextFgColor = termui.ColorRed
		return par2
	}

}


func Display(ch chan *app.AppHealth, stop chan struct{}) {
	paraMap := make(map[string]*termui.Par)
	for {
		var app *app.AppHealth
		select {
		case <- stop:
			break
		case app = <- ch:
		}

		paraMap[app.Name] = GetPar(app)

		v := make([]*termui.Par, len(paraMap))
		idx := 0
		for  _, value := range paraMap {
			v[idx] = value
			idx++
		}

		termui.Render(v[0])

	}
}