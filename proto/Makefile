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
# .cc
MID = $(patsubst %.proto,%.pb.cc,$(SRC))
# .o
PB_OBJ = $(DEP_DIR)/$(patsubst %.proto,%.o,$(SRC))


$(LIB) : $(PB_OBJ)
	ar rcs $@ $^
	
$(PB_OBJ) : $(MID)
	g++ $(CXXFLAGS) $(INC_FLAGS) -c -o $@ $^

$(MID) : $(SRC)
	$(TOOL_PROTOC) $^ --cpp_out=./

clean:
	$(RM) $(LIB) $(DEP_DIR) ./*.h ./*.cc ./*.o -rf

.PHONY: protbuf
