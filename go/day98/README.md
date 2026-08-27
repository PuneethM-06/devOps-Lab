# DAY 98 - AWS SDK GO, IAM, S3, STS & CLOUWATCH

1. ### What is AWS SDK for Go?
- AWS SDK for Go is a set of Go libraries that lets your Go application interact with AWS services using Go native APIs instead of manually making an http request.

2. ### AWS SDK for G0 V2 
- There are 3 thing to understand:
    1.AWS SDK config
    2. Region 
    3. Credentials 

1. **AWS SDK Configuration**
- We typically load the AWS config using SDK's config package
```
Go Application
      │
      ▼
Load AWS Configuration
      │
      ├── Credentials
      ├── Region
      └── Other settings

```
2. **Region**
- AWS resources are associated with the region so we mention them here 

3. **Credentials**
- AWS credentials to authenticate the application 
- so in production we ideally want 
```
Go Application
      │
      ▼
AWS SDK
      │
      ▼
IAM Role
      │
      ▼
Temporary Credentials
```

3. ### How AWS Authentication Works in Go
-  This is where **Credential provider chain** comes in
- **Application asks SDK for credentials, the SDK determines where to obtain them**

4. ### IAM Roles and Credentials Provider chain 
