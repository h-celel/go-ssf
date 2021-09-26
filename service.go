package go_ssf

import (
	"context"
	"sync"
)

type Service interface {
	ComponentSet
	Status(ctx context.Context) error
	Shutdown()
	SetShutdownCallback(func())
	Context() context.Context
}

type defaultService struct {
	ctx          context.Context
	cancel       context.CancelFunc
	m            sync.Mutex
	components   ComponentSet
	shutdownFunc func()
}

func NewService(ctx context.Context) Service {
	ctx, cancel := context.WithCancel(ctx)

	s := &defaultService{
		ctx:        ctx,
		cancel:     cancel,
		components: newComponentSet(),
	}

	go func(ctx context.Context, s *defaultService) {
		<-ctx.Done()
		s.onShutdown()
	}(ctx, s)

	return s
}

func (s *defaultService) Status(ctx context.Context) error {
	s.m.Lock()
	defer s.m.Unlock()

	if err := s.ctx.Err(); err != nil {
		return err
	}

	for _, c := range s.components.GetAllComponents() {
		if err := c.Status(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *defaultService) Shutdown() {
	s.m.Lock()
	defer s.m.Unlock()
	s.cancel()
}

func (s *defaultService) Context() context.Context {
	return s.ctx
}

func (s *defaultService) SetShutdownCallback(f func()) {
	s.m.Lock()
	defer s.m.Unlock()
	s.shutdownFunc = f
}

func (s *defaultService) GetAllComponents() []Component {
	s.m.Lock()
	defer s.m.Unlock()
	return s.components.GetAllComponents()
}

func (s *defaultService) GetComponentsByType(componentType ComponentType) []Component {
	s.m.Lock()
	defer s.m.Unlock()
	return s.components.GetComponentsByType(componentType)
}

func (s *defaultService) GetComponent(componentType ComponentType, i uint32) Component {
	s.m.Lock()
	defer s.m.Unlock()
	return s.components.GetComponent(componentType, i)
}

func (s *defaultService) AddComponent(componentType ComponentType, component Component) {
	s.m.Lock()
	defer s.m.Unlock()
	s.components.AddComponent(componentType, component)
}

func (s *defaultService) onShutdown() {
	if f := s.shutdownFunc; f != nil {
		f()
	}
}
