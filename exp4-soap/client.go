package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AspectRatioResponse struct {
	XMLName xml.Name   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    AspectBody `xml:"Body"`
}

type AspectBody struct {
	AspectRatioResponse AspectRatioResult `xml:"CalculateAspectRatioResponse"`
}

type AspectRatioResult struct {
	AspectRatioResult float64 `xml:"CalculateAspectRatioResult"`
}

func main() {
	url := "http://www.dneonline.com/calculator.asmx"

	requestPayload := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
						<soap:Body>
							<tem:CalculateAspectRatio>
								<tem:x>1920</tem:x>
								<tem:y>1080</tem:y>
								<tem:mdc>10</tem:mdc>
							</tem:CalculateAspectRatio>
						</soap:Body>
					</soap:Envelope>`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestPayload)))
	if err != nil {
		log.Fatal("Erro criando a requisição:", err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://tempuri.org/CalculateAspectRatio")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro durante a requisição:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro lendo a resposta:", err)
	}

	fmt.Println("Resposta SOAP:")
	fmt.Println(string(body))

	var response AspectRatioResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		log.Fatal("Erro durante Unmarshal:", err)
	}

	aspectRatio := response.Body.AspectRatioResponse.AspectRatioResult
	fmt.Println("Aspect Ratio:", aspectRatio)
}
