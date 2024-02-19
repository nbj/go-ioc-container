package Nbj

// Container
// The structure of the (Inversion of Control) Container
type Container struct {
	implementations map[string]any
}

// IocContainer
// The global instance of the IocContainer
var IocContainer *Container

// InitializeIocContainer
// Initializes the IocContainer. This must be called before the
// container can have any implementations registered to it
func InitializeIocContainer() {
	IocContainer = new(Container)
	IocContainer.implementations = make(map[string]any)
}

// Register
// Registers an implementation to a specific key
func (container *Container) Register(key string, implementation any) *Container {
	if _, ok := container.implementations[key]; !ok {
		container.implementations[key] = implementation
	}

	return container
}

// Resolve
// Resolves a specific implementation using its key
func (container *Container) Resolve(key string) any {
	if implementation, ok := container.implementations[key]; ok {
		return implementation
	}

	return nil
}

// Forget
// Forces the IocContainer to forget an implementation at a specific key
func (container *Container) Forget(key string) *Container {
	delete(container.implementations, key)

	return container
}
