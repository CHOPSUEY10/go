package auth

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

var path string = "../account.jsonl"
var pathindex string = "../accountindex.jsonl"

type Account struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
	Password *HashSalt `json:"password"`
	Mut      sync.RWMutex
}

// Penyimpanan password dan compare untuk login
type HashSalt struct {
	Hash []byte
	Salt []byte
}

type Index struct {
	Username string
	Line     int
}

const (
	argon2Time    = 3
	argon2Memory  = 64 * 1024
	argon2Threads = 4
	argon2KeyLen  = 32
	argon2SaltLen = 16
)

type Argon2idHash struct {

	// time represents the number of

	// passed over the specified memory.

	time uint32

	// cpu memory to be used.

	memory uint32

	// threads for parallelism aspect

	// of the algorithm.

	threads uint8

	// keyLen of the generate hash key.

	keyLen uint32

	// saltLen the length of the salt used.

	saltLen uint32
}

func NewArgon2idHash() *Argon2idHash {
	return &Argon2idHash{
		time:    argon2Time,
		saltLen: argon2SaltLen,
		memory:  argon2Memory,
		threads: argon2Threads,
		keyLen:  argon2KeyLen,
	}
}

func (a *Argon2idHash) GenerateHash(password, salt []byte) (*HashSalt, error) {

	var err error

	// If salt is not provided generate a salt of

	// the configured salt length.

	if len(salt) == 0 {

		salt, err = randomSecret(a.saltLen)

	}

	if err != nil {

		return nil, err

	}

	// Generate hash

	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)

	// Return the generated hash and salt used for storage.

	return &HashSalt{Hash: hash, Salt: salt}, nil

}

func randomSecret(length uint32) ([]byte, error) {

	secret := make([]byte, length)

	_, err := rand.Read(secret)

	if err != nil {

		return nil, err

	}

	return secret, nil

}

func CreateAccount(username, password string) *Account {
	hash := NewArgon2idHash()
	salt, err := randomSecret(argon2SaltLen)
	if err != nil {
		return nil
	}

	hashedPassword, err := hash.GenerateHash([]byte(password), salt)
	if err != nil {
		return nil
	}

	account := &Account{
		Id:       uuid.NewString(),
		Username: username,
		Password: hashedPassword,
	}
	return account
}

// // FIX IT LATER AFTER ReadAccount function
func (account *Account) ValidatePassword(ctx context.Context) error {

	hashAdapter := NewArgon2idHash()

	hashed, err := hashAdapter.GenerateHash([]byte(ctx.Value(account).(string)), []byte(account.Password.Salt))
	if err != nil {
		return err
	}

	savedHash := account.Password.Hash
	if string(savedHash) == string(hashed.Hash) {
		return nil
	} else {
		return errors.New("Wrong Password")
	}
}

func (acc *Account) SaveAccount() error {
	acc.Mut.Lock()
	defer acc.Mut.Unlock()

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return fmt.Errorf("Gagal membuka file : %w", err)
	}
	defer file.Close()

	data, err := json.Marshal(acc)
	if err != nil {
		return fmt.Errorf("Gagal konversi ke json : %w", err)
	}

	data = append(data, '\n')

	_, err = file.Write(data)

	if err != nil {
		return fmt.Errorf("Gagal menulis ke JSON: %w", err)
	}

	return acc.makeIndex()
}
