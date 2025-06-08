package main

import "sync"

type SurveyStore interface {
	GetDepartmentData(waitGroup *sync.WaitGroup, target *map[string][]float32, dept string, errChan chan<- error)
}
