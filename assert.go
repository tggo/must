//go:build enable_assertions

package must

import (
	"go.uber.org/zap"
)

func Assert(condition bool, msg string, logger *zap.Logger) {
	if !condition {
		logger.Fatal("[Assertion]", zap.String("msg", msg), zap.Bool("condition", condition))
	}
}
