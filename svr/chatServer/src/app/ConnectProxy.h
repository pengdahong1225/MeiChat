//
// Created by Messi on 2023/8/28.
//

#ifndef CHATSERVER_CONNECTPROXY_H
#define CHATSERVER_CONNECTPROXY_H

#include <TcpServer/TcpServer.h>

class ConnectProxy {
public:
    static void connectionCallback(const core::net::TcpConnectionPtr &conn);
    static void messageCallback(const core::net::TcpConnectionPtr &conn, const std::string &data);
    static void writeCompleteCallback(const core::net::TcpConnectionPtr &conn);
};

#endif //CHATSERVER_CONNECTPROXY_H
