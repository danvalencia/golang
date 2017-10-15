package main

import (
	"fmt"
	"time"
)

// JobRunnable represents a generic Job
type JobRunnable interface {
	ID() string
	Run() *JobResult
}

// JobResult contains an id and the result of the Job
type JobResult struct {
	jobID  string
	result interface{}
}

// Job represents a Job to be run
type Job struct {
	jobID         string
	jobExecutable func() *JobResult
}

// JobRunner is a struct holding a jobQueue in the form of a channel
type JobRunner struct {
	jobQueue   chan JobRunnable
	jobResults chan *JobResult
	quit       chan int
}

// NewJob creates a new *Job struct with the given jobID and jobExecutable
func NewJob(jobID string, jobExecutable func() *JobResult) *Job {
	return &Job{
		jobID:         jobID,
		jobExecutable: jobExecutable,
	}
}

// SubmitJob submits a job to be run sequentially
func (jobRunner *JobRunner) SubmitJob(job JobRunnable) {
	fmt.Printf("Sumbitting job %v\n", job)
	jobRunner.jobQueue <- job
}

// Start with start up the JobRunner making it ready to accept jobs
func (jobRunner *JobRunner) Start() {
	go func() {
		for {
			select {
			case job := <-jobRunner.jobQueue:
				fmt.Printf("Job is %v\n", job)
				res := job.Run()
				jobRunner.jobResults <- res
				fmt.Printf("After running job %v\n", job)
			case <-jobRunner.quit:
				fmt.Println("Quitting job runner!")
				return
			}
		}
	}()
}

// Run will represent the task at hand
func (job *Job) Run() *JobResult {
	return job.jobExecutable()
}

// ID returns the id for the given job
func (job *Job) ID() string {
	return job.jobID
}

func main() {
	jobRunner := &JobRunner{
		jobQueue:   make(chan JobRunnable, 10),
		jobResults: make(chan *JobResult, 10),
		quit:       make(chan int),
	}

	jobRunner.Start()

	jobs := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, jobID := range jobs {
		job := NewJob(jobID, func(jId string) func() *JobResult {
			return func() *JobResult {
				time.Sleep(time.Duration(100) * time.Millisecond)
				fmt.Printf("Doing something super important...%v\n", jId)
				return &JobResult{
					jobID:  jId,
					result: fmt.Sprintf("Result of Job %v", jId),
				}
			}
		}(jobID))
		jobRunner.SubmitJob(job)
	}

	results := make([]JobResult, len(jobs))

	for i := range jobs {
		result := <-jobRunner.jobResults
		results[i] = *result
	}

	fmt.Printf("Results: %v\n", results)
	jobRunner.quit <- 0
	time.Sleep(time.Duration(1500))
}
