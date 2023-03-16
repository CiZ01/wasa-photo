#!/usr/bin/env sh

docker run -it --rm -v "$(pwd):/src" -u "$(id -u):$(id -g)" --network host -p 5173:5000 --workdir /src/webui node:lts /bin/bash
