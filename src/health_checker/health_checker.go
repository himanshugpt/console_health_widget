package health

import (
	"app"
)

type HealthCheck struct {
	Apps map[string]*app.Application
}

func (hc *HealthCheck) AddApp(app *app.Application) {
	hc.Apps[app.Name] = app
}

func (hc *HealthCheck) Run(ch chan *app.AppHealth){
	for {
		for _,v := range hc.Apps {
			 v.IsHealthy(ch)
		}
	}
}
