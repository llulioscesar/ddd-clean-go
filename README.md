# Proyecto en GO (DDD - Clean)

## Estructuras de carpetas

```text
├── LICENSE
├── README.md
├── build                                     # Packaging and Continuous Integration
│   ├── Dockerfile
│   └── init.sql
├── cmd                                       # Main Application
│   └── app
│       └── main.go
├── internal                                  # Private Codes
│   └── app
│       ├── adapter
│       │   ├── controller.go                 # Controller
│       │   ├── postgresql                    # Database
│       │   │   ├── conn.go
│       │   │   └── model                     # Database Model
│       │   │       ├── card.go
│       │   │       ├── cardBrand.go
│       │   │       ├── order.go
│       │   │       ├── parameter.go
│       │   │       ├── payment.go
│       │   │       └── person.go
│       │   ├── repository                    # Repository Implementation
│       │   │   ├── order.go
│       │   │   └── parameter.go
│       │   ├── service                       # Application Service Implementation
│       │   │   └── bitbank.go
│       │   └── view                          # Templates
│       │       └── index.tmpl
│       ├── application
│       │   ├── service                       # Application Service Interface
│       │   │   └── exchange.go
│       │   └── usecase                       # Usecase
│       │       ├── addNewCardAndEatCheese.go
│       │       ├── ohlc.go
│       │       ├── parameter.go
│       │       ├── ticker.go
│       │       └── ticker_test.go
│       └── domain
│           ├── factory                       # Factory
│           │   └── order.go
│           ├── order.go                      # Entity
│           ├── parameter.go
│           ├── parameter_test.go
│           ├── person.go
│           ├── repository                    # Repository Interface
│           │   ├── order.go
│           │   └── parameter.go
│           └── valueobject                   # ValueObject
│               ├── candlestick.go
│               ├── card.go
│               ├── cardbrand.go
│               ├── pair.go
│               ├── payment.go
│               ├── ticker.go
│               └── timeunit.go
└── testdata                                  # Test Data
    └── exchange_mock.go
```

## Capas

### 1. Dominio

En esta capa contiene todas las reglas de negocio del problema de la vida real, **entidades**.

### 2. Aplicacion

En esta capa, implementaremos la logica de negocio de la aplicacion, donde se encuentran los **casos de uso**.

### 3. Adaptador

En esta capa contiene los **controladores**, **gateways**, **presentadores**

### 4. Externa

En esta capa contiene las intefaces externas de comunicacion como **DB**, **UI**, **web**, **devices**.

![clean architecture](https://res.cloudinary.com/practicaldev/image/fetch/s--kunsEE2w--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/i/x5t4b6964ai2d9jtce02.jpg)

# Como pasar los limites de cada capa?

La [regla de depencias](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html#the-dependency-rule)
es uno de los principios de la arquitectura limpia.
> Esta regla dice que las dependencias del código fuente solo pueden apuntar hacia adentro. Nada en un círculo interior puede saber absolutamente nada sobre algo en un círculo exterior. - dsdasd

Por lo tanto, hay que realizar estos 4 pasos:

1. Definir interfaz,
2. Tome el argumento como interfaz y llame funciones de él.
3. Implementarlo.
4. Inyectar la dependencia.

Ejemplo del repositorio

**Estructura de archivos**

```text
.
└── internal
    └── app
        ├── adapter
        │   ├── controller.go    # 4. Dependency Injection
        │   └── repository
        │       └── parameter.go # 3. Implementation
        ├── application
        │   └── usecase
        │       └── parameter.go # 2. Interface Function Call
        └── domain
            ├── parameter.go
            └── repository
                └── parameter.go # 1. Interface
```

**1. Interfaz en la capa de dominio**

```go
package repository

import "dddcleango/internal/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
}
```

**2. Caso de uso en la capa de aplicacion**
```go
package usecase

import (
	"dddcleango/internal/app/domain"
	"dddcleango/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Parameter(r repository.IParameter) domain.Parameter {
	return r.Get()
}
```

**Implementación en la capa de adaptador:**
```go
package repository

import (
	"dddcleango/internal/app/adapter/postgresql"
	"dddcleango/internal/app/adapter/postgresql/model"
	"dddcleango/internal/app/domain"
)

// Parameter is the repository of domain.Parameter
type Parameter struct{}

// Get gets parameter
func (r Parameter) Get() domain.Parameter {
	db := postgresql.Connection()
	var param model.Parameter
	result := db.First(&param, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.Parameter{
		Funds: param.Funds,
		Btc:   param.Btc,
	}
}
```

**4. Inyección de dependencia en el controlador de la capa de adaptador:**
```go
package adapter

import (
	"net/http"

	"dddcleango/internal/app/adapter/repository"
	"dddcleango/internal/app/adapter/service"
	"dddcleango/internal/app/application/usecase"
	"dddcleango/internal/app/domain/valueobject"
	"github.com/gin-gonic/gin"
)

var (
	bitbank             = service.Bitbank{}
	parameterRepository = repository.Parameter{}
	orderRepository     = repository.Order{}
)

// Controller is a controller
type Controller struct{}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	// NOTICE: following path is from CURRENT directory, so please run Gin from root directory
	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/", ctrl.index)
	r.GET("/ticker", ctrl.ticker)
	r.GET("/candlestick", ctrl.candlestick)
	r.GET("/parameter", ctrl.parameter)
	r.GET("/order", ctrl.order)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Goilerplate",
	})
}

func (ctrl Controller) ticker(c *gin.Context) {
	pair := valueobject.BtcJpy
	ticker := usecase.Ticker(bitbank, pair) // Dependency Injection
	c.JSON(200, ticker)
}

func (ctrl Controller) candlestick(c *gin.Context) {
	args := usecase.OhlcArgs{
		E: bitbank, // Dependency Injection
		P: valueobject.BtcJpy,
		T: valueobject.OneMin,
	}
	candlestick := usecase.Ohlc(args)
	c.JSON(200, candlestick)
}

func (ctrl Controller) parameter(c *gin.Context) {
	parameter := usecase.Parameter(parameterRepository) // Dependency Injection
	c.JSON(200, parameter)
}

func (ctrl Controller) order(c *gin.Context) {
	order := usecase.AddNewCardAndEatCheese(orderRepository) // Dependency Injection
	c.JSON(200, order)
}
```