package command_line

import (
	"github.com/saltpay/enterprise-temp-converter/specifications"
	"testing"
)

func TestTempConverter(t *testing.T) {
	specifications.ConvertTemperatures(t, Driver{})
}
