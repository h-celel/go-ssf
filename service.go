package go_ssf

import (
	"context"
	"sync"
)

type Service interface {
	ComponentSet
	Status(ctx context.Context) error
	Close()
}

func NewService() Service {
	return &defaultService{
		components: NewComponentSet(),
	}
}

type defaultService struct {
	m          sync.Mutex
	components ComponentSet
}

func (s *defaultService) Status(ctx context.Context) error {
	s.m.Lock()
	defer s.m.Unlock()
	for _, c := range s.components.GetAllComponents() {
		if err := c.Status(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *defaultService) Close() {
	s.m.Lock()
	defer s.m.Unlock()

	var wg sync.WaitGroup

	cs := s.components.GetAllComponents()
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c Component) {
			c.Close()
			wg.Done()
		}(c)
	}

	wg.Wait()
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
