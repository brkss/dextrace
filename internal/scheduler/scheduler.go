package scheduler

import (
	"log"
	"time"

	"github.com/brkss/dextrace/internal/domain"
	"github.com/brkss/dextrace/internal/usecase"
)

type Scheduler struct {
	sibionicUseCase *usecase.SibionicUseCase
	nightscoutUseCase *usecase.NightscoutUsecase
	userID string
	user domain.User
	stopChan chan bool
}

func NewScheduler(sibionicUseCase *usecase.SibionicUseCase, nightscoutUseCase *usecase.NightscoutUsecase, userID string, user domain.User) *Scheduler {
	return &Scheduler{
		sibionicUseCase:   sibionicUseCase,
		nightscoutUseCase: nightscoutUseCase,
		userID:            userID,
		user:              user,
		stopChan:          make(chan bool),
	}
}

func (s *Scheduler) Start() {
	log.Println("Starting scheduler for push-to-nightscout (every 5 minutes)")
	
	// Run immediately on start
	s.pushToNightscout()
	
	// Schedule to run every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			s.pushToNightscout()
		case <-s.stopChan:
			log.Println("Scheduler stopped")
			return
		}
	}
}

func (s *Scheduler) Stop() {
	s.stopChan <- true
}

func (s *Scheduler) pushToNightscout() {
	log.Println("Executing scheduled push-to-nightscout")
	
	if s.userID == "" {
		log.Println("Error: user ID is required")
		return
	}
	
	data, err := s.sibionicUseCase.GetGlucoseData(s.user, s.userID)
	if err != nil {
		log.Printf("Error getting glucose data: %v", err)
		return
	}
	
	err = s.nightscoutUseCase.PushData(*data)
	if err != nil {
		log.Printf("Error pushing data to Nightscout: %v", err)
		return
	}
	
	log.Println("Successfully pushed data to Nightscout")
} 