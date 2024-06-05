package tokenizer

import "time"

type Tokenizer interface {
	CreateToken(employeeId uint64, companyId uint64) (string, error)
	CheckExpiration(token string) (bool, error)
	DecryptToken(token string) (uint64, uint64, error)
}

func New(
	expiration time.Duration,
	signKey []byte,
) Tokenizer {
	return &tokenizerImpl{
		accessExpiration: expiration,
		signKey:          signKey,
	}
}
