package longpullling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductResponse struct {
	Message      string `json:"Message"`
	ProductID    string `json:"Product_id"`
	ProductName  string `json:"Product_name"`
	ProductPrice string `json:"Product_price"`
}

func Get_last_added_product(c *gin.Context) {
	for {
		response, err := http.Get("http://127.0.0.1:8080/products/get_last_added_product")

		if err != nil {
			fmt.Printf("Error en la petición http: ", err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Respuesta errónea: ", response.StatusCode)
			return
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			return
		}

		var product ProductResponse
		if err := json.Unmarshal(body, &product); err != nil {
			fmt.Println("Error al parsear JSON:", err)
			return
		}

		fmt.Println("Nuevo producto recibido:")
		fmt.Println("🔹", product.Message)
		fmt.Println("🆔 ID:", product.ProductID)
		fmt.Println("📦 Nombre:", product.ProductName)
		fmt.Println("💲 Precio:", product.ProductPrice)

		time.Sleep(10 * time.Second)
	}
}