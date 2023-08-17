package security

import "testing"

func TestGenerateJWT(t *testing.T) {
	t.Run("should generate a valid JWT", func(t *testing.T) {
		jwt := NewJWT("test")
		token, err := jwt.GenerateJWT(TokenDuration, "test")
		if err != nil {
			t.Errorf("could not generate JWT: %v", err)
		}
		if token == "" {
			t.Errorf("token is empty")
		}
	})

	t.Run("should validate a valid token", func(t *testing.T) {
		jwt := NewJWT("test")
		token, err := jwt.GenerateJWT(TokenDuration, "test")
		if err != nil {
			t.Errorf("could not generate JWT: %v", err)
		}
		err = jwt.ValidateToken(token)
		if err != nil {
			t.Errorf("could not validate token: %v", err)
		}
	})

	t.Run("should not validate an invalid token", func(t *testing.T) {
		jwt := NewJWT("test")
		token, err := jwt.GenerateJWT(TokenDuration, "test")
		if err != nil {
			t.Errorf("could not generate JWT: %v", err)
		}
		err = jwt.ValidateToken(token + "a")
		if err == nil {
			t.Errorf("token is valid")
		}
	})

	t.Run("should extract user id from a valid token", func(t *testing.T) {
		jwt := NewJWT("test")
		token, err := jwt.GenerateJWT(TokenDuration, "test")
		if err != nil {
			t.Errorf("could not generate JWT: %v", err)
		}

		userId, err := jwt.GetUserIdClaim(token)

		if err != nil {
			t.Errorf("could not extract user id from token: %v", err)
		}

		if *userId != "test" {
			t.Errorf("user id is not correct")
		}
	})
}
