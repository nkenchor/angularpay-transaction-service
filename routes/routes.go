package routes

import (
	"net/http"
	"transaction/interfaces"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(appPort, hostAddress string, transaction interfaces.TransactionInterface) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: customLogger, NoColor: true})
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/transaction", transactionEndpoint(transaction))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+":"+appPort+"/swagger/doc.json"),
	))

	return router
}

func transactionEndpoint(transaction interfaces.TransactionInterface) http.Handler {
	r := chi.NewRouter()
	r.Post("/requests", transaction.CreateTransaction)
	r.Put("/requests/{reference}/status", transaction.UpdateTransaction)
	r.Delete("/transaction/{reference}", transaction.GetTransactionByRef)
	r.Get("/requests?date={date}&type={type}&service={reference}", transaction.GetTransactionListByCriteria)
	return r
}
