package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserID    int32     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil // Adjust as needed for your use case
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.ExpiredAt), nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.IssuedAt), nil
}

func (p *Payload) GetIssuer() (string, error) {
	return "", nil // Adjust as needed for your use case
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil // Adjust as needed for your use case
}

func (p *Payload) GetSubject() (string, error) {
	return "", nil // Adjust as needed for your use case
}
func NewPayload(userID int32, duration time.Duration) *Payload {
	return &Payload{
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}
