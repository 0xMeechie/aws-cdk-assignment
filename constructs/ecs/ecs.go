package ecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/jsii-runtime-go"
)

type StackInfo struct {
	Stack   awscdk.Stack
	VPC     awsec2.Vpc
	ECR     awsecr.Repository
	Worker  awsecrassets.DockerImageAsset
	Service awsecrassets.DockerImageAsset
}

func NewECS(stack StackInfo) awsecs.Cluster {
	clusterProps := awsecs.ClusterProps{
		ClusterName: jsii.String(*stack.Stack.StackName() + "cluster"),
		Vpc:         stack.VPC,
		DefaultCloudMapNamespace: &awsecs.CloudMapNamespaceOptions{
			Name: jsii.String("default")},
	}
	cluster := awsecs.NewCluster(stack.Stack, jsii.String("cluster"), &clusterProps)
	sg := ecsSG(stack)

	worker := fargateWorkers(stack, cluster, sg)
	service := fargateServices(stack, cluster, sg)

	newWorkerLoadBalancer(stack, worker)
	newServiceLoadBalancer(stack, service)
	return cluster
}

func fargateWorkers(stack StackInfo, cluster awsecs.Cluster, sg awsec2.SecurityGroup) awsecs.FargateService {

	taskProps := awsecs.FargateTaskDefinitionProps{
		Cpu:            jsii.Number(512),
		MemoryLimitMiB: jsii.Number(1024),
	}

	task := awsecs.NewFargateTaskDefinition(stack.Stack, jsii.String("Worker_task"), &taskProps)

	containerProp := awsecs.ContainerDefinitionProps{
		Image:          awsecs.EcrImage_FromDockerImageAsset(stack.Worker),
		ContainerName:  jsii.String("worker-container"),
		TaskDefinition: task,
	}
	container := awsecs.NewContainerDefinition(stack.Stack, jsii.String("Worker_container"), &containerProp)
	container.AddPortMappings(&awsecs.PortMapping{
		ContainerPort: jsii.Number(80),
		HostPort:      jsii.Number(80),
		Protocol:      awsecs.Protocol_TCP,
	})

	mapOptions := &awsecs.CloudMapOptions{
		Name:              jsii.String("worker-container"),
		Container:         container,
		CloudMapNamespace: cluster.DefaultCloudMapNamespace(),
		ContainerPort:     jsii.Number(80),
	}

	serviceProps := awsecs.FargateServiceProps{
		TaskDefinition:  task,
		Cluster:         cluster,
		VpcSubnets:      &awsec2.SubnetSelection{SubnetGroupName: jsii.String("Service")},
		CloudMapOptions: mapOptions,
		SecurityGroups:  &[]awsec2.ISecurityGroup{sg},
		ServiceName:     jsii.String(*stack.Stack.StackName() + "-worker-service"),
	}

	service := awsecs.NewFargateService(stack.Stack, jsii.String("workers_fargate"), &serviceProps)

	return service
}

func fargateServices(stack StackInfo, cluster awsecs.Cluster, sg awsec2.SecurityGroup) awsecs.FargateService {

	taskProps := awsecs.FargateTaskDefinitionProps{
		Cpu:            jsii.Number(512),
		MemoryLimitMiB: jsii.Number(1024),
	}

	task := awsecs.NewFargateTaskDefinition(stack.Stack, jsii.String("Service_task"), &taskProps)

	containerProp := awsecs.ContainerDefinitionProps{
		Image:          awsecs.EcrImage_FromDockerImageAsset(stack.Service),
		ContainerName:  jsii.String("service-container"),
		TaskDefinition: task,
	}
	container := awsecs.NewContainerDefinition(stack.Stack, jsii.String("Service_container"), &containerProp)
	container.AddPortMappings(&awsecs.PortMapping{
		ContainerPort: jsii.Number(80),
		HostPort:      jsii.Number(80),
		Protocol:      awsecs.Protocol_TCP,
	})

	mapOptions := &awsecs.CloudMapOptions{
		Name:              jsii.String("services-container"),
		Container:         container,
		CloudMapNamespace: cluster.DefaultCloudMapNamespace(),
		ContainerPort:     jsii.Number(80),
	}

	serviceProps := awsecs.FargateServiceProps{
		TaskDefinition:  task,
		Cluster:         cluster,
		VpcSubnets:      &awsec2.SubnetSelection{SubnetGroupName: jsii.String("Service")},
		CloudMapOptions: mapOptions,
		SecurityGroups:  &[]awsec2.ISecurityGroup{sg},
		ServiceName:     jsii.String(*stack.Stack.StackName() + "-service"),
	}

	service := awsecs.NewFargateService(stack.Stack, jsii.String("service_fargate"), &serviceProps)

	return service
}

