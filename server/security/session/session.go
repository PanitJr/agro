package session

import (
	"errors"
	"sync"
	"time"

	"agro/components/user"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	signedString = "amoeba-srfc"
)

var (
	storeToken  = make(map[string]*JWTCustomClaims)
	mu          = &sync.Mutex{}
	CurrentUser *user.User
)

func NewJWTCustomClaims(user *user.User) *JWTCustomClaims {
	return &JWTCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
}

type JWTCustomClaims struct {
	*user.User `json:"user"`
	jwt.StandardClaims
}

func (jwtCustom *JWTCustomClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustom)
	t, err := token.SignedString([]byte(signedString))
	if err != nil {
		return "", err
	}
	jwtCustom.New()
	return t, nil
}

func (jwtCustom *JWTCustomClaims) New() {
	mu.Lock()
	storeToken[jwtCustom.StandardClaims.Id] = jwtCustom
	mu.Unlock()
}

func (jwtCustom *JWTCustomClaims) Get() (*JWTCustomClaims, error) {
	mu.Lock()
	_, ok := storeToken[jwtCustom.StandardClaims.Id]
	if ok {
		jwtCustom = storeToken[jwtCustom.StandardClaims.Id]
		mu.Unlock()
		return jwtCustom, nil
	}
	mu.Unlock()
	return jwtCustom, errors.New("token invalid")
}

func (jwtCustom *JWTCustomClaims) Update() {
	mu.Lock()
	storeToken[jwtCustom.StandardClaims.Id] = jwtCustom
	mu.Unlock()

}

func (jwtCustom *JWTCustomClaims) Delete() {
	mu.Lock()
	delete(storeToken, jwtCustom.StandardClaims.Id)
	mu.Unlock()

}

func ParseToJWTCustomClaims(tokenString string) (*JWTCustomClaims, error) {
	var (
		claims *JWTCustomClaims
		ok     bool
	)
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signedString), nil
	})

	claims, ok = token.Claims.(*JWTCustomClaims)
	if ok && token.Valid {
		CurrentUser = claims.User
		return claims, nil
	}

	return nil, err
}
