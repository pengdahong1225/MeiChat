#!/bin/bash

cd $(dirname $0)

# db配置
db_host=127.0.0.1
db_port=3306
db_user=root
db_pwd=1225
db_name=meiChat


# 建库
/usr/local/mysql/bin/mysql -h$db_host -P$db_port -u$db_user -p$db_pwd -e "create database if not exists $db_name default character set = utf8mb4;"

# 建表
IN_FILE=meiChat_uid.sql
OUT_FILE=/tmp/meiChat_uid.sql.out

count=0
while [ $count -lt 1 ]
do
	table_index=`printf %03d $count`
	cat $IN_FILE |sed "s/\#db\#/$db_name/g" |sed "s/\#tid\#/$table_index/g" > $OUT_FILE
	/usr/local/mysql/bin/mysql -h$db_host -P$db_port -u$db_user -p$db_pwd -e "source $OUT_FILE;"

	if [ $? -eq 0 ]
	then
		echo "$table_index done."
	else
		echo "$table_index failed."
	fi

	count=$[$count+1]
done

rm $OUT_FILE
