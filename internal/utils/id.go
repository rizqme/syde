package utils

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand/v2"
	"strings"
	"time"

	"github.com/feedloop/syde/internal/model"
)

const ulidChars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

// GenerateID creates a ULID-like ID with the entity kind prefix.
func GenerateID(kind model.EntityKind) string {
	return kind.IDPrefix() + "_" + generateULID()
}

func generateULID() string {
	// Time component (10 chars, millisecond precision)
	t := time.Now().UnixMilli()
	timePart := encodeTime(t, 10)

	// Random component (16 chars)
	randPart := encodeRandom(16)

	return timePart + randPart
}

func encodeTime(t int64, length int) string {
	b := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		b[i] = ulidChars[t%32]
		t /= 32
	}
	return string(b)
}

func encodeRandom(length int) string {
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(32))
		if err != nil {
			b[i] = ulidChars[mrand.IntN(32)]
		} else {
			b[i] = ulidChars[n.Int64()]
		}
	}
	return string(b)
}

// ParseIDKind extracts the entity kind from an ID prefix.
func ParseIDKind(id string) (model.EntityKind, bool) {
	parts := strings.SplitN(id, "_", 2)
	if len(parts) != 2 {
		return "", false
	}
	for _, k := range model.AllEntityKinds() {
		if k.IDPrefix() == parts[0] {
			return k, true
		}
	}
	return "", false
}
