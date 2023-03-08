package auth

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/borud/tt/pkg/model"
	"github.com/borud/tt/pkg/util"
	"github.com/golang-jwt/jwt/v4"
)

// JWT takes care of generating and parsing JWT tokens.
type JWT struct {
	config JWTConfig
}

// JWTConfig for token manager
type JWTConfig struct {
	Issuer          string
	Secret          []byte
	DefaultValidity time.Duration
	Audience        []string
}

// Errors for auth
var (
	ErrMustHaveDefaultValidity = errors.New("must have default validity period")
	ErrInvalidSignature        = errors.New("invalid signature")
	ErrMustHaveSecret          = errors.New("secret must be specified")
	ErrMustHaveDefaultAudience = errors.New("default audience must be specified")
	ErrClaimsWrongType         = errors.New("claims is wrong type")
	ErrNoClaimsFound           = errors.New("no claims found in context")
	ErrWrongTypeForClaim       = errors.New("wrong type for claim")
	ErrInvalidToken            = errors.New("invalid token")
	ErrTokenExpired            = errors.New("token has expired")
	ErrTokenNotYetValid        = errors.New("token is not yet valid")
	ErrWrongIssuer             = errors.New("wrong issuer")
	ErrNoSubject               = errors.New("subject is missing")
)

// NewJWT creates new token manager
func NewJWT(c JWTConfig) (*JWT, error) {
	if c.DefaultValidity == 0 {
		return nil, ErrMustHaveDefaultValidity
	}

	if len(c.Secret) == 0 {
		return nil, ErrMustHaveSecret
	}

	if len(c.Audience) == 0 {
		return nil, ErrMustHaveDefaultAudience
	}

	return &JWT{config: c}, nil
}

// Create creates a JWT token for user.  Optionally a validFor time.Duration
// argument can be specified to override the default validity period.
func (a *JWT) Create(user model.User, validFor ...time.Duration) (string, error) {
	// Override default validity period if a validFor argument is supplied.
	if len(validFor) == 0 {
		validFor = []time.Duration{a.config.DefaultValidity}
	}

	// Create random unique identifier for token.
	id, err := util.RandomString(128)
	if err != nil {
		return "", err
	}

	claims := jwt.RegisteredClaims{
		ID:        id,
		Issuer:    a.config.Issuer,
		Subject:   user.Username,
		Audience:  a.config.Audience,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(validFor[0])),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(a.config.Secret)
}

// Audience returns the audience for this JWT instance.
func (a *JWT) Audience() []string {
	return a.config.Audience
}

// ParseAndVerify integrity of JWT token.  Returns claims on success and a non-nil
// error on failure.
func (a *JWT) ParseAndVerify(tokenString string) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.config.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	jwtClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrClaimsWrongType, reflect.TypeOf(token.Claims))
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, err := getClaims(jwtClaims)
	if err != nil {
		return nil, err
	}

	if claims.EXP.Before(time.Now()) {
		return nil, ErrTokenExpired
	}

	if time.Now().Before(claims.NBF) {
		return nil, ErrTokenNotYetValid
	}

	if claims.ISS != a.config.Issuer {
		return nil, ErrWrongIssuer
	}

	if claims.SUB == "" {
		return nil, ErrNoSubject
	}

	return claims, nil
}

// getClaims unpacks the claims.
func getClaims(claims jwt.MapClaims) (*Claims, error) {
	jti, ok := claims["jti"].(string)
	if !ok {
		return nil, fmt.Errorf("%w: jti should be string", ErrWrongTypeForClaim)
	}

	iss, ok := claims["iss"].(string)
	if !ok {
		return nil, fmt.Errorf("%w: iss should be string", ErrWrongTypeForClaim)
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("%w: sub should be string", ErrWrongTypeForClaim)
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("%w: exp should be float", ErrWrongTypeForClaim)
	}

	nbf, ok := claims["nbf"].(float64)
	if !ok {
		return nil, fmt.Errorf("%w: nbf should be float", ErrWrongTypeForClaim)
	}

	iat, ok := claims["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("%w: iat should be float", ErrWrongTypeForClaim)
	}

	return &Claims{
		JTI: jti,
		ISS: iss,
		SUB: sub,
		EXP: time.Unix(int64(exp), 0),
		NBF: time.Unix(int64(nbf), 0),
		IAT: time.Unix(int64(iat), 0),
	}, nil
}
