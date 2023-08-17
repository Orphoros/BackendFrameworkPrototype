package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenDuration is the duration that a JWT is valid for. It is valid for 10 minutes.
const TokenDuration = time.Minute * 10

// JWTClaim is a struct that contains the claims for the JWT.
type JWTClaim struct {
	// UserID is UUID of the user that the JWT is for.
	UserID string
	// StandardClaims are the standard claims for the JWT.
	jwt.StandardClaims
}

// JWT is a struct that contains the JWT key.
type JWT struct {
	jwtKey string
}

// NewJWT creates a new JWT struct.
func NewJWT(key string) *JWT {
	return &JWT{jwtKey: key}
}

// GenerateJWT generates a JWT for a user's UUID. The JWT is valid for the specified duration.
func (store JWT) GenerateJWT(validity time.Duration, userID string) (string, error) {
	expirationTime := time.Now().Add(validity)
	claims := &JWTClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(store.jwtKey))
}

// ValidateToken validates a JWT. Returns an error if the JWT is invalid.
func (store JWT) ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(store.jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return fmt.Errorf("could not parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return fmt.Errorf("token is expired")
	}

	return nil
}

// extractClaims extracts a map of the claims from a JWT. Returns the map of claims and a boolean indicating if the extraction was successful.
func (store JWT) extractClaims(tokenString string) (jwt.MapClaims, bool) {
	hmacSecretString := store.jwtKey
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}

// GetUserIdClaim gets the user's UUID claim from a JWT. Returns the user ID claim pointer (nil if error) and an error if the claim could not be extracted.
func (store JWT) GetUserIdClaim(token string) (*string, error) {
	claims, ok := store.extractClaims(token)

	if !ok {
		return nil, fmt.Errorf("could not parse claims: %v", claims)
	}

	userId, ok := claims["UserID"].(string)

	if !ok {
		return nil, fmt.Errorf("could not parse claims: %v", claims)
	}

	return &userId, nil
}
