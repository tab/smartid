package smartid

import "errors"

var (
	ErrAuthenticationIsRunning  = errors.New("authentication is still running")
	ErrSmartIdNoSuitableAccount = errors.New("no suitable account of requested type found")
	ErrSmartIdMaintenance       = errors.New("system is under maintenance, retry again later")
	ErrUserRefused              = errors.New("user refused")
	ErrTimeout                  = errors.New("user didn't respond in time")
)
