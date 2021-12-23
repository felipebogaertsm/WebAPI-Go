// Author: Felipe Bogaerts de Mattos
// email: felipebogaerts@gmail.com
// License: MIT

package main

import (
	. "popt/apps/data"

	"github.com/gin-gonic/gin"
)

const portNumber string = "8080"
const inProduction = false

// main is the main function
func main() {
	router := gin.Default()

	router.GET("/data/read", DataReadCSVController)
	router.GET("/data/mean", DataMeanController)
	router.GET("/data/covariance", DataCovarianceController)
	router.GET("/data/markowitz", MarkowitzOutputController)

	router.Run(":" + portNumber)
}
