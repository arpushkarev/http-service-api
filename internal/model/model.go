package model

type Info struct {
	Id map[string]int64
}

func (i *Info) Add(key string, value int64) Info {
	var myMap = make(map[string]int64)
	myMap[key] = value
	i.Id = myMap
	return *i
}
