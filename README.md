# Go S3 Automation

A simple and lightweight **AWS S3 automation CLI built with Go** using the official AWS SDK for Go v2.

The tool allows you to connect to AWS using credentials and perform common S3 operations directly from the terminal.

## Features

* AWS credential-based authentication
* Configurable AWS Region
* List S3 buckets
* List objects inside a bucket
* Upload files to S3
* Download files from S3
* Delete objects from S3
* Simple interactive CLI
* Built with Go and AWS SDK for Go v2

## Tech Stack

* Go
* AWS SDK for Go v2
* Amazon S3
* AWS IAM

## Project Structure

```text
go-s3-automation/
│
├── main.go
├── go.mod
└── README.md
```

## Requirements

Before running the project, make sure you have:

* Go 1.20+
* AWS Account
* AWS IAM user with appropriate S3 permissions

Check Go installation:

```bash
go version
```

## Installation

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/go-s3-automation.git
```

Go inside the project:

```bash
cd go-s3-automation
```

Install dependencies:

```bash
go mod tidy
```

## Run

Start the application:

```bash
go run main.go
```

The application will ask for:

```text
AWS Access Key:
AWS Secret Key:
AWS Region [us-east-1]:
```

Example:

```text
AWS Access Key: AKIAxxxxxxxxxxxx
AWS Secret Key: xxxxxxxxxxxxxxxxxxxxxxxxx
AWS Region [us-east-1]: ap-south-1
```

## Available Operations

After authentication, the CLI provides:

```text
1. List Buckets
2. List Objects
3. Upload File
4. Download File
5. Delete Object
```

### 1. List Buckets

Displays all S3 buckets accessible by the provided AWS credentials.

```text
S3 Buckets:
- my-project-bucket
- devops-backup
- application-data
```

### 2. List Objects

Enter the bucket name to view its objects.

```text
Bucket name: my-project-bucket

Objects:
- app.zip
- config.json
- backup.tar.gz
```

### 3. Upload File

Upload a local file to an S3 bucket.

Example:

```text
Bucket name: my-project-bucket
Local file path: ./app.zip
S3 object name: releases/app.zip
```

The file will be uploaded as:

```text
s3://my-project-bucket/releases/app.zip
```

### 4. Download File

Download an S3 object to your local machine.

Example:

```text
Bucket name: my-project-bucket
S3 object name: releases/app.zip
Save as: ./app.zip
```

### 5. Delete Object

Delete an object from an S3 bucket.

Example:

```text
Bucket name: my-project-bucket
S3 object name: releases/app.zip
```

## IAM Permissions

For basic functionality, the IAM user needs appropriate S3 permissions.

Example policy:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets",
        "s3:ListBucket",
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject"
      ],
      "Resource": "*"
    }
  ]
}
```

For production environments, avoid using `"Resource": "*"` and restrict permissions to the required buckets and objects.

## Security

**Never commit AWS credentials to GitHub.**

Do not put credentials inside:

```text
main.go
README.md
.env
config files
```

Add sensitive files to `.gitignore`:

```gitignore
.env
*.pem
*.key
credentials
```

For production workloads, prefer:

* IAM Roles
* AWS IAM Identity Center
* Environment variables
* AWS credential profiles
* Temporary credentials

## Build

Create a production binary:

```bash
go build -o s3-automation
```

Run it:

```bash
./s3-automation
```

For Windows:

```bash
go build -o s3-automation.exe
```

## Example Workflow

```text
User
 │
 │ AWS Credentials
 ▼
Go S3 Automation CLI
 │
 │ AWS SDK for Go v2
 ▼
AWS S3
 │
 ├── List Buckets
 ├── List Objects
 ├── Upload
 ├── Download
 └── Delete
```

## Use Cases

This project can be useful for:

* DevOps automation
* Cloud engineering practice
* S3 file management
* CI/CD automation
* Backup automation
* AWS scripting
* Learning Go for DevOps
* Building custom AWS CLI tools

## Future Improvements

Planned improvements:

* Environment variable authentication
* AWS credential profiles
* Bucket creation
* Bucket deletion
* Recursive directory upload
* Recursive S3 download
* Multipart uploads
* Progress bars
* Automatic backups
* Logging
* Configuration file support
* Docker support
* CI/CD with GitHub Actions

## Disclaimer

This project is intended for learning, automation, and development purposes.

Always follow AWS security best practices and use the minimum IAM permissions required for your workload.

## License

MIT License
