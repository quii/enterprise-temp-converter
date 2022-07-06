package command_line

import (
	"github.com/saltpay/enterprise-temp-converter/specifications"
	"testing"
)

func TestTempConverter(t *testing.T) {
	t.Run("it passes the temp converter spec", func(t *testing.T) {
		specifications.ItConvertsTemperatures(t, Driver{})
	})
}
