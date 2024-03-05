package entity

type City struct {
	City string
}

func (c *City) Exist() bool {
	return c.City != ""
}
