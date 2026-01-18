package protocol

import "time"

// represents a registered device in the network
type Device struct {
	ID        string
	Name      string
	PublicKey string
	Endpoints []string // IP:port combos
	VirtualIP string   // assigned mesh IP (10.100.0.x)
	LastSeen  time.Time
	Online    bool
}

// RegisterRequest is sent when a device joins network
type RegisterRequest struct {
	Name      string
	PublicKey string
	Endpoints []string
}

// RegisterResponse is returned after successful registration
type RegisterResponse struct {
	DeviceID  string
	VirtualIP string
	Peers     []Device // current peers in the network
}

// HeartbeatRequest is sent periodically to update presence
type HeartbeatRequest struct {
	DeviceID  string
	Endpoints []string // may change due to NAT
}
