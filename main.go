package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type myHandler struct {
	Contador int
}

// Declarando o método que pertence a struct myHandler
// Estado utilizando ponteiro
func (mh *myHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request){

	if req.URL.Path == "/hello" {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello word"))
		return
	}

	if req.Method == "GET" {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(intToString(mh.Contador)))
	} else if req.Method == "POST" {
		mh.Contador++
	} else if req.Method == "DELETE" {
		mh.Contador = 0
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("Requisição recebido, metodo: %s\n", req.Method)
	writer.Header().Add("Meu-Header", "Meu valor de Header")
	writer.WriteHeader(220)

	writer.Write([]byte("Hello World da API!"))
}

func intToString(a int) string {
	return fmt.Sprintf("%d", a)
}

func stringToInt(a string) int {
	valorNumerico, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return valorNumerico
}



func main() {
	//type Hearder map[string][]string

	//headerMap := Hearder{}
	//headerMap["Meu-Header"] = []string{"Um valor xpto", "um outro valor"}

	//novoMap := map[string]int{}
	//novoMap["xpto"] = 1
	//novoMap["abcd"] = -1

	fmt.Println("Hello World")
	//fmt.Println(intToString(36))
	//fmt.Println(stringToInt("36"))

	handler := myHandler{}
	http.ListenAndServe(":8080", &handler)

}


// Precisa que o app escute numa porta TCP
// precisa declarar funcoes que tratem essa requisições HTTP