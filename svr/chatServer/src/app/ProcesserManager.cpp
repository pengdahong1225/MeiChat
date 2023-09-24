//
// Created by Messi on 2023/8/25.
//

#include "../../../../proto/master_msg.pb.h"
#include "ProcesserManager.h"
#include "../common/session/SessionManager.h"
#include "../processer/ConnectHandler.h"
#include "../common/codec/codec.h"

void ProcesserManager::Init() {
    REGIST_MSG_HANDLER(PBCMsg::kCsRequestChatSingle, ChatP2P);
    REGIST_MSG_HANDLER(PBCMsg::kCsRequestChatGroup, ChatGroup);
}

ProcessBase *ProcesserManager::GetProcess(int cmd) {
    if (handlerMap_.find(cmd) != handlerMap_.end()) {
        return handlerMap_[cmd];
    }
    return nullptr;
}

void ProcessBase::Process(const core::net::TcpConnectionPtr &conn, Session *psession) {
    ENHandlerResult result = EN_Handler_Done;
    if (psession->GetHead().mtype() == EN_Message_Request) {
        psession->SetSessionState(EN_Session_Idle);
        result = ProcessRequestMsg(conn, psession);
    } else if (psession->GetHead().mtype() == EN_Message_Response) {
        switch (psession->GetSessionState()) {
            case EN_Session_Idle:
                std::cout << "ProcessBase::Process"
                          << " -> "
                          << "response too late" << std::endl;
                result = EN_Handler_Done;
                break;
            case EN_Session_Wait_Get_Data:
                result = ProcessResponseMsg(conn, psession);
                break;
            default:
                break;
        }
    }
    if (result == EN_Handler_Done) {
        EndProcess(conn, psession);
    }
}

void ProcessBase::EndProcess(const core::net::TcpConnectionPtr &conn, Session *psession) {
    SessionManager::Instance()->ReleaseSession(psession->GetSessionID());
}

void ProcessBase::SendToClient(const core::net::TcpConnectionPtr &conn, Session *psession) {
    // 组包
    PBHead header{};
    header.CopyFrom(psession->GetHead());
    header.set_mtype(EN_Message_Response);// reverse
    std::string data = codec::Msg2String(header, psession->GetResponse());

    conn->send(data);
}
