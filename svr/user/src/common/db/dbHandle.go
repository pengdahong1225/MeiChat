package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"user/src/common"
	pb "user/src/proto"
)

// 写
func SetData(uid int64, user *pb.PBUser) error {
	sqlInstance := NewSqlInstance()
	defer sqlInstance.Close()

	str := "insert into meiChat(uid,account,pwd,gender,pic) values (?,?,?,?,?)"
	ret, e_exec := sqlInstance.Exec(str, uid, user.Account, user.Pwd, user.Gender, user.PicUrl)
	if e_exec != nil {
		return e_exec
	}
	lastInsertID, e_last := ret.LastInsertId()
	if e_last != nil {
		return e_last
	}
	fmt.Println("Inserted ID:", lastInsertID)

	// update redis
	if err := updateRedis(uid, user); err != nil {
		return err
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////

// 读
func GetData(uid int64) *pb.PBUser {
	conn := RedisPoolInstance.Get()
	defer conn.Close()
	// 身份验证
	_, err := conn.Do("AUTH", "1225")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	param := &common.UserInfo{
		Uid: uid,
	}
	if readByCache(uid, conn, param) == nil {
		return &pb.PBUser{
			Uid:     param.Uid,
			Account: param.Account,
			Pwd:     param.Pwd,
			Gender:  pb.ENGender(param.Gender),
			PicUrl:  param.Pic,
		}
	}
	if readByDB(uid, param) != nil {
		user := &pb.PBUser{
			Uid:     param.Uid,
			Account: param.Account,
			Pwd:     param.Pwd,
			Gender:  pb.ENGender(param.Gender),
			PicUrl:  param.Pic,
		}
		// 同步到redis
		if updateRedis(uid, user) != nil {
			return nil
		}
		return user
	}

	return nil
}

func readByCache(uid int64, conn redis.Conn, param *common.UserInfo) (err error) {
	param.Account, err = redis.String(conn.Do("HGet", "user:"+strconv.FormatInt(uid, 10), "account"))
	if err != nil {
		return err
	}
	param.Pwd, err = redis.String(conn.Do("HGet", "user:"+strconv.FormatInt(uid, 10), "pwd"))
	if err != nil {
		return err
	}
	param.Gender, err = redis.Int(conn.Do("HGet", "user:"+strconv.FormatInt(uid, 10), "gender"))
	if err != nil {
		return err
	}
	param.Pic, err = redis.String(conn.Do("HGet", "user:"+strconv.FormatInt(uid, 10), "pic"))
	if err != nil {
		return err
	}
	return nil
}
func readByDB(uid int64, param *common.UserInfo) (err error) {
	sqlInstance := NewSqlInstance()
	defer sqlInstance.Close()

	str := "select * from meiChat where uid = ?"
	err = sqlInstance.QueryRow(str, uid).Scan(param.Uid, param.Account, param.Pwd, param.Gender, param.Pic)
	return
}

func updateRedis(uid int64, user *pb.PBUser) error {
	conn := RedisPoolInstance.Get()
	defer conn.Close()
	var err error
	// 身份验证
	_, err = conn.Do("AUTH", "1225")
	if err != nil {
		return err
	}

	_, err = conn.Do("HSet", "user:"+strconv.FormatInt(uid, 10), "account", user.Account,
		"pwd", user.Pwd, "gender", user.Gender, "pic", user.PicUrl)
	if err != nil {
		return err
	}
	return nil
}
