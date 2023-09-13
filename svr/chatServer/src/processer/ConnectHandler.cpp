//
// Created by Messi on 2023/8/25.
//

#include "ConnectHandler.h"
#include "../common/redisCliPool/RedisCliPool.h"

ENHandlerResult ChatP2P::ProcessRequestMsg(const core::net::TcpConnectionPtr &conn, Session *session) {
    auto &request = session->GetRequest().cs_request_chat_single();
    int64_t src = request.src_uid();
    int64_t dst = request.dst_uid();
    // 1.是否是好友
    auto redisConn = RedisCliPool::Instance()->GetConn();
    if (redisConn == nullptr)
        return EN_Handler_Done;
    redisReply *reply;
    reply = redisConn->RedisCommand("select 0");
    if (reply == nullptr || reply->type == REDIS_REPLY_ERROR) {
        std::cout << "ChatP2P::ProcessRequestMsg"
                  << " -> "
                  << "redisConn->RedisVCommand error" << std::endl;
        return EN_Handler_Done;
    }
    freeReplyObject(reply);

    reply = redisConn->RedisVCommand("smembers user:%ld:friends", src);
    if (reply == nullptr || reply->type == REDIS_REPLY_ERROR) {
        std::cout << "ChatP2P::ProcessRequestMsg"
                  << " -> "
                  << "redisConn->RedisVCommand error" << std::endl;
        return EN_Handler_Done;
    }
    bool is_friend = false;
    for (int i = 0; i < reply->elements; i++) {
        if (reply->element[i]->integer == dst) {
            is_friend = true;
            break;
        }
    }
    freeReplyObject(reply);
    if (!is_friend) {
        // response
        CSResponseChatSingle &response = *session->GetResponse().mutable_cs_response_chat_single();
        response.set_content(request.content());
        response.set_content_id(request.content_id());
        response.set_contenttype(request.contenttype());
        response.set_src_uid(request.src_uid());
        response.set_dst_uid(request.dst_uid());
        response.set_result(EN_MESSAGE_ERROR_NO_FRIEND);
        SendToClient(conn, session);
        return EN_Handler_Done;
    }

    // 2.是否在线
    reply = redisConn->RedisCommand("select 1");
    if (reply == nullptr || reply->type == REDIS_REPLY_ERROR) {
        std::cout << "ChatP2P::ProcessRequestMsg"
                  << " -> "
                  << "redisConn->RedisVCommand error" << std::endl;
        return EN_Handler_Done;
    }
    freeReplyObject(reply);

    reply = redisConn->RedisVCommand("sismember user:online %ld", dst);
    if (reply == nullptr || reply->type == REDIS_REPLY_ERROR) {
        std::cout << "ChatP2P::ProcessRequestMsg"
                  << " -> "
                  << "redisConn->RedisVCommand error" << std::endl;
        return EN_Handler_Done;
    }
    if (reply->type != REDIS_REPLY_INTEGER || reply->integer != 1) {
        // dst 不在线
        // response
        CSResponseChatSingle &response = *session->GetResponse().mutable_cs_response_chat_single();
        response.set_content(request.content());
        response.set_content_id(request.content_id());
        response.set_contenttype(request.contenttype());
        response.set_src_uid(request.src_uid());
        response.set_dst_uid(request.dst_uid());
        response.set_result(EN_MESSAGE_ERROR_NO_ONLIONE);
        SendToClient(conn, session);
        return EN_Handler_Done;
    }
    freeReplyObject(reply);

    // response
    CSResponseChatSingle &response = *session->GetResponse().mutable_cs_response_chat_single();
    response.set_content(request.content());
    response.set_content_id(request.content_id());
    response.set_contenttype(request.contenttype());
    response.set_src_uid(request.src_uid());
    response.set_dst_uid(request.dst_uid());
    response.set_result(EN_MESSAGE_ERROR_OK);
    SendToClient(conn, session);
    return EN_Handler_Done;
}

ENHandlerResult ChatP2P::ProcessResponseMsg(const core::net::TcpConnectionPtr &conn, Session *session) {
    return EN_Handler_Done;
}

ENHandlerResult ChatGroup::ProcessRequestMsg(const core::net::TcpConnectionPtr &conn, Session *session) {

}

ENHandlerResult ChatGroup::ProcessResponseMsg(const core::net::TcpConnectionPtr &conn, Session *session) {
    return EN_Handler_Done;
}
