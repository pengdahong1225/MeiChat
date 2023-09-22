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
PB_OBJ = \
            $(DEP_DIR)/common_msg.o       \
			$(DEP_DIR)/data_msg.o 		\
            $(DEP_DIR)/cs_msg.o 		\
            $(DEP_DIR)/ss_msg.o       \
			$(DEP_DIR)/master_msg.o    \



$(LIB) : $(PB_OBJ)
	ar rcs $@ $^
	
$(PB_OBJ) : $(DEP_DIR)/%.o : %.proto
	$(TOOL_PROTOC) $< --cpp_out=./
	g++ $(CXXFLAGS) $(INC) -c -o $@ $(patsubst %.proto, %.pb.cc, $<)

clean:
	$(RM) $(LIB) $(DEP_DIR) ./*.h ./*.cc ./*.o -rf

.PHONY: protobuf
