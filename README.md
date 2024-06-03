# ECS CDK Assignment Document

Before deploying there is a few prerequisites.

 - Node.js 14.15.0 or higher must be install
 - AWS CDK cli must be install
 - Go Must be install
 - Docker must be installed and running.  
 - AWS creds should be configured and active.
 - CDK_DEFAULT_REGION and CDK_DEFAULT_ACCOUNT should be set.

# Architecture
This architecure features one vpc. It has three subnet groups(public,data and service). The Public Subnet is internet facing, data and service are private subnets. These three subnet groups are spread across three different AZs. Within each AZ there is a nat gateway for HA. The ecs cluster is placed within the service subnet. In the public subnet lives a ALB that routes are traffic to the ecs cluster.

The ECS clusters holds two services. One for services and another for the workers.



## Installing Node.js 
Follow the directions here depending on your OS and preference.  
https://nodejs.org/en/download/package-manager


## Installing AWS CDK 

```
npm install -g aws-cdk
```

Run the following command to verify a successful installation. The AWS CDK CLI should output the version number:

```
cdk --version
```

## Install Go

Click the link and install go depending on your OS/Machine
https://go.dev/doc/install

## Start Docker
Ensure that Docker is downloaded and running. This can be accomplished by opening the docker desktop application or running the following command. (linux/MacOS)
```
sudo systemctl start docker
```
## Configured AWS Credentials 

This can be done two ways. Getting IAM access credentials or via SSO. You want to get an access key and secret for your iam user via the aws console. 

Or Use ```aws sso login``` from the command line to get credentials. 

Once you get your keys and place them in the /.aws/credential file test your creds with 

```aws sts get-caller-identity```

AWS Documents it pretty well.
https://docs.aws.amazon.com/workspaces-web/latest/adminguide/getting-started-iam-user-access-keys.html



# Set Env Vars 
This is optional but if this is not set than the subnets will only scale across two AZs so it is highly recommended to configure this. 

This can be done with the following commands 
```
export CDK_DEFAULT_REGION = us-east-1
export CDK_DEFAULT_ACCOUNT = 1234567890
```

Replace the region and account number with your account number and where you want to deploy

## Deploying

To deploy run ```cdk deploy``` from the root of the repo. This will build the docker images to your local PC. From there once the images built it will upload them to ecr and proceed to build out the rest of the environment

## Destroying

Run ```cdk destroy``` from the root of the repo. This will proceed to tear down the stack.

## Testing and Verifying

You can log into your aws account and check out all the newly built infrastructure. When the stack is finished you will be able to take the dns URL from the load balancers output and see a message from the work and service task. 
