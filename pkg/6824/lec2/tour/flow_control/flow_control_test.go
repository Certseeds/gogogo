package flow_control

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestForCycle(t *testing.T) {
	forCycle()
}

func TestForCycleWithStateEmpty(t *testing.T) {
	forCycleWithStateEmpty()
}

func TestForCycleWithoutThreePartOrCalledWhile(t *testing.T) {
	forCycleWithoutThreePartOrCalledWhile()
}

func TestForEver(t *testing.T) {
	assert.Equal(t, forEver(), 30)
}

func TestSqrt(t *testing.T) {
	sqrt(-1)
	sqrt(0)
	sqrt(1)
}

func TestPowWithLimit(t *testing.T) {
	assert.Equal(t, powWithLimit(2, 10, 512), float64(512))
	assert.Equal(t, powWithLimit(2, 10, 2048), float64(1024))
}

func TestPowWithLimittButPrint(t *testing.T) {
	assert.Equal(t, powWithLimitButPrint(2, 10, 512), float64(512))
	assert.Equal(t, powWithLimitButPrint(2, 10, 2048), float64(1024))
}

func TestSqrtReal(t *testing.T) {
	assert.InDelta(t, math.Sqrt(2), SqrtReal(2), 0.0000001)
}

func TestSwitchRunningOs(t *testing.T) {
	switchRunningOS()
}

func TestSwitchWithEvaluationCase(t *testing.T) {
	switchWithEvaluationCase()
}

func TestSwitchWithoutInput(t *testing.T) {
	switchWithoutInput()
}
func TestDeferPrint(t *testing.T) {
	deferPrint()
}
func TestDeferPrintStack(t *testing.T) {
	deferPrintStack()
}
