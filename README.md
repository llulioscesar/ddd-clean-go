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

**Interfaz en la capa de dominio**

```go
package repository

import "dddcleango/internal/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
}
```

**Caso de uso en la capa de aplicacion**