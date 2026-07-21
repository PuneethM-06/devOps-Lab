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

