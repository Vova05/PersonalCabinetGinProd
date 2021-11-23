package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
	"time"
)

type ReminderDB struct {
	db *gorm.DB
}


func NewReminderDB(db *gorm.DB) *ReminderDB{
	return &ReminderDB{db: db}
}

func (r *ReminderDB) Create(reminder models.Reminder) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	row :=	r.db.Table("reminders").Create(&reminder)

	if err :=  row.Table("reminders").Where("hint = ?",reminder.Hint).Find(&reminder).Error; err != nil{
		return 0,err
	}
	r.db.Table("reminders").Last(&reminder)
	id = reminder.IdReminder
	return id, nil
}

func (r *ReminderDB) GetAll() ([]models.Reminder, error) {
	var reminder []models.Reminder
	err := r.db.Table("reminders").Find(&reminder).Error
	return reminder,  err
}

func (r *ReminderDB) GetById(reminderId int) (models.Reminder, error) {
	var reminder models.Reminder
	err := r.db.Table("reminders").Where("id_reminder = ?",reminderId).Scan(&reminder).Error
	return reminder, err
}

func (r *ReminderDB) Delete(reminderId int) error {
	var reminder models.Reminder
	err := r.db.Table("reminders").Where("id_reminder = ? ", reminderId).First(&reminder).Error
	if err == nil {
		err = r.db.Table("reminders").Where("id_reminder = ? ", reminderId).Delete(reminder).Error
	}else{
		return err
	}
	return err
}

func (r *ReminderDB) Update(reminderId int, input models.UpdateReminder) error {
	err := r.db.Table("reminders").Where("id_reminder = ?", reminderId).Updates(models.UpdateReminder{
		Hint: input.Hint,
		UpdatedAt: time.Now().UTC(),
	}).Error
	return err
}


