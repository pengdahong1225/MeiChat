//
// Created by Messi on 2023/8/31.
//

#ifndef CHATSERVER_SESSIONMANAGER_H
#define CHATSERVER_SESSIONMANAGER_H

#include <queue>
#include <map>
#include "Common/singleton.h"
#include "Session.h"

#define MaxSessionID 100

class SessionManager : public CSingleton<SessionManager> {
public:
    ~SessionManager();
    void Init();
    Session *AllocSession();
    Session *GetSession(int id);
    void ReleaseSession(int id);

private:
    int alloc_sessionID();

private:
    std::queue<int> _free_id_queue;
    std::map<int, Session *> _alloc_id_map;
};

#endif //CHATSERVER_SESSIONMANAGER_H
