#!/bin/bash
CWD=$(cd "$(dirname $0)";pwd)
"$CWD"/main stop
"$CWD"/main uninstall