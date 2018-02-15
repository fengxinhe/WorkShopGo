package view

import (
    "fmt"
    "html/template"
    "net/http"
    //"net/url"
    "path/filepath"
    "strings"
)

var (
    viewInfo View
    childTemplates     []string
	rootTemplate       string
	templateCollection = make(map[string]*template.Template)
	pluginCollection   = make(template.FuncMap)
	mutex              sync.RWMutex
    mutexPlugins sync.RWMutex

)

type View struct {
    Name string
    BaseURI string
    Folder string
    request *http.Request
	Vars map[string]interface{}
}


func Configure(vi View) {
    viewInfo = vi
}

func ReadConfig() View {
	return viewInfo
}

// LoadTemplates will set the root and child templates
func LoadTemplates(rootTemp string, childTemps []string) {
	rootTemplate = rootTemp
	childTemplates = childTemps
}

func (v *View) PrependBaseURI(s string) string {
	return v.BaseURI + s
}

// New returns a new view
func New(req *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})
	v.Vars["AuthLevel"] = "anon"

	v.BaseURI = viewInfo.BaseURI
	//v.Extension = viewInfo.Extension
	v.Folder = viewInfo.Folder
	v.Name = viewInfo.Name

	// Make sure BaseURI is available in the templates
	v.Vars["BaseURI"] = v.BaseURI

	// This is required for the view to access the request
	v.request = req

	// Get session
	//sess := session.Instance(v.request)

	// Set the AuthLevel to auth if the user is logged in
	// if sess.Values["id"] != nil {
	// 	v.Vars["AuthLevel"] = "auth"
	// }

	return v
}

// AssetTimePath returns a URL with the proper base uri and timestamp appended.
// Works for CSS and JS assets
// Determines if local or on the web
func (v *View) AssetTimePath(s string) (string, error) {
	if strings.HasPrefix(s, "//") {
		return s, nil
	}

	s = strings.TrimLeft(s, "/")
	abs, err := filepath.Abs(s)

	if err != nil {
		return "", err
	}

	time, err2 := FileTime(abs)
	if err2 != nil {
		return "", err2
	}

	return v.PrependBaseURI(s + "?" + time), nil
}

func FileTime(name string) (string, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return "", err
	}
	mtime := fi.ModTime().Unix()
	return fmt.Sprintf("%v", mtime), nil
}

func (v *View) Render(w http.ResponseWriter) {

	// Get the template collection from cache
	mutex.RLock()
	tc, ok := templateCollection[v.Name]
	mutex.RUnlock()

		// Cache the template collection
		mutex.Lock()
		templateCollection[v.Name] = templates
		mutex.Unlock()

	// Display the content to the screen
	err := tc.Funcs(pc).ExecuteTemplate(w, rootTemplate+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}
