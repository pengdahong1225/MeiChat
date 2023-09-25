//
// Created by Messi on 2023/8/28.
//

#include "ConnectProxy.h"
#include "ProcesserManager.h"
#include "common/session/SessionManager.h"
#include "common/codec/codec.h"

void ConnectProxy::connectionCallback(const core::net::TcpConnectionPtr &conn) {
    // add

}

void ConnectProxy::messageCallback(const core::net::TcpConnectionPtr &conn, const std::string &data) {
    PBHead header{};
    PBCMsg msg{};
    // 解包
    codec::String2Msg(data, header, msg);

    // 处理
    ProcessBase *process = ProcesserManager::Instance()->GetProcess(header.cmd());
    if (process == nullptr) {
        std::cout << "ConnectProxy::messageCallback"
                  << " -> "
                  << "GetProcess error" << std::endl;
    }
    // 获取session
    auto session_ = SessionManager::Instance()->GetSession(header.session_id());
    if (session_ == nullptr) {
        // 分配session
        session_ = SessionManager::Instance()->AllocSession();
        if (session_ == nullptr) {
            std::cout << "ConnectProxy::messageCallback"
                      << " -> "
                      << "AllocSession error" << std::endl;
            return;
        }
        // 如果是新的session，要设置state
        session_->SetSessionID(EN_Session_Idle);
    }
    // 更新session
    header.set_session_id(session_->GetSessionID());
    session_->SetHead(header);
    session_->SetRequest(msg);
    process->Process(conn, session_);
}

void ConnectProxy::writeCompleteCallback(const core::net::TcpConnectionPtr &conn) {

}
