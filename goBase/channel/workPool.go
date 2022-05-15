package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Job struct {
    id int
    randomno int
}

type Result struct{
    job Job
    randomno int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(numbers int) int {
    sum := 0
    no := numbers
    for no != 0 {
        sum += no % 10
        no /= 10
    }
    return sum
}

func worker(wg *sync.WaitGroup) {
    for job := range jobs {
        output := Result{job, digits(job.randomno)}
        results <- output
    }
    wg.Done()
}

func creatWorkerPool(noOfWorkers int) {
    var wg sync.WaitGroup
    for i := 0; i < noOfWorkers; i++ {
        wg.Add(1)
        go worker(&wg)
    }
    wg.Wait()
    close(results)
}

func allocate(noOfJobs int) {
    for i := 0; i < noOfJobs; i++ {
        randomno := rand.Intn(999)
        job := Job{i, randomno}
        jobs <- job
    }
    close(jobs)
}

func result(done chan bool){
    for result := range results {
        //time.Sleep(1 * time.Second)// 模拟耗时场景
        fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.randomno)
    }
    done <- true
}

//文中的注释是为了测试go result(done)中的done是否可以去掉，
//测试结果：不可以去掉，当给result()中设置模拟耗时场景时可从输出结果看出并未输出或未
//输出完整的results,因为result()并发执行时没有done信道阻塞主协程main()，所以当main()//正常执行完毕退出时result()仍为完全打印出所有results
//在加入done后即使模拟result()耗时场景也会因为done阻塞主协程而强制等待result()完全输出

func main(){
    startTime := time.Now()
    jobs := 100 //修改100为10减少不必要耗时
    go allocate(jobs)
    done := make(chan bool)
    go result(done)
    workers := 10
    creatWorkerPool(workers) //可以不go此函数,即使go了，下一行代码还是处于阻塞
    <- done
    endTime := time.Now()
    diff := endTime.Sub(startTime)
    fmt.Println("total time taken ",diff.Seconds(), "seconds")
}
