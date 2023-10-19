package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/granitebps/fiber-api/pkg/core"
)

func SetupScheduler(core *core.Core) {
	lc, _ := time.LoadLocation("Asia/Jakarta")

	s := gocron.NewScheduler(lc)

	s.Every(1).Minutes().Do(SendHealthCheckSignal, core)

	s.StartAsync()
}
