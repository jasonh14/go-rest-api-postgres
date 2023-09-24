package user

import (
	"app/internal/model"
	"crypto/aes"
	"crypto/cipher"

	"gorm.io/gorm"
)

type userRepo struct {
	db      *gorm.DB
	gcm     cipher.AEAD
	memory  uint32
	threads uint8
	keylen  uint32
	time    uint32
}

func GetRepository(
	db *gorm.DB,
	secret string,
	memory uint32,
	threads uint8,
	keylen uint32,
	time uint32,
) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &userRepo{
		db:      db,
		gcm:     gcm,
		memory:  memory,
		threads: threads,
		keylen:  keylen,
		time:    time,
	}, nil
}

func (ur *userRepo) RegisterUser(userData model.User) (model.User, error) {
	if err := ur.db.Create(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (ur *userRepo) CheckRegistered(username string) (bool, error) {
	var user model.User
	if err := ur.db.Where(model.User{Username: username}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}

	}

	return user.ID != "", nil
}
