CXXFLAGS += -std=c++11 -Wall -g -pipe -rdynamic -fno-strict-aliasing -Wno-unused-function -Wno-sign-compare -fpermissive -Wno-invalid-offsetof

INC = -I../../../proto -I../../../libs
INC += -I../../../core -I../../../core/Common -I../../../core/Net -I../../../core/Net/Poller -I../../../core/TcpServer
INC += -I./app -I./common -I./common/codec -I./common/redisCliPool -I./common/session -I./processer


LIB = -lpthread
LIB += -L/usr/local/lib/libprotoc.a ../../../libs/libreactor.a ../../../libs/libprotobuf.a

TARGET = chatServer

CPP_FILES = $(wildcard *.cpp) $(wildcard app/*.cpp)
CPP_FILES += $(wildcard common/session/*.cpp) $(wildcard common/redisClient/*.cpp) $(wildcard common/codec/*.cpp)
CPP_FILES += $(wildcard logic/*.cpp) $(wildcard processer/*.cpp)
OBJ = $(patsubst %.cpp,%.o,$(CPP_FILES))

all : $(TARGET)
include makefile.share
