#!/bin/sh
SvrdName='ChatServer'
BIN=${SvrdName}
IP=127.0.0.1
INNER=127.0.0.1
PORT=9000

./$BIN --host=$IP --inner=$INNER --port=$PORT