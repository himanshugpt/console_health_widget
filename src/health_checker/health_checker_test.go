package health

import (
	"app"
	"fmt"
	"testing"
)

func TestAddApp(t *testing.T) {
	map_app := make(map[string]*app.Application)
	health_check := HealthCheck{map_app}
	application := app.Application{Name: "omega", Url: "somebadValue"}
	health_check.AddApp(&application)
	fmt.Println(map_app)
}
