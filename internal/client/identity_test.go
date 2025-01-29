package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Identity_String(t *testing.T) {
	tests := []struct {
		name     string
		kind     string
		country  string
		id       string
		expected string
	}{
		{
			name:     "IDC",
			kind:     TypeIDC,
			country:  "EE",
			id:       "30303039914",
			expected: "IDCEE-30303039914",
		},
		{
			name:     "PAS",
			kind:     TypePAS,
			country:  "EE",
			id:       "30303039914",
			expected: "PASEE-30303039914",
		},
		{
			name:     "PNO",
			kind:     TypePNO,
			country:  "EE",
			id:       "30303039914",
			expected: "PNOEE-30303039914",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := NewIdentity(tt.kind, tt.country, tt.id)
			assert.Equal(t, tt.expected, identity)
		})
	}
}
