#!/bin/bash

table=$1



goctl model mysql datasource --dir ./internal/model --table $table --cache true --url "root:PXDN93VRKUm8TeE7@tcp(0.0.0.0:33069)/beyond_user"