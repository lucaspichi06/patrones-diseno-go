package mediator

import (
	"sync"

	"github.com/lucaspichi06/patrones-diseno-go/src/comportamiento/07-mediator/domain"
)

type StationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	traingQueue    []domain.Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *StationManager) CanLand(t domain.Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}

	s.traingQueue = append(s.traingQueue, t)
	return false
}

func (s *StationManager) NotifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.traingQueue) > 0 {
		firstTrainInQueue := s.traingQueue[0]
		s.traingQueue = s.traingQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
