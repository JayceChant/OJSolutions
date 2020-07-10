package main

import "testing"

var tests3 = []struct {
	name string
	num  string
	k    int
	want string
}{
	{
		name: "1",
		num:  "4321",
		k:    4,
		want: "1342",
	},
	{
		name: "2",
		num:  "100",
		k:    1,
		want: "010",
	},
	{
		name: "3",
		num:  "36789",
		k:    1000,
		want: "36789",
	},
	{
		name: "4",
		num:  "22",
		k:    22,
		want: "22",
	},
	{
		name: "5",
		num:  "9438957234785635408",
		k:    23,
		want: "0345989723478563548",
	},
}

func Test_minInteger(t *testing.T) {
	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			// if got := minIntegerBF(tt.num, tt.k); got != tt.want {
			// 	t.Errorf("minIntegerBF() = %v, want %v", got, tt.want)
			// }

			// if got := minIntegerBF2(tt.num, tt.k); got != tt.want {
			// 	t.Errorf("minIntegerBF2() = %v, want %v", got, tt.want)
			// }

			// if got := minIntegerRec(tt.num, tt.k); got != tt.want {
			// 	t.Errorf("minIntegerRec() = %v, want %v", got, tt.want)
			// }
			// if got := minIntegerST(tt.num, tt.k); got != tt.want {
			// 	t.Errorf("minIntegerST() = %v, want %v", got, tt.want)
			// }
			if got := minIntegerBIT(tt.num, tt.k); got != tt.want {
				t.Errorf("minIntegerBIT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minIntegerBF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests3 {
			minIntegerBF(tt.num, tt.k)
		}
	}
}

func Benchmark_minIntegerBF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests3 {
			minIntegerBF2(tt.num, tt.k)
		}
	}
}

func Benchmark_minIntegerRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests3 {
			minIntegerRec(tt.num, tt.k)
		}
	}
}

func Benchmark_minIntegerST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests3 {
			minIntegerST(tt.num, tt.k)
		}
	}
}

func Benchmark_minIntegerBIT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests3 {
			minIntegerBIT(tt.num, tt.k)
		}
	}
}
