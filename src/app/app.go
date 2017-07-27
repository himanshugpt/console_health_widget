package app

import (
	"net/http"
	"time"
	"errors"
	"strconv"
)

type Application struct {
	Name, Url         string
	Timeout, Interval time.Duration
}

type AppHealth struct {
	// copies name and err all the time, better use pointers
	Name  string
	Err error
	Timestamp time.Time
	URL string
}

func (app *Application) IsHealthy(ch chan *AppHealth) {
		resp, err := http.Get(app.Url)
		//defer resp.Body.Close()
		if err != nil || resp.StatusCode >= 400 {
			code := ""
			if resp != nil && resp.StatusCode > 0 {
				code = strconv.Itoa(resp.StatusCode)
			}
			ch <- &AppHealth{app.Name, errors.New(code),time.Now(), app.Url}
		} else {
			ch <- &AppHealth{app.Name, nil, time.Now(), app.Url}
		}
		if(app.Interval != 0){
			time.Sleep(app.Interval)
		} else {
			time.Sleep(2 * time.Second)
		}
}
