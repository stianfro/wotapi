package v1

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// HealthZ godoc
// @Summary		  Health Check
// @Description simple health check handler that returns 200 OK and is used for Kubernetes liveness probes
// @Tags        health
// @Produce     json
// @Success     200 {string} HealthZ
// @Router      /healthz [get]
func HealthZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to write response")
	}
}
