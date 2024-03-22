package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	type AddRequest struct {
		XMLName xml.Name `xml:"http://tempuri.org/Add Add"`
		IntA    int      `xml:"intA"`
		IntB    int      `xml:"intB"`
	}

	requestBody := &AddRequest{IntA: 10, IntB: 20}
	requestBodyBytes, err := xml.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Erro ao criar solicitação XML: %v", err)
	}
	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Fatalf("Erro ao criar solicitação HTTP: %v", err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://tempuri.org/Add"`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer solicitação HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	fmt.Println(string(body))
}
