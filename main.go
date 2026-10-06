package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func input(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("================================")
	fmt.Println("       Go S3 Automation")
	fmt.Println("================================")

	accessKey := input(reader, "AWS Access Key: ")
	secretKey := input(reader, "AWS Secret Key: ")
	region := input(reader, "AWS Region [us-east-1]: ")

	if region == "" {
		region = "us-east-1"
	}

	creds := credentials.NewStaticCredentialsProvider(
		accessKey,
		secretKey,
		"",
	)

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return
	}

	client := s3.NewFromConfig(cfg)

	fmt.Println("\n1. List Buckets")
	fmt.Println("2. List Objects")
	fmt.Println("3. Upload File")
	fmt.Println("4. Download File")
	fmt.Println("5. Delete Object")

	choice := input(reader, "\nChoose option: ")

	ctx := context.Background()

	switch choice {

	case "1":
		result, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("\nS3 Buckets:")
		for _, bucket := range result.Buckets {
			fmt.Println("-", aws.ToString(bucket.Name))
		}

	case "2":
		bucket := input(reader, "Bucket name: ")

		result, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket: aws.String(bucket),
		})

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("\nObjects:")
		for _, obj := range result.Contents {
			fmt.Println("-", aws.ToString(obj.Key))
		}

	case "3":
		bucket := input(reader, "Bucket name: ")
		file := input(reader, "Local file path: ")
		key := input(reader, "S3 object name: ")

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		_, err = client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   strings.NewReader(string(data)),
		})

		if err != nil {
			fmt.Println("Upload failed:", err)
			return
		}

		fmt.Println("Upload successful!")

	case "4":
		bucket := input(reader, "Bucket name: ")
		key := input(reader, "S3 object name: ")
		file := input(reader, "Save as: ")

		result, err := client.GetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

		if err != nil {
			fmt.Println("Download failed:", err)
			return
		}
		defer result.Body.Close()

		output, err := os.Create(file)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer output.Close()

		buffer := make([]byte, 32*1024)

		for {
			n, readErr := result.Body.Read(buffer)

			if n > 0 {
				_, writeErr := output.Write(buffer[:n])
				if writeErr != nil {
					fmt.Println("Write error:", writeErr)
					return
				}
			}

			if readErr != nil {
				break
			}
		}

		fmt.Println("Download successful!")

	case "5":
		bucket := input(reader, "Bucket name: ")
		key := input(reader, "S3 object name: ")

		_, err := client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

		if err != nil {
			fmt.Println("Delete failed:", err)
			return
		}

		fmt.Println("Object deleted successfully!")

	default:
		fmt.Println("Invalid option")
	}
}
