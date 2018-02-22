package route

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/gorilla/context"
    "../"
)


func LoadHTTPS() http.Handler {
    return middleware(routes())
}

func LoadHTTP() http.Handler {
    return middleware(routes())
}

func routes() *httprouter.Router {
    router := httprouter.New()
    //router.ServeFiles("/static/*filepath", http.Dir("/home/firebug/goweb/"))
    router.GET("/home/firebug/goweb/static/*filepath",wrapHandler(http.HandlerFunc(controller.Static)))
    router.GET("/", wrapHandler(http.HandlerFunc(controller.IndexGet)))
    router.GEt("/classes", wrapHandler(http.HandlerFunc(controller.ClassGet)))

    return router
}

func middleware(h http.Handler) http.Handler {
    //h = logrequest.Handler(h)
    return h
}

func wrapHandler(h http.Handler) httprouter.Handle{
    return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context.Set(r, "params", p)
		h.ServeHTTP(w, r)
}
}