set connect_proto=../svr/connect/src/proto/
if not exist "%connect_proto%" (
   md "%connect_proto%"
)
set user_proto=../svr/user/src/proto/
if not exist "%user_proto%" (
   md "%user_proto%"
)


for %%i in (*.proto) do protoc.exe --proto_path=./ --go_out="%connect_proto%" %%i

for %%i in (*.proto) do protoc.exe --proto_path=./ --go_out="%user_proto%" %%i

pause