func ecsSG(stack StackInfo) awsec2.SecurityGroup {

	sgProps := awsec2.SecurityGroupProps{
		Vpc:               stack.VPC,
		AllowAllOutbound:  jsii.Bool(true),
		Description:       jsii.String("This allows connection to the ecs service from the load balancer"),
		SecurityGroupName: jsii.String("ECS Inbound SG"),
	}
	SG := awsec2.NewSecurityGroup(stack.Stack, jsii.String("ECS_sg"), &sgProps)

	SG.AddIngressRule(awsec2.Peer_Ipv4(jsii.String("0.0.0.0/0")), awsec2.Port_Tcp(jsii.Number(80)), nil, nil)
	return SG
}

func newWorkerLoadBalancer(stack StackInfo, service awsecs.FargateService) {

	lb := awselasticloadbalancingv2.NewApplicationLoadBalancer(stack.Stack, jsii.String("worker_LB"), &awselasticloadbalancingv2.ApplicationLoadBalancerProps{
		Vpc:            stack.VPC,
		InternetFacing: jsii.Bool(true),
	})
	listener := lb.AddListener(jsii.String("worker_Listener"), &awselasticloadbalancingv2.BaseApplicationListenerProps{
		Port: jsii.Number(80),
	})

	targets := &[]awselasticloadbalancingv2.IApplicationLoadBalancerTarget{
		service.LoadBalancerTarget(&awsecs.LoadBalancerTargetOptions{
			ContainerName: jsii.String("worker-container"),
			ContainerPort: jsii.Number(80),
		})}

	listener.AddTargets(jsii.String("ECS1"), &awselasticloadbalancingv2.AddApplicationTargetsProps{
		Port:    jsii.Number(80),
		Targets: targets,
	})

	awscdk.NewCfnOutput(stack.Stack, jsii.String("Worker_lb_output"), &awscdk.CfnOutputProps{

		Value: jsii.String(*lb.LoadBalancerDnsName())})
}

func newServiceLoadBalancer(stack StackInfo, service awsecs.FargateService) {

	lb := awselasticloadbalancingv2.NewApplicationLoadBalancer(stack.Stack, jsii.String("service_LB"), &awselasticloadbalancingv2.ApplicationLoadBalancerProps{
		Vpc:            stack.VPC,
		InternetFacing: jsii.Bool(true),
	})
	listener := lb.AddListener(jsii.String("service_Listener"), &awselasticloadbalancingv2.BaseApplicationListenerProps{
		Port: jsii.Number(80),
	})

	targets := &[]awselasticloadbalancingv2.IApplicationLoadBalancerTarget{
		service.LoadBalancerTarget(&awsecs.LoadBalancerTargetOptions{
			ContainerName: jsii.String("service-container"),
			ContainerPort: jsii.Number(80),
		})}

	listener.AddTargets(jsii.String("ECS1"), &awselasticloadbalancingv2.AddApplicationTargetsProps{
		Port:    jsii.Number(80),
		Targets: targets,
	})

	awscdk.NewCfnOutput(stack.Stack, jsii.String("service_lb_output"), &awscdk.CfnOutputProps{

		Value: jsii.String(*lb.LoadBalancerDnsName())})
}
