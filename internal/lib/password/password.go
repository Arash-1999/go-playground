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
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 12)

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
