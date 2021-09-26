package go_ssf

type ComponentSet interface {
	GetAllComponents() []Component
	GetComponentsByType(componentType ComponentType) []Component
	GetComponent(componentType ComponentType, i uint32) Component
	AddComponent(componentType ComponentType, component Component)
}

type defaultComponentSet map[ComponentType][]Component

func newComponentSet() ComponentSet {
	return &defaultComponentSet{}
}

func (d *defaultComponentSet) GetAllComponents() (res []Component) {
	if d == nil {
		return nil
	}
	for _, v := range *d {
		for _, c := range v {
			res = append(res, c)
		}
	}
	return
}

func (d *defaultComponentSet) GetComponentsByType(componentType ComponentType) (res []Component) {
	if d == nil {
		return nil
	}
	for _, c := range (*d)[componentType] {
		res = append(res, c)
	}
	return
}

func (d *defaultComponentSet) GetComponent(componentType ComponentType, i uint32) Component {
	if d == nil {
		return nil
	}
	if cs, ok := (*d)[componentType]; ok {
		if len(cs) >= int(i+1) {
			return cs[i]
		}
	}

	return nil
}

func (d *defaultComponentSet) AddComponent(componentType ComponentType, component Component) {
	if d == nil {
		return
	}
	(*d)[componentType] = append((*d)[componentType], component)
}
