package metric

import (
	"encoding/json"
	"net/http"
)

type JsonHandler struct {
	http.Handler
}

func (h *JsonHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	log := GetLogger()
	m := GetMiner()
	if m == nil {
		log.Error(nil, "failed to get miner")
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		log.Info("NG", "status-code", http.StatusServiceUnavailable)
		return
	}

	// fetch
	dataFetched, err := m.Fetch()
	if err != nil {
		log.Error(err, "failed to fetch data")
		http.Error(w, "failed to fetch data, upstream not available at this time", http.StatusServiceUnavailable)
		log.Info("NG", "status-code", http.StatusServiceUnavailable)
		return
	}

	// parse
	status, err := m.Parse(dataFetched)
	if err != nil {
		log.Error(err, "failed to parse data")
		http.Error(w, "failed to parse data, upstream responded with wrong response", http.StatusServiceUnavailable)
		log.Info("NG", "status-code", http.StatusServiceUnavailable)
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
