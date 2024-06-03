# aws-cdk-assignment

# AWS ECS Deployment with Service and Workers

## Overview
In this challenge, you will create a simple infrastructure for deploying a service and two workers using AWS ECS (Elastic Container Service) and AWS CDK (Cloud Development Kit). The goal is to automate the provisioning of the necessary resources and deploy the service and workers as separate ECS services.

## Objective
Develop an automated infrastructure provisioning and deployment system that:
- Defines the necessary AWS resources using CDK in TypeScript or Python.
- Deploys a service and two workers as separate ECS services.
- Ensures the service and workers are running and accessible.

## Requirements
### Infrastructure Definition with CDK
1. **CDK Setup:** Set up a new CDK project in TypeScript or Python.
2. **VPC and Subnets:** Define a new VPC with public and private subnets using CDK constructs.
3. **ECS Cluster:** Create an ECS cluster to host the service and workers.
4. **Service Definition:** Define an ECS service for the main service component, specifying the task definition, desired count, and load balancer configuration.
5. **Worker Definitions:** Define two separate ECS services for the workers, specifying the task definitions and desired counts.

### Dockerization
1. **Dockerfiles:** Create Dockerfiles for the service and worker components, specifying the necessary dependencies and configurations.
2. **Docker Images:** Build Docker images for the service and workers and push them to Amazon Elastic Container Registry (ECR).

### Deployment
1. **CDK Deployment:** Use CDK to deploy the infrastructure and ECS services to your AWS account.
2. **Service Accessibility:** Ensure the service is accessible via a load balancer or public IP address.
3. **Worker Execution:** Verify that the workers are running and executing their intended tasks.

### Documentation
- Document the infrastructure architecture, deployment process, and any assumptions made.
- Provide instructions on how to build and deploy the solution.

## Evaluation Criteria
- **Infrastructure as Code:** The effectiveness of using CDK to define and manage the AWS infrastructure.
- **ECS Deployment:** The successful deployment of the service and workers as separate ECS services.
- **Service Accessibility:** The accessibility and functionality of the deployed service.
- **Worker Execution:** The proper execution and behavior of the deployed workers.

## Submission Guidelines
Submit your solution as a Git repository containing:
- **CDK Code:** The CDK project files defining the AWS infrastructure.
- **Dockerfiles:** The Dockerfiles for the service and worker components.
- **Documentation:** Detailed documentation of the architecture, deployment process, and any assumptions or considerations.

## Setup Instructions
1. Set up an AWS account and configure the necessary permissions and credentials.
2. Install and configure AWS CLI and AWS CDK on your local machine.
3. Create a new CDK project and define the infrastructure using CDK constructs.
4. Implement the Dockerfiles for the service and worker components.
5. Build the Docker images and push them to Amazon ECR.
6. Deploy the infrastructure and ECS services using CDK.
7. Verify the accessibility of the service and the execution of the workers.
8. Document the architecture, deployment steps, and any assumptions or considerations.

## Estimated Time
1. **Setting up the CDK project and defining the infrastructure:**
  - If you are familiar with AWS and CDK, this could take around 1-2 hours.
  - If you are new to CDK, it might take an additional 1-2 hours to learn the basics and set up the project.

2. **Implementing the Dockerfiles for the service and worker components:**
  - The time required for this task depends on the complexity of your service and worker components and your experience with Docker.
  - On average, creating Dockerfiles and building the images could take around 1-2 hours.

3. **Deploying the infrastructure and ECS services using CDK:**
  - Once you have the CDK project set up and the Dockerfiles ready, deploying the infrastructure and services should be relatively quick.
  - Deploying and verifying the deployment could take around 30 minutes to 1 hour.

4. **Testing and documenting the solution:**
  - Testing the accessibility of the service and the execution of the workers may take around 30 minutes to 1 hour.
  - Documenting the architecture, deployment process, and assumptions could take another 1-2 hours.

Considering these estimates, the total time required to complete this challenge could range from 4 to 8 hours, depending on your experience level and the complexity of your specific implementation.

**Note:** Provide instructions on how to set up and configure AWS credentials instead of directly accepting them.

Good luck, and we look forward to your solution!
