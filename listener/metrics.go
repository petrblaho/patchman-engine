package listener

import (
	"app/base/utils"
	"app/manager/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	EventUpload             = "upload"
	EventDelete             = "delete"
	ReceivedSuccess         = "success"
	ReceivedErrorIdentity   = "error-identity"
	ReceivedErrorParsing    = "error-parsing"
	ReceivedErrorProcessing = "error-processing"
	ReceivedErrorOtherType  = "error-other-type"
)

var (
	messagesReceivedCnt = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "patchman_engine",
		Subsystem: "listener",
		Name:      "kafka_message_received",
	}, []string{"event", "type"})

	messageHandlingDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "patchman_engine",
		Subsystem: "listener",
		Name:      "kafka_message_handling_duration_seconds",
	}, []string{"event"})
)

func init() {
	prometheus.MustRegister(messagesReceivedCnt, messageHandlingDuration)
}

func RunMetrics() {
	// create web app
	app := gin.New()
	middlewares.Prometheus().Use(app)

	err := app.Run(":8081")
	if err != nil {
		utils.Log("err", err.Error()).Error()
		panic(err)
	}
}
