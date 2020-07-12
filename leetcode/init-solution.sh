read -p "problem No. ?"
mkdir $REPLY
cd $REPLY
go mod init leetcode
echo "package main" > ans.go