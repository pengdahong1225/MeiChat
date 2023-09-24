//
// Created by Messi on 2023/9/1.
//

#ifndef CHATSERVER_CONFIGREDISCLIENT_H
#define CHATSERVER_CONFIGREDISCLIENT_H

#include "../../../../core/Common/CRedisServer.h"
#include "../../../../core/Common/singleton.h"
#include <mutex>

#define MAXCliConn 10

/*
 * redis 连接池
 * 单例
 */
class RedisCliPool : public CSingleton<RedisCliPool> {
public:
    ~RedisCliPool();
    bool Init();
    CRedisServer *GetConn();// 获取空闲连接
    void PutConn(CRedisServer *conn);// 归还连接
private:
    std::vector<CRedisServer> pool_;
    std::mutex mtx_;
};

#endif //CHATSERVER_CONFIGREDISCLIENT_H
