package log

import "go.uber.org/zap"

var log *zap.SugaredLogger

// Init returns logger file
func Init() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./log/app.log",
	}
	logger, _ := cfg.Build()
	log = logger.Sugar()
}

// Errorf
func Errorf(msg string) {
	log.Errorf(msg)
}
