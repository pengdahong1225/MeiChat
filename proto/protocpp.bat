
for %%i in (*.proto) do protoc.exe --proto_path=./ --cpp_out=./ %%i

pause