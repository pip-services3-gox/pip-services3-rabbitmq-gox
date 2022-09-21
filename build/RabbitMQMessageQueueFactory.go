package build

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	cqueues "github.com/pip-services3-gox/pip-services3-messaging-gox/queues"
	queues "github.com/pip-services3-gox/pip-services3-rabbitmq-gox/queues"
)

type RabbitMQMessageQueueFactory struct {
	*cbuild.Factory
	config     *cconf.ConfigParams
	references cref.IReferences
}

func NewRabbitMQMessageQueueFactory() *RabbitMQMessageQueueFactory {
	c := RabbitMQMessageQueueFactory{}
	c.Factory = cbuild.NewFactory()

	memoryQueueDescriptor := cref.NewDescriptor("pip-services", "message-queue", "rabbitmq", "*", "*")

	c.Register(memoryQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}

		return queues.NewEmptyRabbitMQMessageQueue(name)
	})
	return &c
}

func (c *RabbitMQMessageQueueFactory) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.config = config
}

func (c *RabbitMQMessageQueueFactory) SetReferences(ctx context.Context, references cref.IReferences) {
	c.references = references
}

// Creates a message queue component and assigns its name.
//
// Parameters:
//   - name: a name of the created message queue.
func (c *RabbitMQMessageQueueFactory) CreateQueue(name string) cqueues.IMessageQueue {
	queue := queues.NewEmptyRabbitMQMessageQueue(name)

	if c.config != nil {
		queue.Configure(context.Background(), c.config)
	}
	if c.references != nil {
		queue.SetReferences(context.Background(), c.references)
	}

	return queue
}
