package service

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/abdullahjankhan/go_sample/models"
	"golang.org/x/crypto/argon2"
)

type HashingService interface {
	GenerateSalt() (string, error)
	HashWithPepper(plain string) (string, error)
	HashWithSalt(plain, salt string) (string, error)
	SHA256(plain string) (string, error)
	VerifyWithPepper(plain, hash string) (bool, error)
	VerifyPassCode(passCode, hashPassCode, salt string) (bool, error)
}

type hashingService struct {
	format        string
	version       int
	time          uint32
	memory        uint32
	keyLen        uint32
	saltLen       uint32
	threads       uint8
	configService GlobleConfigService
}

func NewHashingService(
	configService GlobleConfigService,
	keyLen uint32,
	memory uint32,
	saltLen uint32,
	threads uint8,
	time uint32,
) HashingService {
	return &hashingService{
		format:        "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		configService: configService,
		keyLen:        keyLen,
		memory:        memory,
		saltLen:       saltLen,
		threads:       threads,
		time:          time,
		version:       argon2.Version,
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func (h *hashingService) GenerateSalt() (string, error) {
	b := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		stdErr := &models.StandardError{
			Code:        models.INTERNAL_SERVER_ERROR,
			ActualError: err,
			Line:        "GenerateSalt():75",
			Message:     models.Hashing_ERROR_MESSAGE,
		}
		return "", stdErr
	}
	secretBase32 := base64.StdEncoding.EncodeToString(b)
	return secretBase32, nil
}

func (h *hashingService) HashWithSalt(plain, salt string) (string, error) {

	hash := argon2.IDKey([]byte(plain), []byte(salt), h.time, h.memory, h.threads, h.keyLen)

	return base64.RawStdEncoding.EncodeToString(hash), nil
}

func (h *hashingService) HashWithPepper(plain string) (string, error) {
	config := h.configService.GetConfig()

	hash := argon2.IDKey([]byte(plain), []byte(config.Aragon.AragonPepper), h.time, h.memory, h.threads, h.keyLen)

	return base64.RawStdEncoding.EncodeToString(hash), nil
}

func (h *hashingService) SHA256(plain string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(plain))
	if err != nil {
		stdErr := &models.StandardError{
			Code:        models.INTERNAL_SERVER_ERROR,
			ActualError: err,
			Line:        "SHA256():103",
			Message:     models.Hashing_ERROR_MESSAGE,
		}
		return "", stdErr
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (h *hashingService) VerifyWithPepper(plain, hash string) (bool, error) {
	config := h.configService.GetConfig()
	decodedHash, err := base64.RawStdEncoding.DecodeString(hash)
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(plain), []byte(config.Aragon.AragonPepper), h.time, h.memory, h.threads, uint32(len(decodedHash)))

	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}

func (u *hashingService) VerifyPassCode(passCode, hashPassCode, salt string) (bool, error) {

	sha256Hash, err := u.SHA256(passCode)
	if err != nil {
		return false, err
	}
	saltedHash, err := u.HashWithSalt(sha256Hash, salt)
	if err != nil {
		return false, err
	}
	pepperedHash, err := u.HashWithPepper(saltedHash)
	if err != nil {
		return false, err
	}

	if pepperedHash != hashPassCode {
		return false, nil
	}
	return true, nil
}
