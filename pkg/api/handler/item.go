package handler

import (
	"encoding/json"
	"github.com/Max425/WB-Tech-level-0/pkg/model/core"
	"github.com/Max425/WB-Tech-level-0/pkg/model/dto"
	"io"
	"net/http"
)

// @Summary create Item
// @Tags item
// @Accept  json
// @Produce  json
// @Param input body core.Item true "New item"
// @Success 200 {object} core.Item
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/item [post]
func (h *Handler) newItem(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	var item core.Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		dto.SendData(w, "invalid input body", http.StatusBadRequest)
		return
	}
	createItemID, err := h.services.Item.CreateItem(r.Context(), &item)
	if err != nil {
		dto.SendData(w, "some error", http.StatusInternalServerError)
		return
	}

	response := dto.CreateResponse{ID: createItemID}
	dto.SendData(w, response, http.StatusOK)
}
