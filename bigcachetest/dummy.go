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
    _ "go.uber.org/zap"
    _ "github.com/prometheus/client_golang/prometheus"
    _ "github.com/gin-gonic/gin"
    _ "github.com/go-redis/redis/v8"
    _ "github.com/gorilla/websocket"
    _ "github.com/sirupsen/logrus"
    _ "github.com/spf13/cobra"
    _ "github.com/spf13/viper"
    _ "go.mongodb.org/mongo-driver/mongo"
    _ "github.com/google/uuid"
    _ "github.com/minio/minio-go/v7"
    _ "k8s.io/client-go/kubernetes"
    _ "github.com/apache/thrift/lib/go/thrift"
    _ "github.com/apache/arrow/go/arrow"
    _ "github.com/nats-io/nats.go"
)
func main() {}