//
// Created by Messi on 2023/8/31.
//

#ifndef CHATSERVER_SESSION_H
#define CHATSERVER_SESSION_H

#include "../../../../../proto/master_msg.pb.h"
#include <Common/noncopyable.h>

enum ENSessionState {
    EN_Session_Idle = 1,
    EN_Session_Wait_Get_Data = 2,
};

class Session : noncopyable {
public:
    Session();
    ~Session() = default;

    int GetSessionID();
    PBHead &GetHead();
    PBCMsg &GetRequest();
    PBCMsg &GetResponse();
    ENSessionState GetSessionState();

    void SetSessionID(int id);
    void SetHead(PBHead &head);
    void SetRequest(PBCMsg &request);
    void SetResponse(PBCMsg &response);
    void SetSessionState(ENSessionState state);

private:
    PBHead head_;
    PBCMsg request_;
    PBCMsg response_;

    int session_id;
    ENSessionState state_;
};

#endif //CHATSERVER_SESSION_H
