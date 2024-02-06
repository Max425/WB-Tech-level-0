package handler

import (
	"encoding/json"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"io"
	"net/http"
)

// @Summary create Delivery
// @Tags delivery
// @Accept  json
// @Produce  json
// @Param input body core.Delivery true "New delivery"
// @Success 200 {object} core.Delivery
// @Failure 400,500 {object} string
// @Router /api/v1/delivery [post]
func (h *Handler) newDelivery(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	var delivery core.Delivery
	err = json.Unmarshal(body, &delivery)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	createDeliveryID, err := h.services.Delivery.Create(r.Context(), &delivery)
	if err != nil {
		dto.SendData(w, "some error", http.StatusInternalServerError)
		return
	}

	response := dto.CreateResponse{ID: createDeliveryID}
	dto.SendData(w, response, http.StatusOK)
}
