package encryption

import (
	"crypto/rand"
	"log"
	"math/big"
	"strconv"
)

// getRandNum returns a otp number of size four
func GetOtpNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	log.Println("otp number", nBig)
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}
