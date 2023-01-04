package main

import (
	"time"

	"github.com/kgrigorev/hellodeb/internal/blap"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("starting blap ...",
		zap.Duration("tick", 3*time.Second),
	)
	blap.Run()
}
