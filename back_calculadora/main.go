package main;

import(
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
);

type historial struct {
	Fecha string `json:"fecha"`
	Operacion string `json:"operacion"`
	Resultado int `json:"resultado"`
}

var historialOperaciones = []historial{}

func getHistorial(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, historialOperaciones);
}

func main () {
	router := gin.Default();
	router.GET("/historial", getHistorial);
	router.Run("localhost:4000")
}