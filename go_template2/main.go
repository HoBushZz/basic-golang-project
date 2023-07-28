package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func f(w http.ResponseWriter, r *http.Request) {
	k := func(name string) (string, error) {
		return name + "测试", nil
	}
	t := template.New("f.html") /* 名字要和文件对应上 */
	t.Funcs(template.FuncMap{
		"test": k,
	})
	_, err := t.ParseFiles("./f.html")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "HeLinfeng"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed,err:%v", err)
		return
	}

}
