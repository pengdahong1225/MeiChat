//
// Created by Messi on 2023/8/28.
//

#include "AppServer.h"
#include "ConnectProxy.h"
#include "ProcesserManager.h"
#include "common/session/SessionManager.h"
#include "common/redisCliPool/RedisCliPool.h"
#include <iostream>

AppServer::AppServer(InetAddr &addr_) {
    loop_ = new core::net::EventLoop;
    server_ = new core::net::TcpServer(loop_, addr_);
    server_->setThreadNum(5);
    server_->setThreadInitCallback([](core::net::EventLoop *) {
        std::cout << "**********初始化成功**********" << std::endl;
    });
    server_->setConnectionCallback(ConnectProxy::connectionCallback);
    server_->setMessageCallback(ConnectProxy::messageCallback);
    server_->setWriteCompleteCallback(ConnectProxy::writeCompleteCallback);
}

AppServer::~AppServer() {
    delete server_;
    delete loop_;
}

void AppServer::start() {
    server_->start();
    loop_->loop();
}

bool AppServer::init() {
    ProcesserManager::Instance()->Init();
    SessionManager::Instance()->Init();
    if (!RedisCliPool::Instance()->Init()) {
        return false;
    }
    return true;
}

int main() {
    InetAddr addr("127.0.0.1", 9000);
    AppServer server(addr);
    if (server.init()) {
        server.start();
    }
    return 0;
}