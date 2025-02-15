// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/config"
)

func TestLenientConfigLoad(t *testing.T) {
	conf := map[string]any{
		"audit": map[string]any{
			"enabled": true,
			"backend": "local",
			"local": map[string]any{
				"storagePath": t.TempDir(),
			},
			"wibble": "wobble",
		},
	}

	require.NoError(t, config.LoadMap(conf))

	c := &audit.Conf{}
	err := config.GetSection(c)

	require.NoError(t, err)
	require.True(t, c.Enabled)
	require.Equal(t, "local", c.Backend)
}
