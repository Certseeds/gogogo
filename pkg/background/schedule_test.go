// SPDX-License-Identifier: AGPL-3.0-or-later
package background

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScheduleNotNull(t *testing.T) {
	NotNilSchedule := Schedule()
	assert.NotNilf(t, NotNilSchedule, "should not nil")
}
