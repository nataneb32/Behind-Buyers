package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAnalytics struct {
	t *testing.T
}

func (m *MockAnalytics) Notify(field string, value interface{}) {
	if field == "test" {
		for a, b := range value.(map[string]interface{}) {
			assert.Equal(m.t, "1", a)
			assert.Equal(m.t, 123, b)
		}
	}
}

func TestHandlerCreate(t *testing.T) {
	assert := assert.New(t)
	mockAnalytics := &MockAnalytics{t: t}
	rh := CreateReportObserver()
	rh.Subscribe(mockAnalytics)
	assert.NotNil(rh)

	report := map[string]interface{}{
		"test": map[string]interface{}{
			"1": 123,
		},
	}
	rh.Report(report)
}
