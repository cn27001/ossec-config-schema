#!/bin/bash

printf "** Building linux/amd64\n"
go build -a github.com/sedarasecurity/ossec/types
go install github.com/sedarasecurity/ossec/types

printf "** Building windows/386\n"
go-windows-386 build -a github.com/sedarasecurity/ossec/types
go-windows-386 install github.com/sedarasecurity/ossec/types

printf "** Building windows/amd64\n"
go-windows-amd64 build -a github.com/sedarasecurity/ossec/types
go-windows-amd64 install github.com/sedarasecurity/ossec/types
