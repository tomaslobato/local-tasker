package utils

import "github.com/tomaslobato/lt/models"

type TaskSorter []models.Task

func (t TaskSorter) Len() int {
	return len(t)
}

func (t TaskSorter) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TaskSorter) Less(i, j int) bool {
	return t[i].Stamp < t[j].Stamp
}
