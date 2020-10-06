// +build azure

package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiagnosticsSettingsResourceExists(t *testing.T) {
	t.Parallel()

	diagnosticsSettingResourceName := "fakename"
	resGroupName := "fakeresgroup"
	subscriptionID := "fakesubid"

	_, err := DiagnosticSettingsResourceExistsE(t, diagnosticsSettingResourceName, resGroupName, subscriptionID)
	require.Error(t, err)
}
