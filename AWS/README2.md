## AWS README PAGE 2

# DEPLOYMENT SCALE AND MANAGING INFRASTRUCTURE

## CLOUD FORMATION

- **It is a declarative way of outlining AWS infrastructure**
- Benefits are:
    - Infrastructure as code - No resources are created manually
    - Cost - Resource creation and termination can be automated
    - Generated diagram for our template
    - Leverage existing templates

## AWS CLOUD DEVELOPMENT KIT (CDK)
- Defining **cloud infrastructure using a familiar coding language**
- Code is **compiled into a cloudformation template**
![CDK](image-1.png)

## AWS ELASTIC BEANSTALK
- Generally we make use of 3 tier for Web application, that is, ELB -> EC2 -> DB
- Elastic beanstalk is a **developer centric view of deploying an application to AWS**
- **BEANSTALK - PaaS**
- Here **developer is managing the code while AWS is responsible for deployments and managing**
- **ELASTIC BEANSTALK HAS ITS OWN HEALTH MONITORING** 

## AWS CODE DEPLOY
- It is a **HYBRID SERVICE - ONPREM AND CLOUD**
- It is a way to **deploy application automatically**
- It helps in **UPGRADING APPLICATIONS FROM V1 TO V2 IN FEW CLICKS**

## AWS CODECOMMIT
- It is a **code repository, using the git technology**
- CodeCommit is a way of **storing code before moving it to deployment in AWS**
- Private, secure

## AWS CODEBUILD
- **SERVERLESS**
- Build code in the cloud
- **Compile source code, test and the output which is a package can be deployed**
- **PAY FOR TIME USED TO BUILD THE CODE**

## AWS CODEPIPELINE
- **We can connect CodeCommit and CodeBuild using a CodePipeline**
- It is basis for **CI/CD**
- Code -> Build -> Test -> Deploy

## OVERALL WORKFLOW

![alt text](image-2.png)

## AWS CODE ARTIFACT

- Storing and retrieving these dependencies is called **ARTIFACT MANAGEMENT**
- **Developers and CodeBuild can retrieve dependencies from CodeArtifact**

## AWS SSM - SYSTEM MANAGER
- Manage **Fleet of EC2 instances in On-prem and cloud**
- It is **HYBRID**
- We can do **AUTOMATED PATCHING**
- We need to install **SSM agent in EC2**

