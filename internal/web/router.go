package web

import (
	"context"
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

	client.Set(ctx, "name", "adb", 0)
	val := client.Get(ctx, "name").Val()
	w.Write([]byte("welcome " + val))
}
