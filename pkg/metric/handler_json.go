package metric

import (
	"encoding/json"
	"net/http"
)

type JsonHandler struct {
	http.Handler
}

func (h *JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := GetLogger().WithName("JsonHandler.ServeHTTP")
	status, err := fetchAndGetStatus()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		log.Error(err, "NG", "msg", "failed to fetch and get status from api endpoint", "status-code", http.StatusServiceUnavailable)
		return
	}

	// marshal
	mOutput, err := json.Marshal(*status)
	if err != nil {
		log.Error(err, "failed to marshal data")
		http.Error(w, "failed to marshal data, unexpected error", http.StatusInternalServerError)
		log.Info("NG", "status-code", http.StatusServiceUnavailable)
		return
	}

	http.Error(w, string(mOutput), http.StatusOK)
	log.Info("OK", "status-code", http.StatusOK, "body-size", len(mOutput))
}
