package sssa

import (
	"github.com/sethvargo/go-password/password"
	"privileges-management/model"
	"time"
)

const (
	PasswordLength = 27
)

func GenerateCredentials() (model.Credentials, error) {
	username := "tempuser_" + time.Now().String()
	pass, err := password.Generate(PasswordLength, PasswordLength/3, PasswordLength/3, false, true)
	if err != nil {
		return model.Credentials{}, err
	}

	return model.Credentials{
		Username: username,
		Password: pass,
	}, nil
}
