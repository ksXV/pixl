package worklist

type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	job := <-w.jobs
	return job
}

func New(bufsize int) Worklist {
	return Worklist{make(chan Entry, bufsize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

func (w *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
