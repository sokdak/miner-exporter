package metric

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleExportMetric(c *gin.Context) {
	m := GetMiner(c)
	log := GetLogger(c)

	// fetch
	dataFetched, err := m.Fetch()
	if err != nil {
		log.Error(err, "failed to fetch data")
		c.String(http.StatusServiceUnavailable, "failed to fetch data, upstream not available at this time")
		return
	}

	// parse
	status, err := m.Parse(dataFetched)
	if err != nil {
		log.Error(err, "failed to parse data")
		c.String(http.StatusServiceUnavailable, "failed to parse data, upstream responded with wrong response")
		return
	}

	c.JSON(200, status)
}
