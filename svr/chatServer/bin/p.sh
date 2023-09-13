#!/bin/bash
SvrdName='ChatServer'
ps -ef | grep ${SvrdName} | grep -v "grep"
