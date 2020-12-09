package handler

import (
	"net/http"
)

// HealthCheckHandler object for health check handler
type HealthCheckHandler struct {
	HandlerOption
	//http.Handler
}

// NewHealtCheckHandler initialize new health check handler
func NewHealtCheckHandler(option HandlerOption) *HealthCheckHandler {
	return &HealthCheckHandler{
		HandlerOption: option,
	}
}

// Check checks if the service is live or not
func (h HealthCheckHandler) Check(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"health_check": "ok"}`))

	return nil
}

// Readiness checking if all work well
func (h HealthCheckHandler) Readiness(w http.ResponseWriter, r *http.Request) (err error) {
	if h.HandlerOption.ProviderConfig.GetBool("mysql.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckDbMysql()
		if err != nil {
			// TODO: add logger
			return
		}
	}

	if h.HandlerOption.ProviderConfig.GetBool("postgre.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckDbPostgres()
		if err != nil {
			// TODO: add logger
			return
		}
	}

	if h.HandlerOption.ProviderConfig.GetBool("cache.is_enabled") {
		err = h.Services.HealthCheck.HealthCheckDbCache()
		if err != nil {
			// TODO: add logger
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"readiness": "ok"}`))

	return
}
