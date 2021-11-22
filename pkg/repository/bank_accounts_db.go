package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type BankAccountsDB struct {
	db *gorm.DB
}


func NewBankAccountsDB(db *gorm.DB) *BankAccountsDB{
	return &BankAccountsDB{db: db}
}

func (r *BankAccountsDB) Create(account models.BankAccounts) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)


	row :=	r.db.Table("bank_account").Create(&account)

	if err :=  row.Table("bank_account").Where("number_account = ?",account.NumberAccount).Find(&account).Error; err != nil{
		return 0,err
	}
	r.db.Table("bank_account").Where("number_account = ?",account.NumberAccount).Find(&account)
	id = account.IdBankAccount
	return id, nil
}

func (r *BankAccountsDB) GetAllUser(userId int) ([]models.BankAccounts, error) {
	var accounts []models.BankAccounts
	err := r.db.Table("bank_account").Where("client_id = ?", userId).Find(&accounts).Error
	return accounts,  err
}

func (r *BankAccountsDB) GetAll() ([]models.BankAccounts, error) {
	var accounts []models.BankAccounts
	err := r.db.Table("bank_account").Find(&accounts).Error
	return accounts,  err
}

func (r *BankAccountsDB) GetById(accountId int) (models.BankAccounts, error) {
	var account models.BankAccounts
	err := r.db.Table("bank_account").Where("id_bank_account = ?",accountId).Scan(&account).Error
	return account, err
}

func (r *BankAccountsDB) Delete(accountId int) error {
	var account models.BankAccounts
	err := r.db.Table("bank_account").Where("id_bank_account = ? ",accountId).First(&account).Error
	if err == nil {
		err = r.db.Table("bank_account").Where("id_bank_account = ? ",accountId).Delete(account).Error
	}else{
		return err
	}
	return err
}

func (r *BankAccountsDB) Update(accountId int, input models.UpdateBankAccounts) error {
	//var account models.BankAccounts
	err := r.db.Table("bank_account").Where("id_bank_account = ?", accountId).Updates(models.UpdateBankAccounts{
		NumberAccount: input.NumberAccount,
		ClientId: input.ClientId,
		StatusId: input.StatusId,
	}).Error
	return err
}


