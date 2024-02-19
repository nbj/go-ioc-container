package Nbj

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_implementations_can_be_registered_and_resolved(t *testing.T) {
	// Arrange
	InitializeIocContainer()
	someImplementation := "This string represent some kind of implementation"

	// Act
	IocContainer.Register("some_implementation", someImplementation)

	// Assert
	assert.Equal(t, "This string represent some kind of implementation", IocContainer.Resolve("some_implementation"))
}

func Test_implementations_can_be_forgotten(t *testing.T) {
	// Arrange
	InitializeIocContainer()
	someImplementation := "This string represent some kind of implementation"

	// Act
	IocContainer.Register("some_implementation", someImplementation)
	assert.Equal(t, "This string represent some kind of implementation", IocContainer.Resolve("some_implementation"))
	IocContainer.Forget("some_implementation")

	// Assert
	assert.Equal(t, nil, IocContainer.Resolve("some_implementation"))
}
