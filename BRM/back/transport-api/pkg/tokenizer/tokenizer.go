package tokenizer

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type tokenizerImpl struct {
	signKey []byte
}

func (t *tokenizerImpl) DecryptToken(tokenStr string) (TokenData, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrParsingToken
		}
		return t.signKey, nil
	})
	if err != nil || !token.Valid {
		return TokenData{}, ErrParsingToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return TokenData{}, ErrParsingToken
	}

	employeeIdStr, ok := claims["employee-id"].(string)
	if !ok {
		return TokenData{}, ErrParsingToken
	}

	companyIdStr, ok := claims["company-id"].(string)
	if !ok {
		return TokenData{}, ErrParsingToken
	}

	employeeId, err := strconv.ParseUint(employeeIdStr, 10, 64)
	if err != nil {
		return TokenData{}, ErrParsingToken
	}

	companyId, err := strconv.ParseUint(companyIdStr, 10, 64)
	if err != nil {
		return TokenData{}, ErrParsingToken
	}

	expTimeStr, ok := claims["exp"].(string)
	if !ok {
		return TokenData{}, ErrParsingToken
	}

	expTimeInt, err := strconv.ParseUint(expTimeStr, 10, 64)
	if err != nil {
		return TokenData{}, ErrParsingToken
	}

	return TokenData{
		EmployeeId: uint(employeeId),
		CompanyId:  uint(companyId),
		IsExpired:  time.Now().UTC().After(time.Unix(int64(expTimeInt), 0)),
	}, nil
}
