package baseservice

// JWTData holds the decoded information of the JWT.
type JWTData struct {
	ID string
}

// JWTGeneratorSvc jwt generator service.
type JWTGeneratorSvc interface {
	// NewJWT should return the jwt generated from the JWTData.
	NewJWT(JWTData) (string, error)
}
