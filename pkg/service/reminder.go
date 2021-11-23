package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type ReminderService struct {
	repo repository.Reminder
}

func NewReminderService(repo repository.Reminder) *ReminderService{
	return &ReminderService{repo: repo}
}

func (s *ReminderService) Create(reminder models.Reminder) (int, error) {
	return s.repo.Create(reminder)
}


func (s *ReminderService) GetAll() ([]models.Reminder, error) {
	return s.repo.GetAll()
}

func (s *ReminderService) GetById(reminderId int) (models.Reminder, error) {
	return s.repo.GetById(reminderId)
}

func (s *ReminderService) Delete(reminderId int) error {
	return s.repo.Delete(reminderId)
}

func (s *ReminderService) Update (reminderId int, input models.UpdateReminder) error {
	return s.repo.Update(reminderId,input)
}



