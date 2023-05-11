package model

import "strconv"

type Id struct {
	Id int64
}

func (i *Id) GetId(id int64) string {
	idStr := strconv.Itoa(int(id + 1))
	return idStr
}
