package test

import (
	"testing"
	"time"

	"KOHO/repository"
	"KOHO/util"
)

type Setup struct {
	dailyLimit   *repository.DailyLimit
	dayStartTime time.Time
}

func SetupTest(t time.Time) *Setup {
	return &Setup{
		dailyLimit:   repository.NewDailyLimitRule(util.GetDayStartTime(t)),
		dayStartTime: util.GetDayStartTime(t),
	}
}

func TestDailyLimitAmount_Validate_SingleDay(t *testing.T) {
	type args struct {
		transactionTime time.Time
		amount          float64
	}

	s := SetupTest(time.Now())

	tests := []struct {
		name string
		dl   *repository.DailyLimit
		args args
		want bool
	}{
		{
			name: "tc1",
			dl:   s.dailyLimit,
			args: args{transactionTime: s.dayStartTime,
				amount: 2000,
			},
			want: true,
		},
		{
			name: "tc2",
			dl:   s.dailyLimit,
			args: args{transactionTime: s.dayStartTime.Add(time.Hour),
				amount: 3000.1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dl.Validate(tt.args.amount); got != tt.want {
				t.Errorf("DailyLimit.Validate() = %v, want %v", got, tt.want)
			} else {
				tt.dl.UpdateQuota(tt.args.amount)
			}
		})
	}
}

func TestDailyLimitAmount_Validate_MultipleDay(t *testing.T) {
	type args struct {
		transactionTime time.Time
		amount          float64
	}

	s1 := SetupTest(time.Now())
	s2 := SetupTest(time.Now().AddDate(0, 0, 1))

	tests := []struct {
		name string
		dl   *repository.DailyLimit
		args args
		want bool
	}{
		{
			name: "tc1",
			dl:   s1.dailyLimit,
			args: args{transactionTime: s1.dayStartTime,
				amount: 5000,
			},
			want: true,
		},
		{
			name: "tc2",
			dl:   s2.dailyLimit,
			args: args{transactionTime: s2.dayStartTime.Add(time.Hour),
				amount: 3000,
			},
			want: true,
		},
		{
			name: "tc3",
			dl:   s2.dailyLimit,
			args: args{transactionTime: s2.dayStartTime.Add(time.Hour * 2),
				amount: 3000,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dl.Validate(tt.args.amount); got != tt.want {
				t.Errorf("DailyLimit.Validate() = %v, want %v", got, tt.want)
			} else {
				tt.dl.UpdateQuota(tt.args.amount)
			}
		})
	}
}

func TestDailyLimitCount_Validate(t *testing.T) {
	type args struct {
		transactionTime time.Time
		amount          float64
	}

	s := SetupTest(time.Now())

	tests := []struct {
		name string
		dl   *repository.DailyLimit
		args args
		want bool
	}{
		{
			name: "tc1",
			dl:   s.dailyLimit,
			args: args{transactionTime: s.dayStartTime,
				amount: 2000,
			},
			want: true,
		},
		{
			name: "tc2",
			dl:   s.dailyLimit,
			args: args{transactionTime: s.dayStartTime.Add(time.Hour),
				amount: 500,
			},
			want: true,
		},
		{
			name: "tc3",
			dl:   s.dailyLimit,
			args: args{transactionTime: s.dayStartTime.Add(time.Hour * 2),
				amount: 2000,
			},
			want: true,
		},
		{
			name: "tc4",
			dl:   s.dailyLimit,
			args: args{transactionTime: s.dayStartTime.Add(time.Hour * 3),
				amount: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dl.Validate(tt.args.amount); got != tt.want {
				t.Errorf("DailyLimit.Validate() = %v, want %v", got, tt.want)
			} else {
				tt.dl.UpdateQuota(tt.args.amount)
			}
		})
	}
}
