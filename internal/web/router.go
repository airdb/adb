package web

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/airdb/sailor/faas"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
)

func Run() {
	r := chi.NewRouter()
	r.Get("/", HandleRoot)

	faas.RunChi(r)
}

var ctx = context.Background()

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	redisURL := os.Getenv("REDIS_URL")

	opt, _ := redis.ParseURL(redisURL)
	client := redis.NewClient(opt)

	_, err := client.Get(ctx, "counter").Int64()
	if err != nil {
		client.Set(ctx, "counter", 1, 0)
		// client.Incr(ctx, "counter")
	}

	client.Incr(ctx, "counter")

	// client.Set(ctx, "name", "adb", 0)
	val, _ := client.Get(ctx, "counter").Int64()
	msg := fmt.Sprintf("welcome counter: %d", val)
	w.Write([]byte(msg))
}
