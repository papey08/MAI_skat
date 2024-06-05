package tokenizer

import (
	"auth/internal/model"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type tokenizerImpl struct {
	accessExpiration time.Duration
	signKey          []byte
}

func (t *tokenizerImpl) CreateToken(employeeId uint64, companyId uint64) (string, error) {
	claims := jwt.MapClaims{
		"employee-id": strconv.FormatUint(employeeId, 10),
		"company-id":  strconv.FormatUint(companyId, 10),
		"exp":         strconv.FormatInt(time.Now().UTC().Add(t.accessExpiration).Unix(), 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(t.signKey)
	if err != nil {
		return "", model.ErrCreateAccessToken
	}
	return signedToken, nil
}

// CheckExpiration returns true if token still could be used and false if it has expired
func (t *tokenizerImpl) CheckExpiration(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrParsingAccessToken
		}
		return t.signKey, nil
	})
	if err != nil || !token.Valid {
		return false, model.ErrParsingAccessToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, model.ErrParsingAccessToken
	}

	expTime := time.Unix(int64(claims["exp"].(float64)), 0)
	return time.Now().UTC().Before(expTime), nil
}

func (t *tokenizerImpl) DecryptToken(tokenStr string) (uint64, uint64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrParsingAccessToken
		}
		return t.signKey, nil
	})
	if err != nil || !token.Valid {
		return 0, 0, model.ErrParsingAccessToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, 0, model.ErrParsingAccessToken
	}

	employeeIdStr, ok := claims["employee-id"].(string)
	if !ok {
		return 0, 0, model.ErrParsingAccessToken
	}

	companyIdStr, ok := claims["company-id"].(string)
	if !ok {
		return 0, 0, model.ErrParsingAccessToken
	}

	employeeId, err := strconv.ParseUint(employeeIdStr, 10, 64)
	if err != nil {
		return 0, 0, model.ErrParsingAccessToken
	}

	companyId, err := strconv.ParseUint(companyIdStr, 10, 64)
	if err != nil {
		return 0, 0, model.ErrParsingAccessToken
	}

	return employeeId, companyId, nil
}
