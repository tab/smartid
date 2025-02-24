package identity

import (
	"regexp"

	"github.com/tab/smartid/internal/errors"
)

// https://github.com/SK-EID/smart-id-documentation#2322-etsisemantics-identifier
const (
	// TypeIDC is the identity type for identity card
	TypeIDC = "IDC"
	// TypePAS is the identity type for passport
	TypePAS = "PAS"
	// TypePNO is the identity type for personal number
	TypePNO = "PNO"
)

// Regex is the regular expression for identity numbers
var Regex = regexp.MustCompile(`^(PAS|IDC|PNO)([A-Z]{2})-([A-Za-z0-9-]+)$`)

// Identity represents semantics identifier
type Identity struct {
	Country string
	Type    string
	ID      string
}

// Parse parses the identity number
func Parse(value string) (*Identity, error) {
	if value == "" {
		return nil, errors.ErrEmptyIdentityNumber
	}

	matches := Regex.FindStringSubmatch(value)
	if len(matches) != 4 {
		return nil, errors.ErrInvalidIdentityNumber
	}

	return &Identity{
		Type:    matches[1],
		Country: matches[2],
		ID:      matches[3],
	}, nil
}
