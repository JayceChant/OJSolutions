package main

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s:    "0000",
			want: []string{"0.0.0.0"},
		},
		{
			s:    "010010",
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			s:    "172162541",
			want: []string{"17.216.25.41", "17.216.254.1", "172.16.25.41", "172.16.254.1", "172.162.5.41", "172.162.54.1"},
		},
		{
			s:    "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s:    "3923156",
			want: []string{"3.9.23.156", "3.9.231.56", "3.92.3.156", "3.92.31.56", "39.2.3.156", "39.2.31.56", "39.23.1.56", "39.23.15.6", "39.231.5.6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
