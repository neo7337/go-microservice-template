package handler

import (
	"github.com/neo7337/go-microservice-template/internal/service"
	"github.com/neo7337/go-microservice-template/pkg"
	"oss.nandlabs.io/golly/rest"
)

func UsersHandler(ctx rest.ServerContext) {
	users, err := service.GetUsers()
	if err != nil {
		pkg.ResponseJSON(ctx, 500, map[string]string{"error": err.Error()})
		return
	}
	pkg.ResponseJSON(ctx, 200, users)
}

func ConcurrencyDemoHandler(ctx rest.ServerContext) {
	// Simulate concurrent tasks: e.g., fetching data from two sources
	type result struct {
		Source string      `json:"source"`
		Data   interface{} `json:"data"`
	}
	results := make(chan result, 2)

	go func() {
		// Simulate a task (e.g., DB call)
		users, _ := service.GetUsers()
		results <- result{Source: "users", Data: users}
	}()

	go func() {
		// Simulate another task (e.g., external API call)
		info := map[string]string{"message": "Hello from goroutine!"}
		results <- result{Source: "info", Data: info}
	}()

	resp := make(map[string]interface{})
	for i := 0; i < 2; i++ {
		r := <-results
		resp[r.Source] = r.Data
	}

	pkg.ResponseJSON(ctx, 200, resp)
}
