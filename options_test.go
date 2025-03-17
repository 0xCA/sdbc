package sdbc

import (
	"log/slog"
	"testing"

	"gotest.tools/v3/assert"
)

func TestEmptyLogHandler(t *testing.T) {
	t.Parallel()

	handler := emptyLogHandler{}.
		WithAttrs(nil).
		WithGroup("group")

	assert.Check(t, !handler.Enabled(nil, 0))              //lint:ignore SA1012 intentional nil
	assert.NilError(t, handler.Handle(nil, slog.Record{})) //lint:ignore SA1012 intentional nil
}
