// SPDX-License-Identifier: AGPL-3.0-or-later
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
