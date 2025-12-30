// internal/wrap/handler.go
package wrap

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

// POST /api/wraps
func (h *Handler) CreateWrap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	wrap, err := h.service.MakeWrap(req.Name)
	if err != nil {
		log.Printf("Error creating wrap: %v", err)
		http.Error(w, "Failed to create wrap: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wrap)
}

// GET /api/wraps/all
func (h *Handler) GetAllWraps(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	wraps, err := h.service.GetAllWraps()
	if err != nil {
		log.Printf("Error getting all wraps: %v", err)
		http.Error(w, "Failed to retrieve wraps", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wraps)
}

// GET /api/wraps/{uuid}
func (h *Handler) GetWrap(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract UUID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/get/wrap/")
	if path == "" {
		http.Error(w, "UUID is required", http.StatusBadRequest)
		return
	}

	wrap, err := h.service.GetWrap(path)
	if err != nil {
		http.Error(w, "Failed to get wrap", http.StatusInternalServerError)
		return
	}

	if wrap == nil {
		http.Error(w, "Wrap not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wrap)
}
