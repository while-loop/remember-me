package util

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/while-loop/remember-me/remme/api/services/v1/changer"
	"math/big"
	"os"
)

var (
	stdNums  = []byte("0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")
	stdSpecs = []byte("!@#$%^&*()-_=+,.?/:;{}[]`~")
	stdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
)

type PasswdFunc func() string

var (
	DefaultPasswdFunc = NewPasswordGen(32, true, true).Generate
)

type PasswdGen struct {
	length uint
	chars  []byte
}

func NewPasswordGen(len uint, specialChars, numbers bool) *PasswdGen {
	data := bytes.NewBuffer(stdChars)
	if specialChars {
		data.Write(stdSpecs)
	}

	if numbers {
		data.Write(stdNums)
	}

	return &PasswdGen{len, data.Bytes()}
}

func NewPasswordGenP(config *changer.PasswdConfig) *PasswdGen {
	if config == nil {
		return NewPasswordGen(32, true, true)
	}
	return NewPasswordGen(uint(config.Length), config.SpecialChars, config.Numbers)
}

func (p PasswdGen) Generate() string {
	return rand_char(p.length, p.chars)
}

func rand_char(length uint, chars []byte) string {
	new_pword := make([]byte, length)
	clen := int64(len(chars))
	i := uint(0)
	max := big.NewInt(clen)

	for i = 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			i--
			fmt.Fprint(os.Stderr, err)
		}
		new_pword[i] = chars[n.Int64()%clen]
	}

	return string(new_pword)
}
