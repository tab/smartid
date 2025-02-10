package config

import "time"

// Config is a struct holds the client configuration options
type Config struct {
	RelyingPartyName string
	RelyingPartyUUID string
	CertificateLevel string
	HashType         string
	InteractionType  string
	Text             string
	URL              string
	Timeout          time.Duration
}

// WorkerConfig is a struct holds the worker pool configuration options
type WorkerConfig struct {
	Concurrency int
	QueueSize   int
}
