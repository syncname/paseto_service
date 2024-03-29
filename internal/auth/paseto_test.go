package auth

import (
	"github.com/stretchr/testify/require"
	"pasetoservice/internal/models"
	"testing"
	"time"
)

func TestServiceGenerateNewKey(t *testing.T) {

	key := []byte("000f3e5799296cc4ce32c444cfde4962")

	pasetoToken, err := NewPaseto(key)
	require.NoError(t, err)

	token, err := pasetoToken.NewToken(models.TokenData{
		Subject:  "test",
		Duration: 5 * time.Second,
		AdditionalClaims: models.AdditionalClaims{
			Name: "add name",
			Role: "test-role",
		},
		Footer: models.Footer{MetaData: "footer"},
	})

	require.NoError(t, err)
	require.NotEmpty(t, token)

	sc, err := pasetoToken.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, "footer", sc.Footer.MetaData)
	require.Equal(t, "add name", sc.AdditionalClaims.Name)
	require.Equal(t, "test-role", sc.AdditionalClaims.Role)

}

func TestInvalidCases(t *testing.T) {

	badKey := []byte("00")
	pasetoToken, err := NewPaseto(badKey)
	require.ErrorIs(t, err, ErrInvalidSize)

	key := []byte("000f3e5799296cc4ce32c444cfde4962")
	pasetoToken, err = NewPaseto(key)
	require.NoError(t, err)

	token, err := pasetoToken.NewToken(models.TokenData{
		Duration: -5 * time.Second,
	})
	require.NoError(t, err)

	sc, err := pasetoToken.VerifyToken(token)
	require.Error(t, err)
	require.Error(t, sc.Valid())

}
