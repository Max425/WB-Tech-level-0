package handler

import (
	"encoding/json"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"io"
	"net/http"
)

// @Summary create Payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Param input body core.Payment true "New payment"
// @Success 200 {object} core.Payment
// @Failure 400,500 {object} string
// @Router /api/v1/payment [post]
func (h *Handler) newPayment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	var payment core.Payment
	err = json.Unmarshal(body, &payment)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	createPaymentID, err := h.services.Payment.Create(r.Context(), &payment)
	if err != nil {
		dto.SendData(w, "some error", http.StatusInternalServerError)
		return
	}

	response := dto.CreateResponse{ID: createPaymentID}
	dto.SendData(w, response, http.StatusOK)
}
