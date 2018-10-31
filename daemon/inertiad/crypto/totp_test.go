package crypto

import (
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGeneration(t *testing.T) {
	key, err := generateSecretKey("TestAccountName")
	assert.Nil(t, err)

	// string form of secret
	secret := key.Secret()
	assert.Equal(t, len(secret), 16)

	// img form of secret
	_, err = key.Image(400, 400)
	assert.Nil(t, err)
}

func TestVerification(t *testing.T) {

	key, err := generateSecretKey("TestAccountName")
	assert.Nil(t, err)
	currentTime := time.Now()

	verificationTests := []struct {
		name string
		in   time.Time
		out  bool
	}{
		{
			"valid TOTP",
			currentTime,
			true,
		},
		{
			"TOTP before current period window",
			currentTime.Add(time.Duration(-(TotpPeriod * 2) * time.Second)),
			false,
		},
		{
			"TOTP after current period window",
			currentTime.Add(time.Duration((TotpPeriod * 2) * time.Second)),
			false,
		},
	}

	for _, test := range verificationTests {
		t.Run(test.name, func(t *testing.T) {
			code, err := totp.GenerateCode(key.Secret(), test.in)
			assert.Nil(t, err)
			assert.Equal(t, validatePasscode(code, key.Secret()), test.out)
		})

	}
}

func TestBackupCodes(t *testing.T) {
	codes := generateBackupCodes()
	assert.Equal(t, len(codes), TotpNoBackupCodes)
	for _, code := range codes {
		assert.Equal(t, len(code), 11)
	}
}