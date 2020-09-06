package engine

type ConcurrentEngine struct {
	Scheduler		Scheduler
	WorkerCount		int
	ItemChan		chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady (chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i:=0; i<e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	for{
		result := <-out
		for _, item := range result.Items{
			//log.Printf("Got item %v", item)
			go func() {e.ItemChan <- item}()
		}

		for _, request := range result.Request{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier){
	go func() {
		for{
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}
