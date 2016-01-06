package main

import(
	"cfg"
	"flag"
//	"fmt"
//	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("localhost", ":1234", "http service address") // Q=17, R=18

func main() {
	var config = cfg.NewConfig()

//	fmt.Printf("%+v", config)

	flag.Parse()
	for _, value := range config.Pages {
		http.Handle(value.Url, http.HandlerFunc(value.Cntr.Exec))
	}

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
