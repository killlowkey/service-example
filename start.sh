#!/bin/bash
CWD=$(cd "$(dirname $0)";pwd)
root_path=$PWD

echo "$root_path"
echo "$LD_LIBRARY_PATH"

touch /etc/ld.so.conf.d/main.conf
echo "$root_path" > /etc/ld.so.conf.d/main.conf
ldconfig
sleep 1

export LD_LIBRARY_PATH=$root_path:$LD_LIBRARY_PATH
export PATH=$PATH:$root_path

"$CWD"/main install
"$CWD"/main start