package background

import (
	"gogogo/pkg/background/internal"
	"time"
)

func Schedule() func() {
	return schedule
}
func schedule() {
	internal.Logger.Info(time.Now())
}
