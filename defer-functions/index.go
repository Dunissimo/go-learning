package main

import (
	"errors"
	"fmt"
)

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.FinishJob()

	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	job.Merged = make(map[string]string)

	for _, v := range job.Dicts {
		if v == nil {
			return job, errNilDict
		}

		for k2, v2 := range v {
			job.Merged[k2] = v2
		}
	}

	return job, nil
}

func (job *MergeDictsJob) FinishJob() {
	job.IsFinished = true
}

func main() {
	job, _ := ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b", "b": "c"},
		{"d": "e", "f": "g"},
		{"a": "z", "f": "g"},
	}})

	job2, err2 := ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
		nil,
	}})

	job3, err3 := ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
	}})

	fmt.Println(job)
	fmt.Println(job2, err2)
	fmt.Println(job3, err3)
}
