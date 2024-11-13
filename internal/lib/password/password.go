package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	plain *string
	Hash  []byte
}

func (p *Password) Set(pass string) error {
	// TODO: manage password hash before bcrypt
	// sha := sha256.New()
	// salt := make([]byte, saltSize)
	// _, err := rand.Read(salt[:])
	//
	// if err != nil {
	// 	panic(err)
	// }
	// passBytes := []byte(pass)
	// passBytes = append(passBytes, salt...)
	//
	// sha.Write(passBytes)
	// hash, err := bcrypt.GenerateFromPassword(sha.Sum(nil), bcrypt.DefaultCost)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	p.Hash = hash
	p.plain = &pass

	return nil
}

func (p *Password) Match(pass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(pass))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
