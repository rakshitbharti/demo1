package mapstore

import (
	"errors"

	"assignment3/domain"
)

type MapStore struct {
	Store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{Store: make(map[string]domain.Customer)}
}

func (m *MapStore) Create(c domain.Customer) error {
	if _, ok := m.Store[c.ID]; ok {
		return errors.New("customer already exists")
	}
	m.Store[c.ID] = c
	return nil

}
func (m *MapStore) Update(s string, c domain.Customer) error {
	if _, ok := m.Store[s]; ok {
		m.Store[s] = c
	} else {
		return errors.New("Customer does not exists")
	}
	return nil
}
func (m *MapStore) Delete(s string) error {
	if _, ok := m.Store[s]; ok {
		delete(m.Store, s)
	} else {
		return errors.New("Customer doesn't exists")
	}
	return nil
}
func (m *MapStore) GetById(s string) (domain.Customer, error) {
	if val, ok := m.Store[s]; ok {
		return val, nil
	} else {
		return domain.Customer{}, errors.New("No such customer exists")
	}
}
func (m *MapStore) GetAll() ([]domain.Customer, error) {
	var Customers []domain.Customer
	for _, v := range m.Store {
		Customers = append(Customers, v)
	}
	return Customers, nil
}
