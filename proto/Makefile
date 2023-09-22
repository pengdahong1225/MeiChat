DEP_DIR := ./objdep
$(shell mkdir -p $(DEP_DIR))

LIB_DIR := $(PROJECT_DIR)/libs
TOOL_PROTOC := /usr/local/bin/protoc

CXXFLAGS += -std=c++11 -Wall -g -pipe -rdynamic -fno-strict-aliasing -Wno-unused-function -Wno-sign-compare -fpermissive

LIB = $(LIB_DIR)/libprotobuf.a

# .proto
SRC = $(wildcard *.proto)
# .h
INC = -I./proto
# .o
PB_OBJ = $(DEP_DIR)/$(patsubst %.proto,%.o,$(SRC))


$(LIB) : $(PB_OBJ)
	ar rcs $@ $^
	
%.o : %.cc
	g++ $(CXXFLAGS) $(INC) -c -o $@ $^

%.cc : %.proto
	$(TOOL_PROTOC) $^ --cpp_out=./

clean:
	$(RM) $(LIB) $(DEP_DIR) ./*.h ./*.cc ./*.o -rf

.PHONY: protobuf
