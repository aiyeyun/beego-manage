package global

func (ram *roleAuthMenus) Get(key uint8) interface{} {
	defer ram.m.RUnlock()
	ram.m.RLock()
	return ram.List[key]
}

func (ram *roleAuthMenus) Set(key uint8, val interface{})  {
	ram.m.Lock()
	defer ram.m.Unlock()
	ram.List[key] = val
}