package app

import (
	"fmt"
	"testing"
)

func TestIsHealthy(t *testing.T) {
	app := Application{Name: "omega", Url: "http://employee-management.mongodb.cc/api/ping"}
	ch := make(chan AppHealth)
	go app.IsHealthy(ch)
	health := <-ch
	if health.err != "" {
		t.Error("test failed")
	}
	fmt.Println(health.Name)
}
