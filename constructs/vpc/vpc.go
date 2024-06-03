package vpc

import (
	"aws-ecs-cdk-go/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func NewVPC(stack awscdk.Stack) awsec2.Vpc {
	subnets := []*awsec2.SubnetConfiguration{
		{
			Name:                jsii.String("Public"),
			SubnetType:          awsec2.SubnetType_PUBLIC,
			CidrMask:            jsii.Number(config.CidrMask),
			MapPublicIpOnLaunch: jsii.Bool(true),
		},
		{
			Name:       jsii.String("Service"),
			SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
			CidrMask:   jsii.Number(config.CidrMask),
		},
		{
			Name:       jsii.String("Data"),
			SubnetType: awsec2.SubnetType_PRIVATE_ISOLATED,
			CidrMask:   jsii.Number(config.CidrMask),
		},
	}

	vpcProp := awsec2.VpcProps{
		CreateInternetGateway: jsii.Bool(true),
		EnableDnsHostnames:    jsii.Bool(true),
		EnableDnsSupport:      jsii.Bool(true),
		IpAddresses:           awsec2.IpAddresses_Cidr(jsii.String(config.VPCCidr)),
		VpcName:               jsii.String(*stack.StackName() + "-VPC"),
		SubnetConfiguration:   &subnets,
	}

	//

	vpc := awsec2.NewVpc(stack, jsii.String("VPC"), &vpcProp)

	for _, subnet := range *vpc.PublicSubnets() {
		subnetName := jsii.String(*stack.StackName() + "-Public-" + *subnet.AvailabilityZone())
		awscdk.Tags_Of(subnet).Add(jsii.String("Name"), jsii.String(*subnetName), &awscdk.TagProps{})
	}

	for _, subnet := range *vpc.SelectSubnetObjects(&awsec2.SubnetSelection{SubnetGroupName: jsii.String("Service")}) {

		subnetName := jsii.String(*stack.StackName() + "-Service-" + *subnet.AvailabilityZone())
		awscdk.Tags_Of(subnet).Add(jsii.String("Name"), jsii.String(*subnetName), &awscdk.TagProps{})
	}

	for _, subnet := range *vpc.SelectSubnetObjects(&awsec2.SubnetSelection{SubnetGroupName: jsii.String("Data")}) {

		subnetName := jsii.String(*stack.StackName() + "-Data-" + *subnet.AvailabilityZone())
		awscdk.Tags_Of(subnet).Add(jsii.String("Name"), jsii.String(*subnetName), &awscdk.TagProps{})
	}

	return vpc

}
