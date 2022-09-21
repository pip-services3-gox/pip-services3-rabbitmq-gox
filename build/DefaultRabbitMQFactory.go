package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	queues "github.com/pip-services3-gox/pip-services3-rabbitmq-gox/queues"
)

// Creates RabbitMQMessageQueue components by their descriptors.
// See RabbitMQMessageQueue
type DefaultRabbitMQFactory struct {
	*cbuild.Factory
}

// NewDefaultRabbitMQFactory method are create a new instance of the factory.
func NewDefaultRabbitMQFactory() *DefaultRabbitMQFactory {
	c := DefaultRabbitMQFactory{}
	c.Factory = cbuild.NewFactory()

	rabbitMQMessageQueueFactoryDescriptor := cref.NewDescriptor("pip-services", "queue-factory", "rabbitmq", "*", "1.0")
	rabbitMQMessageQueueDescriptor := cref.NewDescriptor("pip-services", "message-queue", "rabbitmq", "*", "1.0")

	c.RegisterType(rabbitMQMessageQueueFactoryDescriptor, NewRabbitMQMessageQueueFactory)

	c.Register(rabbitMQMessageQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}

		return queues.NewEmptyRabbitMQMessageQueue(name)
	})

	return &c
}
