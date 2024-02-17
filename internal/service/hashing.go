package service

import "strconv"

type ConsistentHashing struct {
}

func (h *ConsistentHashing) Hash(key string) (int, error) {
	return strconv.Atoi(key)
}
