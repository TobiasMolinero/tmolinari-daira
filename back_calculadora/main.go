package main;

import(
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
);

type Historial struct {
	Id_operacion int `json:"id_operacion"`
	Fecha string `json:"fecha"`
	Operacion string `json:"operacion"`
	Resultado string `json:"resultado"`
}

var historialOperaciones = []Historial{}

func enableCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*");
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true");
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE");
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization");
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200);
		return
	}
	c.Next()
}

func getHistorial(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, historialOperaciones);
}

func guardarOperacion(c *gin.Context){
	var nuevaOperacion Historial;
	c.BindJSON(&nuevaOperacion);
	historialOperaciones = append(historialOperaciones, nuevaOperacion);
	c.IndentedJSON(http.StatusCreated, historialOperaciones);
}

func main () {
	router := gin.Default();
	router.Use(enableCors);
	router.GET("/historial", getHistorial);
	router.POST("/historial/add", guardarOperacion);
	router.Run("localhost:4000")
}