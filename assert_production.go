//go:build !enable_assertions

package must

import (
	"go.uber.org/zap"
)

func Assert(condition bool, msg string, logger *zap.Logger) {
	if !condition {
		if logger != nil {
			logger.Debug("[Assertion]", zap.String("msg", msg), zap.Bool("condition", condition))
		}
	}
}
