package metric

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sokdak/miner-exporter/pkg/dto"
)

const (
	namespace         = "mining"
	subsystemForGpu   = "gpu"
	subsystemForMiner = "miner"
)

type GpuMetric struct {
	Type   prometheus.ValueType
	Desc   *prometheus.Desc
	Value  func(*dto.Status, int) float64
	Labels func(*dto.Status, int) []string
}

type MinerMetric struct {
	Type   prometheus.ValueType
	Desc   *prometheus.Desc
	Value  func(*dto.Status) float64
	Labels func(*dto.Status) []string
}

var (
	defaultMinerMetricLabels    = []string{"worker", "address", "algorithm", "pool", "name", "version"}
	defaultMinerGpuMetricLabels = append(defaultMinerMetricLabels, "gpu_id", "model")
)

type MinerCollector struct {
	logger logr.Logger

	up                                prometheus.Gauge
	totalScrapes, fetchOrParseFailure prometheus.Counter

	metricsGpu   []*GpuMetric
	metricsMiner []*MinerMetric
}

func NewMinerMetric(log logr.Logger) *MinerCollector {
	return &MinerCollector{
		logger: log.WithName("miner-metric"),
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: prometheus.BuildFQName(namespace, subsystemForGpu, "up"),
			Help: "Was the last scrape of miner api endpoint successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystemForGpu, "total_scrapes"),
			Help: "Current total miner stat scrapes.",
		}),
		fetchOrParseFailure: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(namespace, subsystemForGpu, "fetch_or_parse_failure"),
			Help: "Number of errors while fetching or parsing data.",
		}),
		metricsGpu: []*GpuMetric{
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "hashrate"),
					"hashrate",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].Hashrate)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "coretemp"),
					"core temp",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].CoreTemp)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "memtemp"),
					"memory temp",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].MemoryTemp)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "power"),
					"power consumption",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].PowerConsumption)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "fan"),
					"fan speed",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].FanSpeed)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "share_accepted"),
					"share accepted",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].ShareAccepted)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "share_rejected"),
					"share rejected",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].ShareRejected)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForGpu, "share_stale"),
					"share stale",
					defaultMinerGpuMetricLabels, nil),
				Value: func(m *dto.Status, id int) float64 {
					return float64(m.Devices[id].ShareStale)
				},
				Labels: defaultMinerGpuMetricLabelValueGenerator,
			},
		},
		metricsMiner: []*MinerMetric{
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForMiner, "gpu_count"),
					"gpu count",
					defaultMinerMetricLabels, nil),
				Value: func(m *dto.Status) float64 {
					return float64(len(m.Devices))
				},
				Labels: defaultMinerMetricLabelValueGenerator,
			},
			{
				Type: prometheus.GaugeValue,
				Desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystemForMiner, "uptime"),
					"uptime",
					defaultMinerMetricLabels, nil),
				Value: func(m *dto.Status) float64 {
					return float64(m.Miner.Uptime)
				},
				Labels: defaultMinerMetricLabelValueGenerator,
			},
		},
	}
}

func defaultMinerMetricLabelValueGenerator(m *dto.Status) []string {
	return []string{m.Miner.Worker, m.Miner.Address, m.Miner.Algorithm, m.Miner.Pool, m.Miner.Name, m.Miner.Version}
}

func defaultMinerGpuMetricLabelValueGenerator(m *dto.Status, id int) []string {
	return append(defaultMinerMetricLabelValueGenerator(m), fmt.Sprintf("%d", m.Devices[id].GpuId), m.Devices[id].Name)
}

func (c *MinerCollector) Collect(ch chan<- prometheus.Metric) {
	c.totalScrapes.Inc()
	defer func() {
		ch <- c.up
		ch <- c.totalScrapes
		ch <- c.fetchOrParseFailure
	}()

	status, err := fetchAndGetStatus()
	if err != nil {
		c.up.Set(0)
		c.logger.Error(err, "failed to fetch and get status")
		return
	}
	c.up.Set(1)

	// set miner metrics
	for _, metric := range c.metricsMiner {
		ch <- prometheus.MustNewConstMetric(
			metric.Desc,
			metric.Type,
			metric.Value(status),
			metric.Labels(status)...)
	}

	// set gpu metrics
	for _, metric := range c.metricsGpu {
		for i := range status.Devices {
			ch <- prometheus.MustNewConstMetric(
				metric.Desc,
				metric.Type,
				metric.Value(status, i),
				metric.Labels(status, i)...)
		}
	}

	c.logger.Info("collect completed", "gpu-metric-count", len(c.metricsGpu), "miner-metric-count", len(c.metricsMiner))
}

func (c *MinerCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metricsGpu {
		ch <- metric.Desc
	}

	for _, metric := range c.metricsMiner {
		ch <- metric.Desc
	}

	ch <- c.up.Desc()
	ch <- c.totalScrapes.Desc()
	ch <- c.fetchOrParseFailure.Desc()
}
