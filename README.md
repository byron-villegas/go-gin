# Go Gin

Proyecto base para aplicaciones Gin con ejemplos de configuración, testing y buenas prácticas.

## Tabla de Contenidos

- [Instalación](#instalación)
- [Ejecutar Aplicación](#ejecutar-aplicación)
- [Testing](#testing)
- [Tests de Aceptación](#tests-de-aceptación)
- [Tests de Rendimiento](#tests-de-rendimiento)
- [Swagger](#swagger)
- [Links de Referencia](#links-de-referencia)

## Instalación
### Instalar Go
Para instalar go debemos bajarlo e instalarlo de la siguiente pagina https://go.dev/doc/install

### Instalar Paquetes
Para instalar los paquetes debemos ejecutar el siguiente comando

```shell
go mod tidy
```

## Ejecutar Aplicación
Se debe ejecutar el siguiente comando

```shell
go run main.go
```

## Testing

### Ejecutar
Se debe ejecutar el siguiente comando

```shell
go test ./repository ./service ./controller ./routes ./config
```

### Ejecutar con Cobertura
Se debe ejecutar el siguiente comando

```shell
go test -cover ./repository ./service ./controller ./routes ./config
```

### Generar Reporte Cobertura Formato Consola
Se deben ejecutar los siguientes comandos

```shell
go test -coverprofile="coverage.out" ./repository ./service ./controller ./routes ./config
go tool cover -func="coverage.out"
```

### Generar Reporte Cobertura Formato HTML
Se deben ejecutar los siguientes comandos

```shell
go test -coverprofile="coverage.out" ./repository ./service ./controller ./routes ./config
go tool cover -html="coverage.out" -o coverage.html
```

## Tests de Aceptación
### Configuración
Se debe crear un archivo **acceptance-test/main_test.go** con el siguiente contenido

```go
package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	var outputFile *os.File
	var err error
	outputFile, err = os.Create("godog-report.json")

	if err != nil {
		t.Fatalf("failed to create report file: %v", err)
	}

	opts := godog.Options{
		Format:   "cucumber",
		Paths:    []string{"."},
		Output:   outputFile,
		TestingT: t,
		Strict:   true,
	}
	godog.TestSuite{
		Name:                "acceptance",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()
}
```

Esto nos permite definir la ruta de los features, el archivo y formato de salida de reporte de los tests de aceptación

### Ejecución
Se debe ejecutar el siguiente comando

```shell
go test ./acceptance-test
```

Al finalizar generara un reporte **godog-report.json**

## Tests de Rendimiento
### Configuración
Se debe crear un archivo **performance-test/vegeta/main_test.go** con el siguiente contenido

```go
package performance

import (
	"os"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

var apiHost string = "http://localhost:8080/api"

func init() {
	if os.Getenv("API_HOST") != "" {
		apiHost = os.Getenv("API_HOST")
	}
}

func TestProductEndpointPerformance(t *testing.T) {
	rate := vegeta.Rate{Freq: 10, Per: time.Second} // 10 requests por segundo
	duration := 5 * time.Second

	endpoint := "products"

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    apiHost + "/products",
	})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "GET /products") {
		metrics.Add(res)
	}
	metrics.Close()

	// Reporte en formato texto
	textFile, err := os.Create("vegeta-" + endpoint + "-report.txt")
	if err != nil {
		t.Fatalf("failed to create text report file: %v", err)
	}
	defer textFile.Close()

	textReporter := vegeta.NewTextReporter(&metrics)
	if err := textReporter.Report(textFile); err != nil {
		t.Errorf("failed to write vegeta text report: %v", err)
	}

	// Reporte en formato JSON
	jsonFile, err := os.Create("vegeta-" + endpoint + "-report.json")
	if err != nil {
		t.Fatalf("failed to create JSON report file: %v", err)
	}
	defer jsonFile.Close()

	jsonReporter := vegeta.NewJSONReporter(&metrics)
	if err := jsonReporter.Report(jsonFile); err != nil {
		t.Errorf("failed to write vegeta JSON report: %v", err)
	}

	if metrics.Success < 0.95 {
		t.Errorf("success rate too low: %.2f%%", metrics.Success*100)
	}
	t.Logf("99th percentile latency: %s", metrics.Latencies.P99)
}

```

Tenemos que crear una funcion (test) por endpoint a validar

### Ejecución
Se debe ejecutar el siguiente comando

```shell
go test -v ./performance-test/vegeta
```

Al finalizar generara dos reportes **vegeta-product-report.json** y **vegeta-product-report.txt**

## Swagger
### Instalación
Debemos instalar swaggo/swag con el siguiente comando

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

### Agregar Dependencias
Debemos agregar las siguientes dependencias

```shell
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

### Documentar Endpoints
Debemos documentar los endpoints con comentarios ejemplo

```go
// GetProducts godoc
// @Summary      Lista todos los productos
// @Description  Obtiene todos los productos disponibles
// @Tags         products
// @Produce      json
// @Success      200  {array}  models.Product
// @Router       /products [get]
func (ec *ProductController) GetProducts(c *gin.Context) {
    // ...
}
```

Swag generará la documentacion de swagger en base a esos comentarios


### Generar Archivo Swagger
Para generar el archivo swagger simplemente ejecutamos el siguiente comando

```shell
swag init
```

Esto nos generará los archivos de swagger en la carpeta **/docs** con la siguiente estructura

/docs
├── docs.go
├── swagger.json
└── swagger.yml

### Configurar Swagger UI
Para configurar Swagger UI simplemente agregamos el siguiente codigo al archivo **main.go**

```go
import (
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "go-gin/docs"
)

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

Quedando así

```go
package main

import (
	"go-gin/config"
	"go-gin/routes"

	_ "go-gin/docs" // docs generated by Swag

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Show banner with application information
	config.ShowBanner()

	// Create a server instance
	r := gin.Default()

	// Config CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{config.AppConfig.CORSHeaders},
		AllowCredentials: true,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Config server path
	routerGroup := r.Group(config.AppConfig.ServerPath)

	// Config routes for the server
	routes.SetupRoutes(routerGroup)

	// Start server with port
	r.Run(":" + config.AppConfig.ServerPort)
}
```

Cuando ejecutemos a la aplicacion debemos entrar a la pagina /swagger/index.html

## Links de Referencia
A continuación dejo links utilizados para realizar este proyecto

[Gin](https://go.dev/doc/tutorial/web-service-gin)

[Swagger](https://github.com/swaggo/swag)