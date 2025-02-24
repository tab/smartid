package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tab/smartid/internal/errors"
)

func Test_Identity_Parse(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		valid  bool
		expect *Identity
		err    error
	}{
		{
			name:  "Success: PNOEE-30303039914",
			input: "PNOEE-30303039914",
			valid: true,
			expect: &Identity{
				Type:    TypePNO,
				Country: "EE",
				ID:      "30303039914",
			},
		},
		{
			name:  "Success: PASEE-30303039914",
			input: "PASEE-30303039914",
			valid: true,
			expect: &Identity{
				Type:    TypePAS,
				Country: "EE",
				ID:      "30303039914",
			},
		},
		{
			name:  "Success: IDCEE-30303039914",
			input: "IDCEE-30303039914",
			valid: true,
			expect: &Identity{
				Type:    TypeIDC,
				Country: "EE",
				ID:      "30303039914",
			},
		},
		{
			name:  "Success: PNOLT-40504040001",
			input: "PNOLT-40504040001",
			valid: true,
			expect: &Identity{
				Type:    TypePNO,
				Country: "LT",
				ID:      "40504040001",
			},
		},
		{
			name:  "Success: PNOLV-050405-10009",
			input: "PNOLV-050405-10009",
			valid: true,
			expect: &Identity{
				Type:    TypePNO,
				Country: "LV",
				ID:      "050405-10009",
			},
		},
		{
			name:  "Success: PNOBE-05040400032",
			input: "PNOBE-05040400032",
			valid: true,
			expect: &Identity{
				Type:    TypePNO,
				Country: "BE",
				ID:      "05040400032",
			},
		},
		{
			name:   "Error: invalid identity",
			input:  "PNOEE-30303039914#",
			valid:  false,
			expect: nil,
			err:    errors.ErrInvalidIdentityNumber,
		},
		{
			name:   "Error: missing identifier",
			input:  "PNOEE-",
			valid:  false,
			expect: nil,
			err:    errors.ErrInvalidIdentityNumber,
		},
		{
			name:   "Error: missing hyphen",
			input:  "PNOEE30303039914",
			valid:  false,
			expect: nil,
			err:    errors.ErrInvalidIdentityNumber,
		},
		{
			name:   "Error: missing type",
			input:  "EE-30303039914",
			valid:  false,
			expect: nil,
			err:    errors.ErrInvalidIdentityNumber,
		},
		{
			name:   "Error: missing country",
			input:  "PNO-30303039914",
			valid:  false,
			expect: nil,
			err:    errors.ErrInvalidIdentityNumber,
		},
		{
			name:   "Error: missing country and type",
			input:  "30303039914",
			valid:  false,
			expect: nil,
			err:    errors.ErrInvalidIdentityNumber,
		},
		{
			name:   "Error: empty",
			input:  "",
			valid:  false,
			expect: nil,
			err:    errors.ErrEmptyIdentityNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity, err := Parse(tt.input)

			if tt.valid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, identity)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
			}
		})
	}
}
