package handler

import (
	"encoding/json"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// @Summary get customer orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param uid path string true "customer UID"
// @Success 200 {object} []core.Order
// @Failure 500 {object} string
// @Router /api/v1/customers/{uid}/orders [get]
func (h *Handler) customerOrders(w http.ResponseWriter, r *http.Request) {
	uid, has := mux.Vars(r)["uid"]
	if !has {
		dto.SendData(w, "invalid params", http.StatusBadRequest)
		return
	}
	orders, err := h.services.Order.GetCustomerOrders(r.Context(), uid)
	if err != nil {
		dto.SendData(w, "some error", http.StatusInternalServerError)
		return
	}

	dto.SendData(w, orders, http.StatusOK)
}

// @Summary create Order
// @Tags order
// @Accept  json
// @Produce  json
// @Param input body core.Order true "New order"
// @Success 200 {object} core.Order
// @Failure 400,500 {object} string
// @Router /api/v1/order [post]
func (h *Handler) newOrder(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	var order core.Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	createOrder, err := h.services.Order.CreateOrder(r.Context(), &order)
	if err != nil {
		dto.SendData(w, "some error", http.StatusInternalServerError)
		return
	}

	response := dto.CreateResponse{ID: createOrder}
	dto.SendData(w, response, http.StatusOK)
}
