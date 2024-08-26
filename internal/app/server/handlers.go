package server

import (
	"encoding/json"
	"net/http"

	"github.com/tonitaga/url-shortener/internal/model/alias"
	"github.com/tonitaga/url-shortener/internal/model/redirect"
)

func (s *Server) redirectCreationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "frontend/405/index.html")
		return
	}

	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected application/json", http.StatusBadRequest)
		return
	}

	var body JsonBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.logger.Infof(
		"Processing request with data: alias '%s', originalUrl '%s'",
		body.Alias,
		body.OriginalUrl,
	)

	// Lookup for originalUrl in storage
	model, err := s.storage.Redirect().FindByTarget(body.OriginalUrl)
	if err == nil {
		body.Alias = model.Alias
		if err := json.NewEncoder(w).Encode(body); err != nil {
			s.logger.WithError(err).Error("Failed to encode body")
			http.Error(w, "Failed to encode body", http.StatusInternalServerError)
			return
		}
		s.logger.Infof(
			"Requested url '%s' already exists in storage, returned alias '%s' to user",
			model.Target,
			model.Alias,
		)
		return
	}

	if len(body.Alias) == 0 {
		// If user doesn't set alias, generate random one
		body.Alias = alias.New(6)
	}

	if err := s.storage.Redirect().Create(&redirect.RedirectModel{
		Alias:  body.Alias,
		Target: body.OriginalUrl,
	}); err != nil {
		s.logger.WithError(err).Error("Failed to create redirect")
		http.Error(w, "Failed to create redirect", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		s.logger.WithError(err).Error("Failed to encode response")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	s.logger.Infof("Created alias '%s' for target '%s'", body.Alias, body.OriginalUrl)
}

func (s *Server) redirectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	alias := r.URL.String()
	alias = alias[1:]

	// Trying to find alias in storage
	model, err := s.storage.Redirect().FindByAlias(alias)
	if err != nil {
		s.logger.WithError(err).Info("Someone tried redirection with unknonw alias")
		http.ServeFile(w, r, "frontend/404/index.html")
		return
	}

	http.Redirect(w, r, model.Target, http.StatusFound)
	s.logger.Infof("Redirection to '%s' finished successfully", model.Target)
}

func (s *Server) showMainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "frontend/index.html")
}
