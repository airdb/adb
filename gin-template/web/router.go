package web

import (
	"io"
	"log"
	"net/http/httptest"
	"os"

	"{{ .GoModulePath }}/mocks"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Run() {
	defaultAddr := "0.0.0.0:8080"
	log.Println("web server starting, listen", defaultAddr)
	err := NewRouter().Run(defaultAddr)

	if err != nil {
		log.Println("error: ", err)
	}
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	pprof.Register(router)

	router.GET("/", DefaultRoot)

	v1API := router.Group("/apis/user/v1")
	v1API.Use()
	v1API.GET("/:user", ListUser)

	return router
}

func APIRequest(uri, method string, param io.Reader) *httptest.ResponseRecorder {
	// Change to the root directory for handler test case.
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = os.Chdir(wd)
		if err != nil {
			panic(err)
		}
	}()

	err = os.Chdir("../")
	if err != nil {
		panic(err)
	}

	db, err := mocks.SetUpMockDatabases()
	if err != nil {
		panic(err)
	}

	defer mocks.DestroyMockDatabases(db)

	req := httptest.NewRequest(method, uri, param)

	if method == "GET" {
		req.Header.Set("Content-Type", "application/json")
	} else if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	return w
}
