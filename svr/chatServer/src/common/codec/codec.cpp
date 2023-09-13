//
// Created by Messi on 2023/9/5.
//

#include "codec.h"

std::string codec::Msg2String(PBHead &header, PBCMsg &msg) {
    std::string data;

    // header
    {
        int length = header.ByteSizeLong();
        std::vector<uint32_t> buffer(length + 4);
        header.SerializeToArray(buffer.data() + 4, length);// 从第5个字节开始写数据
        std::string serialized_data(reinterpret_cast<char *>(buffer.data()), buffer.size());
        *reinterpret_cast<uint32_t *>(serialized_data.data()) = uint32_t(length);// 0~4字节写长度

        data.append(serialized_data);
    }

    // body
    {
        int length = msg.ByteSizeLong();
        std::vector<uint32_t> buffer(length + 4);
        msg.SerializeToArray(buffer.data() + 4, length);// 从第5个字节开始写数据
        std::string serialized_data(reinterpret_cast<char *>(buffer.data()), buffer.size());
        *reinterpret_cast<uint32_t *>(serialized_data.data()) = uint32_t(length);// 0~4字节写长度

        data.append(serialized_data);
    }
    return data;
}

void codec::String2Msg(const std::string &data, PBHead &header, PBCMsg &msg) {

}
