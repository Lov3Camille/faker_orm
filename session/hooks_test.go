package session

import (
	"faker_orm/log"
	"testing"
)

type Account struct {
	ID       int `fakerorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before insert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(
		&Account{
			ID:       1,
			Password: "123456",
		},
		&Account{
			ID:       2,
			Password: "qwerty",
		})

	u := &Account{}

	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
