//
// Created by Messi on 2023/8/28.
//

#ifndef CHATSERVER_APPSERVER_H
#define CHATSERVER_APPSERVER_H

#include "../../../../core/TcpServer.h"
#include <EventLoop.h>

class AppServer {
public:
    AppServer(InetAddr &addr_);
    ~AppServer();
    bool init();
    void start();

private:
    core::net::EventLoop *loop_;
    core::net::TcpServer *server_;
};

#endif //CHATSERVER_APPSERVER_H
