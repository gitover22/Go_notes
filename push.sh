#!/bin/bash

# 检查参数数量
if [ $# -ne 1 ]; then
    echo "Usage: $0 <commit_message>"
    exit 1
fi

# 添加所有修改
git add .

# 提交修改并指定提交信息
git commit -m "$1"

# 推送到远程仓库
git push

