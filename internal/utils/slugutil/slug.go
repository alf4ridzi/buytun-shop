package slugutil

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/gosimple/slug"
)

func generateSecureRandom() (int64, error) {
	max := new(big.Int).SetInt64(1 << 62)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

func GenerateProductSlug(name string) string {
	text := slug.Make(name)

	randScure, err := generateSecureRandom()
	if err != nil {
		return ""
	}

	secureRandom := strconv.FormatInt(randScure, 10)
	return text + "-" + secureRandom
}
