//
// Created by Messi on 2023/9/1.
//

#include "RedisCliPool.h"
#include <iostream>

// Debug
const std::string ip = "127.0.0.1";
const int port = 6379;
const std::string passwd = "1225";

bool RedisCliPool::Init() {
    pool_.clear();
    for (int i = 0; i < MAXCliConn; i++) {
        pool_.emplace_back();
    }

    for (auto &conn: pool_) {
        if (conn.Connect(ip, port) != 0) {
            std::cout << "RedisCliPool::Init"
                      << " -> "
                      << " Connect error" << std::endl;
            return false;
        }
        if (conn.RedisVCommand("auth %s", passwd.c_str()) == nullptr) {
            std::cout << "RedisCliPool::Init"
                      << " -> "
                      << " AUTH Passwd error" << std::endl;
            return false;
        }
    }
    return true;
}

RedisCliPool::~RedisCliPool() {
    pool_.clear();
}

CRedisServer *RedisCliPool::GetConn() {
    CRedisServer *freeConn = nullptr;
    mtx_.lock();
    if (pool_.empty()) {
        return nullptr;
    }
    // 获取空闲的redis连接
    for (auto &conn: pool_) {
        if (!conn.in_use_) {
            freeConn = &conn;
            freeConn->in_use_ = true;
        }
    }
    mtx_.unlock();
    return freeConn;
}

void RedisCliPool::PutConn(CRedisServer *conn) {
    mtx_.lock();
    conn->ResetConn();
    conn->in_use_ = false;
    mtx_.unlock();
}

RedisCliPool::RedisCliPool() {

}
