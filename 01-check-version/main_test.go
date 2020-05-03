package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckVersion(t *testing.T) {
	version, err := GetAPIVersion()
	require.NoError(t, err)
	require.Contains(t, version, "2.0.")
}
