package route

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/gorilla/context"
    "../static"
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
    router.GET("/home/firebug/goweb/static/*filepath",wrapHandler(http.HandlerFunc(static.Static)))
    router.GET("/", wrapHandler(http.HandlerFunc(controller.IndexGet)))
    router.GET("/classes", wrapHandler(http.HandlerFunc(controller.ClassGet)))

    router.GET("/classes/create",wrapHandler(http.HandlerFunc(controller.CreateClassGet)))

    router.POST("/classes/create",wrapHandler(http.HandlerFunc(controller.CreateClassPost)))

    router.GET("/featured/create",wrapHandler(http.HandlerFunc(controller.CreateProjectGet)))
    router.POST("/featured/create",wrapHandler(http.HandlerFunc(controller.CreateProjectPost)))

    router.GET("/featured/show",wrapHandler(http.HandlerFunc(controller.ShowProjectGet)))


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
