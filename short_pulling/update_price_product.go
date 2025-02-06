package shortpulling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdatePriceProducts(c *gin.Context) {
	search_id := 1
	price := 0
	for {
		price++
		data := []byte(`{"Price": ` + strconv.Itoa(price) + `}`,)
		request, err := http.NewRequest(
			"PUT", 
			"http://127.0.0.1:8080/products/update_price/" + strconv.Itoa(search_id), 
			bytes.NewBuffer(data),
		)

		request.Header.Set("Content-Type", "application/json")

		client := http.Client{}

		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error al hacer la solicitud HTTP:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if response.StatusCode == http.StatusAccepted {
			type ResponseMap map[string]string
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Error al leer el cuerpo de la respuesta:", err)
				response.Body.Close()
				time.Sleep(5 * time.Second)
				continue
			}

			type Response struct {
				Message string `json:"Message"`
			}

			var result Response

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al deserializar la respuesta:", err)
				response.Body.Close()
				time.Sleep(5 * time.Second)
				continue
			}
			fmt.Printf("Mensaje: %s", result.Message)
		}else {
			fmt.Printf("Respuesta inesperada del servidor: %d\n", response.StatusCode)
		}

		response.Body.Close()
		time.Sleep(5 * time.Second)
		search_id++
	}
}