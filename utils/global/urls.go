package global

func (ram *roleAuthUrls) Get(key uint8) map[string]string {
	defer ram.m.RUnlock()
	ram.m.RLock()
	return ram.List[key]
}

func (ram *roleAuthUrls) Set(key uint8, val map[string]string)  {
	ram.m.Lock()
	defer ram.m.Unlock()
	ram.List[key] = val
}