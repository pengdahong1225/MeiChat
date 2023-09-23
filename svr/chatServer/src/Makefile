CXXFLAGS += -std=c++11 -Wall -g -pipe -rdynamic -fno-strict-aliasing -Wno-unused-function -Wno-sign-compare -fpermissive -Wno-invalid-offsetof

INC += -I. -I../../../core -I../../../core/Common -I../../../core/Net -I../../../core/Net/Poller -I../../../core/TcpServer
INC += -I../../../proto -I../../../libs

LIB = -lpthread
LIB += -L/usr/local/lib/libprotoc.a ../../../libs/libreactor.a ../../../libs/libprotobuf.a

TARGET = chatServer

CPP_FILES = $(wildcard *.cpp) $(wildcard app/*.cpp)
CPP_FILES += $(wildcard common/session/*.cpp) $(wildcard common/redisClient/*.cpp) $(wildcard common/codec/*.cpp)
CPP_FILES += $(wildcard logic/*.cpp) $(wildcard processer/*.cpp)
OBJ = $(patsubst %.cpp,%.o,$(CPP_FILES))

OBJ_DIR = .dep
SUBDIRS=$(shell ls -l | grep ^d | awk '{print $$9}')
$(shell if [ ! -d $(OBJ_DIR) ]; then for dir in $(SUBDIRS); do mkdir -p $(OBJ_DIR)/$$dir; done fi)

OBJ = $(addprefix $(OBJ_DIR)/, $(patsubst %.cpp,%.o,$(CPP_FILES)))

all : $(TARGET)
include makefile.share
