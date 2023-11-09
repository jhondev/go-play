package models

import "golang.org/x/crypto/bcrypt"

// Create a custom Password type which is a struct containing the plaintext and hashed
// versions of the Password for a user. The plaintext field is a *pointer* to a string,
// so that we're able to distinguish between a plaintext Password not being present in
// the struct at all, versus a plaintext Password which is the empty string "".
type Password struct {
	Plaintext *string
	Hash      []byte
}

// Set calculates the bcrypt hash of a plaintext password, and stores both
// the hash and the plaintext versions in the struct.
func (p *Password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Plaintext = &plaintextPassword
	p.Hash = hash

	return nil
}
