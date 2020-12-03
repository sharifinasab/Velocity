package test

import (
	"testing"
	"time"

	"KOHO/repository"
	"KOHO/util"
)

func TestWeeklyLimit_Validate(t *testing.T) {
	type args struct {
		transactionTime time.Time
		amount          float64
	}

	weeklyLimit := repository.NewWeeklyLimitRule(time.Now())
	weekStartTime := util.GetWeekStartTime(time.Now())

	tests := []struct {
		name string
		wl   *repository.WeeklyLimit
		args args
		want bool
	}{
		{
			name: "tc1",
			wl:   weeklyLimit,
			args: args{transactionTime: weekStartTime, amount: 5000},
			want: true,
		},
		{
			name: "tc2",
			wl:   weeklyLimit,
			args: args{transactionTime: weekStartTime.AddDate(0, 0, 1),
				amount: 5000,
			},
			want: true,
		},
		{
			name: "tc3",
			wl:   weeklyLimit,
			args: args{transactionTime: weekStartTime.AddDate(0, 0, 2),
				amount: 5000,
			},
			want: true,
		},
		{
			name: "tc4",
			wl:   weeklyLimit,
			args: args{transactionTime: weekStartTime.AddDate(0, 0, 3),
				amount: 5000,
			},
			want: true,
		},
		{
			name: "tc5",
			wl:   weeklyLimit,
			args: args{transactionTime: weekStartTime.AddDate(0, 0, 4),
				amount: 0.1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wl.Validate(tt.args.amount); got != tt.want {
				t.Errorf("WeeklyLimit.Validate() = %v, want %v", got, tt.want)
			} else {
				weeklyLimit.UpdateQuota(tt.args.amount)
			}
		})
	}
}
