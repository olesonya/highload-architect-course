package metrics

import (
	"context"
	"errors"
	"net/http"

	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	logger "github.com/sirupsen/logrus"
)

type Config interface {
	GetMetricsBindAddress() string
}

type Server struct {
	httpServer *http.Server
}

func NewServer(config Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: config.GetMetricsBindAddress(),
		},
	}
}

func (s *Server) Start() {
	http.Handle("/metrics", promhttp.Handler())

	logger.Infof("Metrics server is running on %s", s.httpServer.Addr)

	go func() {
		if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("s.httpServer.ListenAndServe(): %v", err)
		}
	}()
}

func (s *Server) Stop() {
	if err := s.httpServer.Shutdown(context.TODO()); err != nil {
		logger.Fatalf("s.httpServer.Shutdown(...): %v", err)
	}

	logger.Infoln("Metrics server graceful stop")
}
