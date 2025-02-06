package shortpulling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Get_quantity_of_users(c *gin.Context){
	for {
		response, err := http.Get("http://127.0.0.1:8080/users/comprobate")

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
				Users int `json:"Users"`
			}

			var result Response

			err = json.Unmarshal(body, &result)

			if err != nil {
				fmt.Println("Error al deserializar la respuesta:", err)
				response.Body.Close()
				time.Sleep(5 * time.Second)
				continue
			}

			fmt.Printf("Mensaje: %s, Usuarios: %d\n", result.Message, result.Users)
		}else {
			fmt.Printf("Respuesta inesperada del servidor: %d\n", response.StatusCode)
		}

		response.Body.Close()

		time.Sleep(5 * time.Second)
	}
}