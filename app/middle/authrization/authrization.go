package authrization

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

type SupabaseJwtPayload struct {
	jwt.StandardClaims
	Email        string               `json:"email,omitempty"`
	Phone        string               `json:"phone,omitempty"`
	AppMetadata  SupabaseAppMetadata  `json:"app_metadata,omitempty"`
	UserMetadata SupabaseUserMetadata `json:"user_metadata,omitempty"`
	Role         string               `json:"role,omitempty"`
	Aal          string               `json:"aal,omitempty"`
	Amr          []SupabaseAmr        `json:"amr,omitempty"`
	SessionID    string               `json:"session_id,omitempty"`
}

type SupabaseAppMetadata struct {
	Provider  string   `json:"provider,omitempty"`
	Providers []string `json:"providers,omitempty"`
}

type SupabaseUserMetadata struct {
	AvatarURL         string `json:"avatar_url,omitempty"`
	Email             string `json:"email,omitempty"`
	EmailVerified     bool   `json:"email_verified,omitempty"`
	FullName          string `json:"full_name,omitempty"`
	Iss               string `json:"iss,omitempty"`
	Name              string `json:"name,omitempty"`
	Picture           string `json:"picture,omitempty"`
	PreferredUsername string `json:"preferred_username,omitempty"`
	ProviderID        string `json:"provider_id,omitempty"`
	Sub               string `json:"sub,omitempty"`
	UserName          string `json:"user_name,omitempty"`
}

func verifyToken(tokenString string) (*CustomClaims, error) {
	SUPABASE_JWT_SECRET := os.Getenv("SUPABASE_JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := []byte(SUPABASE_JWT_SECRET)
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := authorizationHeader[len("Bearer "):]
		claims, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add the claims to the request context for later use in the handler
		ctx := r.Context()
		ctx = context.WithValue(ctx, "claims", claims)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
