package handler

import (
	"encoding/json"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"io"
	"net/http"
)

// @Summary create Customer
// @Tags customer
// @Accept  json
// @Produce  json
// @Param input body core.Customer true "New customer"
// @Success 200 {object} core.Customer
// @Failure 400,500 {object} string
// @Router /api/v1/customer [post]
func (h *Handler) newCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	var customer core.Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	err = h.services.Customer.Create(r.Context(), &customer)
	if err != nil {
		dto.SendData(w, "some error", http.StatusInternalServerError)
		return
	}

	response := dto.CreateResponse{ID: customer.CustomerUid}
	dto.SendData(w, response, http.StatusOK)
}
