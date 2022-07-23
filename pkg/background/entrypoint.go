// SPDX-License-Identifier: AGPL-3.0-or-later
package background

import (
	"github.com/robfig/cron/v3"
	"gogogo/pkg/background/internal"
)

func Main() {
	internal.Logger.Infow("INFO start output")
	internal.Logger.Error("ERROR start output")
	_, err := GetPureConfig()
	if err != nil {
		internal.Logger.Fatal("real config fail")
	}
	configRecord := GetConfig()
	work := cron.New(cron.WithSeconds())
	_, err = work.AddFunc(configRecord.Schedule.Cron, Schedule())
	if err != nil {
		internal.Logger.Fatal("start schedule fail")
	}
	go work.Start()
	defer work.Stop()
	select {}
}
