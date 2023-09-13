//
// Created by Messi on 2023/8/31.
//

#include "Session.h"

Session::Session() {
    session_id = -1;
}

PBHead &Session::GetHead() {
    return head_;
}

PBCMsg &Session::GetRequest() {
    return request_;
}

PBCMsg &Session::GetResponse() {
    return response_;
}

void Session::SetHead(PBHead &head) {
    head_ = head;
}

void Session::SetRequest(PBCMsg &request) {
    request_ = request;
}

void Session::SetResponse(PBCMsg &response) {
    response_ = response;
}

int Session::GetSessionID() {
    return session_id;
}

void Session::SetSessionID(int id) {
    session_id = id;
}

ENSessionState Session::GetSessionState() {
    return state_;
}

void Session::SetSessionState(ENSessionState state) {
    state_ = state;
}
