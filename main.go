package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Channel struct {
	pulumi.ResourceState
	Id pulumi.IDOutput `pulumi:"channelId"`
}

func NewChannel(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*Channel, error) {
	channel := &Channel{}
	err := ctx.RegisterComponentResource("pulumi:gcp:Channel", name, channel, opts...)
	if err != nil {
		return nil, err
	}

	topic, err := pubsub.NewTopic(ctx, name, &pubsub.TopicArgs{
		Name: pulumi.String(name),
	}, pulumi.Parent(channel))
	if err != nil {
		return nil, err
	}

	_, err = pubsub.NewTopic(ctx, name+"-dlq", &pubsub.TopicArgs{
		Name: pulumi.String(name + "-dlq"),
	}, pulumi.Parent(channel))
	if err != nil {
		return nil, err
	}

	ctx.RegisterResourceOutputs(channel, pulumi.Map{
		"Id": topic.ID(),
	})

	return channel, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
