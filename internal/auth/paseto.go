package auth

import (
	"fmt"
	"github.com/vk-rv/pvx"
	"pasetoservice/internal/models"
	"time"
)

type PasetoAuth struct {
	pasetoKey    *pvx.SymKey
	symmetricKey []byte
}

const keySize = 32

var ErrInvalidSize = fmt.Errorf("bad key size: it must be %d bytes", keySize)

func NewPaseto(key []byte) (*PasetoAuth, error) {

	if len(key) != keySize {
		return nil, ErrInvalidSize
	}

	pasetoKey := pvx.NewSymmetricKey(key, pvx.Version4)

	return &PasetoAuth{
		symmetricKey: key,
		pasetoKey:    pasetoKey,
	}, nil
}

func (pa *PasetoAuth) NewToken(data models.TokenData) (string, error) {

	serviceClaims := &models.ServiceClaims{}

	iss := time.Now()
	exp := iss.Add(data.Duration)

	serviceClaims.IssuedAt = &iss
	serviceClaims.Expiration = &exp
	serviceClaims.Subject = data.Subject

	serviceClaims.AdditionalClaims = data.AdditionalClaims
	serviceClaims.Footer = data.Footer

	pv4 := pvx.NewPV4Local()

	authToken, err := pv4.Encrypt(pa.pasetoKey, serviceClaims,
		pvx.WithFooter(serviceClaims.Footer))
	if err != nil {
		return "", err
	}

	return authToken, nil

}

func (pa *PasetoAuth) VerifyToken(token string) (*models.ServiceClaims, error) {
	pv4 := pvx.NewPV4Local()
	tk := pv4.Decrypt(token, pa.pasetoKey)

	f := models.Footer{}
	sc := models.ServiceClaims{
		Footer: f,
	}

	err := tk.Scan(&sc, &f)
	if err != nil {
		return &sc, err
	}

	return &sc, nil
}
