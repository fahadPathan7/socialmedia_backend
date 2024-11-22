package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fahadPathan7/socialmedia_backend/batman"
	pb "github.com/fahadPathan7/socialmedia_backend/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var instantiated Authenticator
var once sync.Once

// Authenticator provides access to encode and decode JWT
type Authenticator interface {
	EncodeUser(pb.User, int) (string, error)
	DecodeToken(string) (pb.User, error)
	UserFromContext(context.Context) (pb.User, error)
	IsAuthorized(ctx context.Context, ownerID string) (pb.User, bool, error)
	IsAdmin(ctx context.Context) (bool, error)
}

type authenticator struct {
	signingKey []byte
}

type tokenClaims struct {
	jwt.StandardClaims
	UserName    string
	UserEmail string
	Roles     []string
}

// EncodeUser encodes a user
func (a *authenticator) EncodeUser(user pb.User, t int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(t) * time.Hour).Unix(),
			Issuer:    "Social",
			Subject:   user.Id,
		},
		UserName:    user.Username,
		UserEmail: user.Email,
		Roles:     user.Roles,
	})

	ss, err := token.SignedString(a.signingKey)
	if err != nil {
		return "", errors.New("Invalid sign in key")
	}
	return ss, nil
}

// DecodeToken decodes a token
func (a *authenticator) DecodeToken(token string) (pb.User, error) {

	res, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.signingKey, nil
	})
	if err != nil {
		return pb.User{}, err
	}

	if !res.Valid {
		return pb.User{}, errors.New("Invalid token")
	}
	claims := res.Claims.(*tokenClaims)
	return pb.User{
		Username:  claims.UserName,
		Email: 	   claims.UserEmail,
		Roles:     claims.Roles,
	}, nil
}

// UserFromContext gets user from context
func (a *authenticator) UserFromContext(ctx context.Context) (pb.User, error) {
	// Extract headers
	const bearerSchema string = "Bearer"
	var md metadata.MD
	var ok bool
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return pb.User{}, fmt.Errorf("no headers found")
	}

	// Extract authorization header
	authHeader, ok := md["authorization"]
	if !ok {
		return pb.User{}, fmt.Errorf("authorization header not found")
	}
	if len(authHeader) < 1 {
		return pb.User{}, errors.New("authorization header is invalid")
	}

	if !strings.HasPrefix(authHeader[0], bearerSchema) {
		return pb.User{}, errors.New("authorization header is invalid")
	}

	return a.DecodeToken(authHeader[0][len(bearerSchema):])
}

// IsAuthorized extracts auth header, fetches the user, check the logged in user's role and if it's the same user as owner id.
func (a *authenticator) IsAuthorized(ctx context.Context, ownerID string) (pb.User, bool, error) {
	// Is system generated call
	yes := isSystemGeneratedCall(ctx)
	if yes { // System call, no validation needed
		return pb.User{}, true, nil
	}

	// Fetch user from JWT
	user, err := a.UserFromContext(ctx)
	if err != nil {
		errs := batman.Errors{
			Errors: []batman.ErrorDetails{
				{
					Error:   fmt.Sprintf("Token not valid: %v", err),
					Field:   "",
					Message: "Token not valid",
				},
			},
		}
		st := batman.ComposeErrors(codes.Unauthenticated, "Token not valid", errs)
		return pb.User{}, false, st.Err()
	}

	// Check if admin
	if stringInSlice("admin", user.Roles) {
		return user, true, nil
	}

	if user.Id != ownerID {
		st := batman.ComposeMultipleErrorStr(
			codes.PermissionDenied,
			"Permission denied",
			[]string{
				"You do not have the permission to perform this action.",
			},
		)
		return user, false, st.Err()
	}

	return user, true, nil
}

// IsAuthorized extracts auth header, fetches the user, check the logged in user's role and if the user is an admin
func (a *authenticator) IsAdmin(ctx context.Context) (bool, error) {
	// Fetch user from JWT
	user, err := a.UserFromContext(ctx)
	if err != nil {
		errs := batman.Errors{
			Errors: []batman.ErrorDetails{
				{
					Error:   fmt.Sprintf("Token not valid: %v", err),
					Field:   "",
					Message: "Token not valid",
				},
			},
		}
		st := batman.ComposeErrors(codes.Unauthenticated, "Token not valid", errs)
		return false, st.Err()
	}

	// Check if admin
	if stringInSlice("admin", user.Roles) {
		return true, nil
	} else {
		return false, nil
	}
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// New returns a new authenticator
func New() Authenticator {
	key := "5AFFA4C31416C5DFA64065B899A515DE11FBEFB5B607A4894DBF96040EA48725"
	once.Do(func() {
		instantiated = &authenticator{
			signingKey: []byte(key),
		}
	})
	return instantiated
}

func isSystemGeneratedCall(ctx context.Context) bool {
	// Extract headers
	var md metadata.MD
	var ok bool
	md, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	fmt.Println("Extracting system token")
	// Extract system-authorized header
	_, ok = md["system-authorized"]
	if !ok {
		return false
	}

	return true
}
