package main //pacote principal

import ( //pacotes importados

	"log"
	"net/http"
)

func main() { //função principal
	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", fs)
	log.Print("Listening on: 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)

	}
}
