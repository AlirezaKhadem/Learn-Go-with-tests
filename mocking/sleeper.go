package mocking

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	counter int
}

func (s *SpySleeper) Sleep() {
	s.counter += 1
}

func (s *SpySleeper) GetCounter() int {
	return s.counter
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
