package protocol

// EnableEncryption is sent by the server to enable websocket encryption. The server sends its public key and
// salt in the command itself, and the client submits its own public key in a response.
//
// This is a client-side command which can be called even on third-party servers that do not implement it.
type EnableEncryption struct {
	// PublicKey is the public key of the client. It is base64 encoded, so it must first be decoded before it
	// can be converted to a key.
	PublicKey string `json:"publicKey"`
}
