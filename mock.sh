#!/bin/bash

mockery --dir pkg/domain/ --name PostRepository --output internal/post/repository/mock/ --case underscore --outpkg mock --with-expecter