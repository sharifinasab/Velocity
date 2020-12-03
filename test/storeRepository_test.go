package test

import (
	"testing"

	"KOHO/repository"
)

func TestStore_IsDuplicatedTransaction(t *testing.T) {
	type args struct {
		transactionID string
		accountID     string
	}
	store := repository.NewStore()
	tests := []struct {
		name string
		s    *repository.Store
		args args
		want bool
	}{
		{name: "tc1", s: store, args: args{transactionID: "1611", accountID: "12"}, want: false},
		{name: "tc2", s: store, args: args{transactionID: "1611", accountID: "12"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsDuplicatedTransaction(tt.args.transactionID, tt.args.accountID); got != tt.want {
				t.Errorf("Store.IsDuplicatedTransaction() = %v, want %v", got, tt.want)
			} else {
				tt.s.AddTransaction(tt.args.transactionID, tt.args.accountID)
			}
		})
	}
}
