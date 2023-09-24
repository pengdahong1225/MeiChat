//
// Created by Messi on 2023/9/5.
//

#ifndef CHATSERVER_CODEC_H
#define CHATSERVER_CODEC_H

#include "../../../proto/master_msg.pb.h"
#include <string>

class codec {
public:
    static std::string Msg2String(PBHead &header, PBCMsg &msg);
    static void String2Msg(const std::string &data, PBHead &header, PBCMsg &msg);
};


#endif //CHATSERVER_CODEC_H
