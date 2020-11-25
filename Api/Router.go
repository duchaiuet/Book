package Api

import (
	"Book/Controller"
	"Book/Database"
	_ "Book/docs"
	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"strings"
)

const (
	basePath string = "/api/v1/book"
)

func Router() http.Handler {
	r := chi.NewRouter()

	bookController := Controller.NewBookController()
	categoryController := Controller.NewCategoryController()
	roleController := Controller.NewRoleController()
	userController := Controller.NewUserController()
	//authorization := middleware.NewAuthorizeMiddleware()

	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(Database.HttpSwagger),
	))

	r.Route(basePath, func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))
		r.Route("/", func(r chi.Router) {
			r.Route("/books", func(r chi.Router) {
				r.Post("/", bookController.CreateBook)
				r.Get("/", bookController.GetListBook)
				r.Get("/{id}", bookController.GetBookById)
				r.Put("/activate/{code}", bookController.ActiveBook)
				r.Put("/deactivate/{code}", bookController.DeActiveBook)
			})
			r.Route("/categories", func(r chi.Router) {
				r.Post("/", categoryController.CreateCategory)
				r.Get("/", categoryController.GetListCategory)
				r.Get("/{code}", categoryController.GetCategoryById)
				r.Put("/{code}", categoryController.UpdateCategory)
			})
			r.Route("/roles", func(r chi.Router) {
				r.Post("/", roleController.CreateRole)
				r.Get("/", roleController.GetList)
				r.Get("/{id}", roleController.GetById)
				r.Put("/{id}", roleController.Update)
				r.Put("/active/{id}", roleController.Active)
				r.Put("/deactivate/{id}", roleController.DeActive)
			})
			r.Route("/users", func(r chi.Router) {
				r.Post("/", userController.Create)
				r.Get("/", userController.GetList)
				r.Get("/{id}", userController.GetById)
				r.Put("/{id}", userController.Update)
				r.Put("/activate/{id}", userController.Active)
				r.Put("/deactivate/{id}", userController.DeActive)

			})
		})

	})

	return r
}

func HandleHttpServer(port string) {
	Database.InfoLog.Printf("server  %s", port)
	Database.ErrLog.Printf("swagger : %s", strings.Replace(Database.HttpSwagger, "doc.json", "index.html", 1))
	err := http.ListenAndServe(port, Router())
	if err != nil {
		Database.ErrLog.Print(err)
	}
}
