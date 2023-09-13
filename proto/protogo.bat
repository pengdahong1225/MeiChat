
for %%i in (*.proto) do protoc.exe --proto_path=./ --go_out=../svr/connect/src/proto/ %%i

for %%i in (*.proto) do protoc.exe --proto_path=./ --go_out=../svr/user/src/proto/ %%i

pause