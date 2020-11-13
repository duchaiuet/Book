package Api

import (
	"Book/Controller"
	"Book/Database"
	_ "Book/docs"
	"github.com/go-chi/chi"
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
	//authorization := middleware.NewAuthorizeMiddleware()

	r.Get(basePath+"/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(Database.HttpSwagger),
	))
	r.Route(basePath, func(r chi.Router) {

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
