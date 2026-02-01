package logging

import "testing"

func TestLogger(t *testing.T) {
	logger := Logger()

	t.Run("Logger should not be nil", func(t *testing.T) {
		if logger.GetSink() == nil {
			t.Error("Expected logger to be initialized, got nil")
		}
	})

	t.Run("Logger should be singleton", func(t *testing.T) {
		logger2 := Logger()
		if logger.GetSink() != logger2.GetSink() {
			t.Error("Expected same logger instance, got different instances")
		}
	})
}
