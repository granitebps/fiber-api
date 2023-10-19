package scheduler

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/granitebps/fiber-api/pkg/core"
)

func SetupScheduler(c *core.Core) {
	lc, _ := time.LoadLocation("Asia/Jakarta")

	s := gocron.NewScheduler(lc)

	_, err := s.Every(1).Minutes().Do(SendHealthCheckSignal, c)
	if err != nil {
		log.Panic(err)
	}

	s.StartAsync()
}
