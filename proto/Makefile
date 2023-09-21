DEP_DIR := ./objdep
$(shell mkdir -p $(DEP_DIR))

PROJECT_DIR := ..
LIB_DIR 	:= $(PROJECT_DIR)/libs
TOOL_PROTOC   := /usr/local/bin/protoc

LIB = $(LIB_DIR)/libprotobuf.a

CXXFLAGS += -std=c++2a -Wall -g -pipe -rdynamic -fno-strict-aliasing -Wno-unused-function -Wno-sign-compare -fpermissive

SRC = $(wildcard *.proto)
PB_OBJ = $(DEP_DIR)/$(patsubst %.proto,%.o,$(SRC))

INC_FLAGS =  -I$(PROJECT_DIR)/proto

$(LIB) : $(PB_OBJ) 
	ar rcs $@ $^
	
$(PB_OBJ) : $(SRC)
	$(TOOL_PROTOC) $< --cpp_out=./
	g++ $(CXXFLAGS) $(INC_FLAGS) -c -o $@ $(patsubst %.proto, %.pb.cc, $<)

clean:
	$(RM) $(LIB) $(DEP_DIR) -rf 

.PHONY: protbuf
