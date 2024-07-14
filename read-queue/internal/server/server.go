package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"readq/internal/utils"

	"github.com/go-redis/redis/v8"
)

// Remove the declaration of the unused variable
var StreamData = make(chan []byte, 1000)

func StringBase64ToByteArray(s string) []byte {
	b := make([]byte, len(s))
	for i := range s {
		b[i] = s[i]
	}
	return b
}

func GetDataFromStreamInfinite() {

	rdb := utils.GetRDB()
	ctx := context.Background()
	lastID := "0-0" // Start from the beginning of the stream

	go Run(StreamData)

	for {
		// Use XRead in a blocking mode, waiting for new messages
		streams, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"streamExporter", lastID},
			Count:   100,
			Block:   1000,
		}).Result()

		if err != nil {
			fmt.Printf("Error reading from stream: %v\n", err)
			continue
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				data, ok := message.Values["data"].(string)
				if !ok {
					fmt.Println("Error converting data to string")
					continue
				}

				// Decode the Base64 string to a byte array
				byteArray, err := base64.StdEncoding.DecodeString(data)
				if err != nil {
					fmt.Printf("Error decoding Base64 string: %v\n", err)
					continue
				}

				StreamData <- byteArray

				_, err = rdb.XDel(ctx, "streamExporter", message.ID).Result()
				if err != nil {
					fmt.Printf("Error deleting message %s: %v\n", message.ID, err)
				}
				lastID = message.ID // Update lastID to the latest processed message ID
			}
		}

	}
}
