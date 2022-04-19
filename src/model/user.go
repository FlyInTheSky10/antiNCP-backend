package model

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/FlyInThesky10/antiNCP-backend/controller/param"
	"github.com/go-redis/redis/v7"
	"time"
)

const (
	timePrivateKey = time.Minute * 15 // 15 min
	keyPrivateKey  = "priv_key:%s"
)

func GetUserByID(id string) (param.User, bool, error) {
	user := make([]param.User, 0)
	err := Db.Select(&user, fmt.Sprintf("SELECT * FROM user WHERE id=%s", id))
	if err != nil {
		return param.User{}, false, err
	}
	if len(user) < 1 {
		return param.User{}, false, nil
	}
	return user[0], true, nil
}
func AddUser(user param.ReqUserRegister) error {
	tx := Db.MustBegin()
	tx.MustExec(
		"INSERT INTO user(id, password, name, academy, id_number, phone_number, admin)values(?, ?, ?, ?, ?, ?, ?)",
		user.Id,
		user.Password,
		user.Name,
		user.Academy,
		user.IdNumber,
		user.PhoneNumber,
		0,
	)
	tx.MustExec(
		"INSERT INTO code(id, status, verify_id, verify_time)values(?, ?, ?, ?)",
		user.Id, 0, "", 0,
	)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func GetPrivateKey(id string) (*rsa.PrivateKey, bool, error) {
	result := redisClient.Get(fmt.Sprintf(keyPrivateKey, id))
	if result.Err() == redis.Nil {
		return &rsa.PrivateKey{}, false, nil
	}
	if result.Err() != nil {
		return &rsa.PrivateKey{}, false, result.Err()
	}
	key, err := x509.ParsePKCS1PrivateKey([]byte(result.Val()))
	if err != nil {
		return &rsa.PrivateKey{}, false, err
	}
	return key, true, nil
}
func AddPrivateKey(id string, key *rsa.PrivateKey) error {
	return redisClient.Set(fmt.Sprintf(keyPrivateKey, id),
		x509.MarshalPKCS1PrivateKey(key), timePrivateKey).Err()
}
