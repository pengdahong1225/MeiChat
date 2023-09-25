//
// Created by Messi on 2023/8/25.
//

#ifndef CHATSERVER_CONNECTHANDLER_H
#define CHATSERVER_CONNECTHANDLER_H

#include "app/ProcesserManager.h"

class ChatP2P : public ProcessBase {
public:
    ENHandlerResult ProcessRequestMsg(const core::net::TcpConnectionPtr &conn, Session *session) override;
    ENHandlerResult ProcessResponseMsg(const core::net::TcpConnectionPtr &conn, Session *session) override;

private:
    bool CheckFriendShip(int64_t src, int64_t dst);
};

class ChatGroup : public ProcessBase {
    ENHandlerResult ProcessRequestMsg(const core::net::TcpConnectionPtr &conn, Session *session) override;
    ENHandlerResult ProcessResponseMsg(const core::net::TcpConnectionPtr &conn, Session *session) override;
};

#endif //CHATSERVER_CONNECTHANDLER_H
