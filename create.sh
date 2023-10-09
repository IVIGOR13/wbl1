#!/bin/bash
echo "Bash version ${BASH_VERSION}..."
for i in $(seq 1 1 26)
do
    mkdir "task$i"
    cd "task$i"
    go mod init main
    touch main.go
    echo "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello task [$i]\")\n}" > main.go
    cd ..
    echo "run$i:\n\tgo run task$i/main.go" >> Makefile
done