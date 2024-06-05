package tokenizer

import "errors"

var (
	ErrParsingToken = errors.New("token has invalid struct or it is expired")
)

type TokenData struct {
	EmployeeId uint
	CompanyId  uint
	IsExpired  bool
}

type Tokenizer interface {
	DecryptToken(token string) (TokenData, error)
}

func New(signKey string) Tokenizer {
	return &tokenizerImpl{
		signKey: []byte(signKey),
	}
}
