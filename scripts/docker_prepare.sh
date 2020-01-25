#!/bin/sh
set -e
docker build -f Dockefile_qmk_container -t go_qmk_base_container .
