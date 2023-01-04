package main

import (
	"time"

	"github.com/kgrigorev/hellodeb/internal/blop"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("starting blop...",
		zap.Duration("tick", 3*time.Second),
	)
	blop.Run()
}
