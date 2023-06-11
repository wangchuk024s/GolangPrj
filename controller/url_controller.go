package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your-username/project-name/model"
)

var urlMappings = make(map[string]model.URLMapping)

// AddURLMapping adds a new URL mapping
func AddURLMapping(w http.ResponseWriter, r *http.Request) {
	var mapping model.URLMapping
	err := json.NewDecoder(r.Body).Decode(&mapping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a random ID for the mapping
	mapping.ID = model.GenerateID()

	// Save the URL mapping
	urlMappings[mapping.ID] = mapping

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapping)
}

// DeleteURLMapping deletes a URL mapping by ID
func DeleteURLMapping(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Check if the mapping exists
	_, exists := urlMappings[id]
	if !exists {
		http.Error(w, "URL mapping not found", http.StatusNotFound)
		return
	}

	// Delete the mapping
	delete(urlMappings, id)

	w.WriteHeader(http.StatusOK)
}

// UpdateURLMapping updates a URL mapping by ID
func UpdateURLMapping(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Check if the mapping exists
	_, exists := urlMappings[id]
	if !exists {
		http.Error(w, "URL mapping not found", http.StatusNotFound)
		return
	}

	var mapping model.URLMapping
	err := json.NewDecoder(r.Body).Decode(&mapping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the mapping
	urlMappings[id] = mapping

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapping)
}

// GetURLMappings retrieves all URL mappings
func GetURLMappings(w http.ResponseWriter, r *http.Request) {
	mappings := make([]model.URLMapping, 0, len(urlMappings))
	for _, mapping := range urlMappings {
		mappings = append(mappings, mapping)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mappings)
}
