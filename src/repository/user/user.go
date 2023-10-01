package user

import (
	"app/src/model"
	"app/src/tracing"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"time"

	"gorm.io/gorm"
)

type userRepo struct {
	db        *gorm.DB
	gcm       cipher.AEAD
	memory    uint32
	threads   uint8
	keylen    uint32
	time      uint32
	signKey   *rsa.PrivateKey
	accessExp time.Duration
}

func GetRepository(
	db *gorm.DB,
	secret string,
	memory uint32,
	threads uint8,
	keylen uint32,
	time uint32,
	signKey *rsa.PrivateKey,
	accessExp time.Duration,
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
		db:        db,
		gcm:       gcm,
		memory:    memory,
		threads:   threads,
		keylen:    keylen,
		time:      time,
		signKey:   signKey,
		accessExp: accessExp,
	}, nil
}

func (ur *userRepo) RegisterUser(ctx context.Context, userData model.User) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	if err := ur.db.WithContext(ctx).Create(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (ur *userRepo) CheckRegistered(ctx context.Context, username string) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckRegistered")
	defer span.End()
	var user model.User
	if err := ur.db.WithContext(ctx).Where(model.User{Username: username}).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}

	}

	return user.ID != "", nil
}

func (ur *userRepo) VerifyLogin(ctx context.Context, username, password string, userData model.User) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "VerifyLogin")
	defer span.End()

	if username != userData.Username {
		return false, nil
	}

	verified, err := ur.comparePassword(ctx, userData.Hash, password)
	if err != nil {
		return false, err
	}

	return verified, nil

}

func (ur *userRepo) GetUserData(ctx context.Context, username string) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetUserData")
	defer span.End()

	var user model.User
	if err := ur.db.WithContext(ctx).Where(model.User{Username: username}).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
