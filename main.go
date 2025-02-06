package main

import (
	"fmt"
	"net/http"
	"time"

	longpullling "example.com/n/long_pullling"
	shortpulling "example.com/n/short_pulling"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/users", shortpulling.Get_quantity_of_users)
	r.GET("/update_price", shortpulling.UpdatePriceProducts)
	r.GET("/last_product", longpullling.Get_last_added_product)
	
	srv := &http.Server{
		Addr:         ":4000",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Minute,
		IdleTimeout:  1 * time.Hour,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Error: Server Main hasn't begin")
	}
}