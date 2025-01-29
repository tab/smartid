package smartid

import "fmt"

// https://github.com/SK-EID/smart-id-documentation#2322-etsisemantics-identifier
const (
	// TypeIDC is the identity type for identity card
	TypeIDC = "IDC"
	// TypePAS is the identity type for passport
	TypePAS = "PAS"
	// TypePNO is the identity type for personal number
	TypePNO = "PNO"
)

type Identity struct {
	Country string
	Type    string
	ID      string
}

// NewIdentity creates a new identity instance
func NewIdentity(typ, country, id string) string {
	identity := Identity{
		Country: country,
		Type:    typ,
		ID:      id,
	}

	return identity.String()
}

// String returns a string representation
func (i *Identity) String() string {
	return fmt.Sprintf("%v%v-%v", i.Type, i.Country, i.ID)
}
