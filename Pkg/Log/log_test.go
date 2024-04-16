package Log

import "testing"

func TestNewrpcLog(t *testing.T) {
	log := NewrpcLog()
	log.Logger.Debug("test")
}
