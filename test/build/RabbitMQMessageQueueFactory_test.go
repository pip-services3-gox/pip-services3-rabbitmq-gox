package test_build

import (
	"testing"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	build "github.com/pip-services3-gox/pip-services3-rabbitmq-gox/build"
	queues "github.com/pip-services3-gox/pip-services3-rabbitmq-gox/queues"
	"github.com/stretchr/testify/assert"
)

func TestRabbitMQMessageQueueFactory(t *testing.T) {
	factory := build.NewRabbitMQMessageQueueFactory()
	descriptor := cref.NewDescriptor("pip-services", "message-queue", "rabbitmq", "test", "1.0")

	canResult := factory.CanCreate(descriptor)
	assert.NotNil(t, canResult)

	comp, err := factory.Create(descriptor)
	assert.Nil(t, err)
	assert.NotNil(t, comp)

	queue := comp.(*queues.RabbitMQMessageQueue)
	assert.Equal(t, "test", queue.Name())
}
