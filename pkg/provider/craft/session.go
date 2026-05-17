package craft

type SessionState uint8

const (
	SessionStateHandshake SessionState = iota
	SessionStateStatus
	SessionStateLogin
	SessionStateConfiguration
	SessionStatePlay
)
