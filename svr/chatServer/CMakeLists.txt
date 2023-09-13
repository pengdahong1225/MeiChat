cmake_minimum_required(VERSION 3.24)

set(TARGET chatServer)
set(CMAKE_CXX_STANDARD 20)

if (WIN32)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /Zc:__cplusplus")
endif ()

project(${TARGET})

include_directories(src)
include_directories(../../core)
include_directories(../../proto)
include_directories(../../libs)

aux_source_directory(src/app SRCS)
aux_source_directory(src/common SRCS)
aux_source_directory(src/common/codec SRCS)
aux_source_directory(src/common/session SRCS)
aux_source_directory(src/common/redisCliPool SRCS)
aux_source_directory(src/processer SRCS)
aux_source_directory(src/logic SRCS)

add_executable(${TARGET} ${SRCS})