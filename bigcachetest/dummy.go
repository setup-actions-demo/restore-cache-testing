// bigcachetest/dummy.go
package main

import (
    _ "github.com/aws/aws-sdk-go-v2/aws"
    _ "cloud.google.com/go"
    _ "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
    _ "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
    _ "github.com/lib/pq"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/golang/protobuf/proto"
    _ "github.com/elastic/go-elasticsearch/v8"
)
func main() {}