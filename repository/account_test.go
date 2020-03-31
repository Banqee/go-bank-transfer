package repository

import (
	"reflect"
	"testing"

	"github.com/gsabadini/go-bank-transfer/infrastructure/database/stub"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gsabadini/go-bank-transfer/domain"
)

func TestAccountStore(t *testing.T) {
	type args struct {
		account domain.Account
	}

	tests := []struct {
		name          string
		args          args
		expected      domain.Account
		expectedError interface{}
		repository    Account
		expectedErr   bool
	}{
		{
			name:       "Success to create account",
			args:       args{account: domain.Account{}},
			expected:   domain.Account{},
			repository: NewAccount(stub.MongoHandlerSuccessStub{}),
		},
		{
			name:        "Error to create account",
			args:        args{account: domain.Account{}},
			repository:  NewAccount(stub.MongoHandlerErrorStub{}),
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repository.Store(tt.args.account)

			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Error: '%v' | ExpectedErr: '%v'", tt.name, err, tt.expectedErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, got, tt.expected)
			}
		})
	}
}

func TestAccountUpdate(t *testing.T) {
	type args struct {
		query  bson.M
		update bson.M
	}

	tests := []struct {
		name        string
		repository  Account
		args        args
		expectedErr bool
	}{
		{
			name:       "Success to account update",
			repository: NewAccount(stub.MongoHandlerSuccessStub{}),
			args: args{
				query:  bson.M{},
				update: bson.M{},
			},
		},
		{
			name:       "Error to account update",
			repository: NewAccount(stub.MongoHandlerErrorStub{}),
			args: args{
				query:  bson.M{},
				update: bson.M{},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repository.Update(tt.args.query, tt.args.update); (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Error: '%v' | ExpectedErr: '%v'", tt.name, err, tt.expectedErr)
			}
		})
	}
}

func TestAccountFindAll(t *testing.T) {
	tests := []struct {
		name        string
		repository  Account
		expected    []domain.Account
		expectedErr bool
	}{
		{
			name:       "Success in finding all accounts",
			repository: NewAccount(stub.MongoHandlerSuccessStub{}),

			expected: []domain.Account{},
		},
		{
			name:        "Error in finding all accounts",
			repository:  NewAccount(stub.MongoHandlerErrorStub{}),
			expected:    []domain.Account{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repository.FindAll()

			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Error: '%v' | ExpectedErr: '%v'", tt.name, err, tt.expectedErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, got, tt.expected)
			}
		})
	}
}

func TestAccountFindOne(t *testing.T) {
	type args struct {
		query bson.M
	}

	tests := []struct {
		name        string
		repository  Account
		args        args
		expected    *domain.Account
		expectedErr bool
	}{
		{
			name:       "Success to find account",
			repository: NewAccount(stub.MongoHandlerSuccessStub{}),
			args:       args{query: bson.M{}},
			expected:   &domain.Account{},
		},
		{
			name:        "Error to find account",
			repository:  NewAccount(stub.MongoHandlerErrorStub{}),
			args:        args{query: bson.M{}},
			expected:    &domain.Account{},
			expectedErr: true,
		},
		{
			name:        "Account not found",
			repository:  NewAccount(stub.MongoHandlerErrorStub{TypeErr: mgo.ErrNotFound}),
			args:        args{query: bson.M{}},
			expected:    &domain.Account{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repository.FindOne(tt.args.query)

			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Error: '%v' | ExpectedErr: '%v'", tt.name, err, tt.expectedErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, got, tt.expected)
			}
		})
	}
}

func TestAccountFindOneWithSelector(t *testing.T) {
	type args struct {
		query    bson.M
		selector interface{}
	}

	tests := []struct {
		name        string
		repository  Account
		args        args
		expected    domain.Account
		expectedErr bool
	}{
		{
			name:       "Success to find account with selector",
			repository: NewAccount(stub.MongoHandlerSuccessStub{}),
			expected:   domain.Account{},
		},
		{
			name:        "Error to find account with selector",
			repository:  NewAccount(stub.MongoHandlerErrorStub{}),
			expected:    domain.Account{},
			expectedErr: true,
		},
		{
			name:        "Account with selector not found",
			repository:  NewAccount(stub.MongoHandlerErrorStub{TypeErr: mgo.ErrNotFound}),
			args:        args{query: bson.M{}},
			expected:    domain.Account{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repository.FindOneWithSelector(tt.args.query, tt.args.selector)

			if (err != nil) != tt.expectedErr {
				t.Errorf("[TestCase '%s'] Error: '%v' | ExpectedErr: '%v'", tt.name, err, tt.expectedErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, got, tt.expected)
			}
		})
	}
}