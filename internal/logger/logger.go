package logger

import "go.uber.org/zap"

// New returns a configured zap logger (production config by default)
func New(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	switch level {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	return cfg.Build()
}
