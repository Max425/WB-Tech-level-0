package handler

import (
	"errors"
	"github.com/Max425/WB-Tech-level-0/pkg/constants"
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
// @Success 200 {object} []dto.Order
// @Failure 500 {object} string
// @Router /api/v1/customers/{uid}/orders [get]
func (h *Handler) customerOrders(w http.ResponseWriter, r *http.Request) {
	UID, has := mux.Vars(r)["uid"]
	if !has {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}
	orders, err := h.services.Order.GetCustomerOrders(UID)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "unknown error")
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, orders)
}

// @Summary get order by UID
// @Tags order
// @Accept  json
// @Produce  json
// @Param uid path string true "order UID"
// @Success 200 {object} dto.Order
// @Failure 500 {object} string
// @Router /api/v1/orders/{uid} [get]
func (h *Handler) order(w http.ResponseWriter, r *http.Request) {
	UID, has := mux.Vars(r)["uid"]
	if !has {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}
	order, err := h.services.Order.GetOrderByUID(r.Context(), UID)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "unknown error")
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, order)
}

// @Summary create Order
// @Tags order
// @Accept  json
// @Produce  json
// @Param input body dto.Order true "New order"
// @Success 200 {object} dto.ClientResponseDto
// @Failure 400,409,500 {object} string
// @Router /api/v1/order [post]
func (h *Handler) newOrder(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	createOrder, err := h.services.Order.CreateOrder(r.Context(), body)
	if err != nil {
		if errors.Is(err, constants.InvalidInputBodyError) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
			return
		}
		if errors.Is(err, constants.AlreadyExistsError) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusConflict, "already created")
			return
		}
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "unknown error")
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, createOrder)
}
