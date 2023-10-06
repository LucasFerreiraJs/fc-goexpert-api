package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/lucasferreirajs/fc-goexpert/apis/configs"
	_ "github.com/lucasferreirajs/fc-goexpert/apis/docs"
	"github.com/lucasferreirajs/fc-goexpert/apis/internal/entity"
	"github.com/lucasferreirajs/fc-goexpert/apis/internal/infra/database"
	"github.com/lucasferreirajs/fc-goexpert/apis/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// httpSwagger "github.com/swaggo/http/swagger"

// @title           Go expert API EXAMPLE
// @version         1.0
// @description     product api with authentication.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Lucas Ferreira
// @contact.url    github.com/lucasferreirajs
// @contact.email  casferreiraa@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)

	// * before
	// userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(LogRequest)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProductById)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)
	// http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe("127.0.0.1:8000", r)

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
