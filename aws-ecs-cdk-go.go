package main

import (
	"aws-ecs-cdk-go/constructs/ecr"
	"aws-ecs-cdk-go/constructs/ecs"
	"aws-ecs-cdk-go/constructs/vpc"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AwsEcsCdkGoStackProps struct {
	awscdk.StackProps
}

func NewAwsEcsCdkGoStack(scope constructs.Construct, id string, props *AwsEcsCdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	stackinfo := ecs.StackInfo{
		Stack:   stack,
		VPC:     vpc.NewVPC(stack),
		ECR:     ecr.NewRepo(stack),
		Worker:  ecr.WorkerImage(stack),
		Service: ecr.ServiceImage(stack),
	}

	ecs.NewECS(stackinfo)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewAwsEcsCdkGoStack(app, "AwsEcsCdkGoStack", &AwsEcsCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
