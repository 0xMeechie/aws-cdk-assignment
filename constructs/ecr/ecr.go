package ecr

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/jsii-runtime-go"
)

func NewRepo(stack awscdk.Stack) awsecr.Repository {
	repoProps := awsecr.RepositoryProps{
		RepositoryName: jsii.String(strings.ToLower(*stack.StackName()) + "-repo"),
		EmptyOnDelete:  jsii.Bool(true),
		RemovalPolicy:  awscdk.RemovalPolicy_DESTROY,
	}
	repo := awsecr.NewRepository(stack, jsii.String("Repo"), &repoProps)

	return repo
}

func WorkerImage(stack awscdk.Stack) awsecrassets.DockerImageAsset {

	wokerProps := &awsecrassets.DockerImageAssetProps{
		Directory: jsii.String("workers"),
		Platform:  awsecrassets.Platform_LINUX_AMD64(),
	}
	image := awsecrassets.NewDockerImageAsset(stack, jsii.String("worker"), wokerProps)

	return image

}

func ServiceImage(stack awscdk.Stack) awsecrassets.DockerImageAsset {

	serviceProps := &awsecrassets.DockerImageAssetProps{
		Directory: jsii.String("services"),
		Platform:  awsecrassets.Platform_LINUX_AMD64(),
	}
	image := awsecrassets.NewDockerImageAsset(stack, jsii.String("service"), serviceProps)
	return image

}
