package handler

import (
	"fmt"
	_ "github.com/Max425/WB-Tech-level-0/docs"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
	"github.com/Max425/WB-Tech-level-0/pkg/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

func NewHandler(services *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", constants.Host)),
	))

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/customers/{uid}/orders", h.customerOrders).Methods("GET")
	apiRouter.HandleFunc("/orders/{uid}", h.order).Methods("GET")
	apiRouter.HandleFunc("/order", h.newOrder).Methods("POST")

	apiRouter.Use(
		h.panicRecoveryMiddleware,
		h.corsMiddleware,
	)

	return r
}
