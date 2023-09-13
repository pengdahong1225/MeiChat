//
// Created by Messi on 2023/8/31.
//

#include "SessionManager.h"

SessionManager::~SessionManager() {
    for (auto &it: _alloc_id_map) {
        if (it.second == nullptr)
            continue;
        delete it.second;
    }
    _alloc_id_map.clear();
}

void SessionManager::Init() {
    _alloc_id_map.clear();
    for (int i = 0; i < MaxSessionID; i++) {
        _free_id_queue.push(i);
    }
}

Session *SessionManager::AllocSession() {
    int id = alloc_sessionID();
    if (id < 0) {
        return nullptr;
    }
    if (_alloc_id_map.find(id) != _alloc_id_map.end()) {
        return nullptr;
    }
    auto *newSession = new Session();
    _alloc_id_map[id] = newSession;
    newSession->SetSessionID(id);
    return newSession;
}

Session *SessionManager::GetSession(int id) {
    if (id < 0) {
        return nullptr;
    }
    if (_alloc_id_map.find(id) == _alloc_id_map.end()) {
        return nullptr;
    }
    return _alloc_id_map[id];
}

void SessionManager::ReleaseSession(int id) {
    if (_alloc_id_map.find(id) != _alloc_id_map.end()) {
        Session *session = _alloc_id_map[id];
        delete session;
        _alloc_id_map[id] = nullptr;
        _alloc_id_map.erase(id);
        _free_id_queue.push(id);
    }
}

int SessionManager::alloc_sessionID() {
    if (_free_id_queue.empty())
        return -1;
    int id = _free_id_queue.front();
    _free_id_queue.pop();
    return id;
}
