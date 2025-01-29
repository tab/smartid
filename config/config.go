package config

// Config is a struct holds the client configuration options
type Config struct {
	RelyingPartyName string
	RelyingPartyUUID string
	CertificateLevel string
	HashType         string
	InteractionType  string
	Text             string
	URL              string
	Timeout          int
}

// Option is a type for functional options
type Option func(*Config)

// WithRelayingPartyName is option to set the RelyingPartyName
func WithRelyingPartyName(name string) Option {
	return func(c *Config) {
		c.RelyingPartyName = name
	}
}

// WithRelayingPartyUUID is option to set the RelyingPartyUUID
func WithRelyingPartyUUID(id string) Option {
	return func(c *Config) {
		c.RelyingPartyUUID = id
	}
}

// WithCertificateLevel is option to set the certificate level
func WithCertificateLevel(level string) Option {
	return func(c *Config) {
		c.CertificateLevel = level
	}
}

// WithHashType is option to set the hash type
func WithHashType(hashType string) Option {
	return func(c *Config) {
		c.HashType = hashType
	}
}

// WithInteractionType is option to set the interaction type
func WithInteractionType(interactionType string) Option {
	return func(c *Config) {
		c.InteractionType = interactionType
	}
}

// WithText is option to set the display text
func WithText(text string) Option {
	return func(c *Config) {
		c.Text = text
	}
}

// WithURL is option to set the Smart-ID service URL
func WithURL(url string) Option {
	return func(c *Config) {
		c.URL = url
	}
}

// WithTimeout is option to set the request timeout
func WithTimeout(timeout int) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}
