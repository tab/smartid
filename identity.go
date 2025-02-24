package smartid

import (
	"fmt"

	"github.com/tab/smartid/internal/identity"
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

// Identity represents semantics identifier
type Identity struct {
	Country string
	Type    string
	ID      string
}

// NewIdentity creates a new identity instance
func NewIdentity(typ, country, id string) string {
	i := Identity{
		Country: country,
		Type:    typ,
		ID:      id,
	}

	return i.String()
}

// String returns a string representation
func (i *Identity) String() string {
	return fmt.Sprintf("%v%v-%v", i.Type, i.Country, i.ID)
}

// Parse parses the identity number
func Parse(value string) (*Identity, error) {
	result, err := identity.Parse(value)
	if err != nil {
		return nil, err
	}

	return &Identity{
		Country: result.Country,
		Type:    result.Type,
		ID:      result.ID,
	}, nil
}
