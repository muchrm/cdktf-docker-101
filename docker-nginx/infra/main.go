package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v3/container"
	"github.com/cdktf/cdktf-provider-docker-go/docker/v3/image"
	"github.com/cdktf/cdktf-provider-docker-go/docker/v3/network"
	dockerprovider "github.com/cdktf/cdktf-provider-docker-go/docker/v3/provider"
)

func NewNginxImage(stack cdktf.TerraformStack) image.Image {
	return image.NewImage(stack, jsii.String("nginxImage"), &image.ImageConfig{
		Name:        jsii.String("nginx:latest"),
		KeepLocally: jsii.Bool(false),
	})
}

func NewInternalNetwork(stack cdktf.TerraformStack) network.Network {

	network := network.NewNetwork(stack, jsii.String("internalNetwork"), &network.NetworkConfig{
		Name: jsii.String("frontend-internal"),
	})

	return network
}

func NewWebStack(stack cdktf.TerraformStack, image image.Image, network network.Network) cdktf.TerraformStack {

	container.NewContainer(stack, jsii.String("webContainer"), &container.ContainerConfig{
		Image:    image.Latest(),
		Name:     jsii.String("web"),
		Networks: jsii.Strings(*network.Name()),
	})

	return stack
}

func NewProxyStack(stack cdktf.TerraformStack, image image.Image, network network.Network) cdktf.TerraformStack {

	container.NewContainer(stack, jsii.String("proxyContainer"), &container.ContainerConfig{
		Image:    image.Latest(),
		Name:     jsii.String("proxy"),
		Networks: jsii.Strings(*network.Name()),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(80), External: jsii.Number(8000),
		}},
	})

	return stack
}

func NewFrontendStack(scope constructs.Construct) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, jsii.String("web-infra"))

	dockerprovider.NewDockerProvider(stack, jsii.String("docker"), &dockerprovider.DockerProviderConfig{})

	network := NewInternalNetwork(stack)
	nginxImage := NewNginxImage(stack)
	NewProxyStack(stack, nginxImage, network)
	NewWebStack(stack, nginxImage, network)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewFrontendStack(app)

	app.Synth()
}
