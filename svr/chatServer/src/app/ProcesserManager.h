//
// Created by Messi on 2023/8/25.
//

#ifndef CHATSERVER_PROCESSERMANAGER_H
#define CHATSERVER_PROCESSERMANAGER_H

#include <unordered_map>
#include "common/session/Session.h"
#include "Common/Callbacks.h"
#include "TcpConnection.h"
#include "Common/singleton.h"

#define REGIST_MSG_HANDLER(cmd, handler) \
do { \
    ProcessBase* phandler = new handler(); \
    handlerMap_[cmd] = phandler; \
} while(0)

enum ENHandlerResult {
    EN_Handler_Done = 1,
    EN_Handler_Succ = 2,
    EN_Handler_Get = 3,
    EN_Handler_Save = 4,
};

class ProcessBase {
public:
    ProcessBase() = default;
    virtual ~ProcessBase() = default;

    // 状态机
    void Process(const core::net::TcpConnectionPtr &conn, Session *psession);
    virtual ENHandlerResult ProcessRequestMsg(const core::net::TcpConnectionPtr &, Session *) = 0;
    virtual ENHandlerResult ProcessResponseMsg(const core::net::TcpConnectionPtr &, Session *) = 0;

protected:
    static void SendToClient(const core::net::TcpConnectionPtr &conn, Session *psession);

private:
    static void EndProcess(const core::net::TcpConnectionPtr &conn, Session *psession);
};

typedef std::unordered_map<int, ProcessBase *> HandlerMap;

class ProcesserManager : public CSingleton<ProcesserManager> {
public:
    void Init();
    ProcessBase *GetProcess(int cmd);

private:
    HandlerMap handlerMap_;
};

#endif //CHATSERVER_PROCESSERMANAGER_H
