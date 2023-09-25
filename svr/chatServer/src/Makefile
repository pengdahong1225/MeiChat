CXX = g++
CXXFLAGS += -std=c++2a -Wall -g -pipe -rdynamic -fno-strict-aliasing -Wno-unused-function -Wno-sign-compare -fpermissive -Wno-invalid-offsetof

INC += -I. -I../../../core -I../../../proto -I../../../libs \
	-I./app -I./common -I./common/codec -I./common/redisCliPool -I./common/session -I./processer

LIB += -L/usr/local/lib/libprotoc.a ../../../libs/libreactor.a ../../../libs/libprotobuf.a

TARGET = chatServer

CPP_FILES = $(wildcard *.cpp) $(wildcard app/*.cpp)
CPP_FILES += $(wildcard common/session/*.cpp) $(wildcard common/redisClient/*.cpp) $(wildcard common/codec/*.cpp)
CPP_FILES += $(wildcard logic/*.cpp) $(wildcard processer/*.cpp)
OBJ = $(patsubst %.cpp,%.o,$(CPP_FILES))

$(TARGET) : $(OBJ)
	@$(CXX) -o $@ $^ $(LIB)
	install $(TARGET) ../bin/
%.o : %.cpp
	@$(CXX) $(CXXFLAGS) $(INC) -c -o $@ $<
clean:
	@rm -f $(OBJ) $(TARGET) ../bin/$(TARGET)
