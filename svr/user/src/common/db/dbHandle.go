package db

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"user/src/common"
	pb "user/src/proto"
)

// 增
func AddData(uid int64, user *pb.PBUser) error {
	sqlConn := SqlInstance()
	defer sqlConn.Close()

	str := "insert into meiChat(uid,account,pwd,gender,pic) values (?,?,?,?,?)"
	ret, e_exec := sqlConn.Exec(str, uid, user.Account, user.Pwd, user.Gender, user.PicUrl)
	if e_exec != nil {
		return e_exec
	}

	lastInsertID, e_last := ret.LastInsertId() // 新插入数据的id
	if e_last != nil {
		return e_last
	}
	log.Println("Inserted ID:", lastInsertID)

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
		log.Println(err)
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
	sqlConn := SqlInstance()
	// 查询
	sqlStr := "select * from meiChat where uid = ?"
	rows, e := sqlConn.Query(sqlStr, uid)
	if e != nil {
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据 且调用Scan才会释放我们的连接
	for rows.Next() {
		err = rows.Scan(param.Uid, param.Account, param.Pwd, param.Gender, param.Pic)
		if err != nil {
			return
		}
	}

	// 也可查询和读取一起 QueryRow最多返回一条数据
	// err = sqlConn.QueryRow(str, uid).Scan(param.Uid, param.Account, param.Pwd, param.Gender, param.Pic)
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
