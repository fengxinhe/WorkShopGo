package controller

import(
    "net/http"
    "strings"
    "fmt"
    "io"
    "time"
)


const STATIC_URL string = "/home/firebug/goweb/static/"
const STATIC_ROOT string = "/home/firebug/goweb/static/"


func StaticHandler(w http.ResponseWriter, r *http.Request) {
    if strings.HasSuffix(r.URL.Path, "/") {
		Error404(w, r)
		return
	}
    http.ServeFile(w, r, r.URL.Path[1:])
}

func Static(w http.ResponseWriter, req *http.Request) {
    static_file := req.URL.Path[len(STATIC_URL):]
    if len(static_file) != 0 {
        f, err := http.Dir(STATIC_ROOT).Open(static_file)
        if err == nil {
            content := io.ReadSeeker(f)
            http.ServeContent(w, req, static_file, time.Now(), content)
            return
        }
    }
    http.NotFound(w, req)
}


func GETStaticURL() string {
    return "/home/firebug/goweb/static/"
}

func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found 404")
}
