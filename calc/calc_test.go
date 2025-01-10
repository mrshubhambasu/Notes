package calc

import (
	"os"
	"runtime/pprof"
	"testing"
)

// func TestSum(t *testing.T) {
// 	type args struct {
// 		a int
// 		b int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"testpass",
// 			args{
// 				1,
// 				2,
// 			},
// 			3,
// 		},
// 		{
// 			"test with black value",
// 			args{},
// 			0,
// 		},
// 	}
// 	for _, tt := range tests {

// 		f := func(t *testing.T) {
// 			got := Sum(tt.args.a, tt.args.b)
// 			if got != tt.want {
// 				t.Errorf("Sum() = %v, want %v", got, tt.want)
// 			}
// 		}

// 		t.Run(tt.name, f)
// 	}
// }

// func BenchmarkAdd(b *testing.B) {
// 	b.ReportAllocs()
// 	for i := 0; i < b.N; i++ {
// 		Sum(i, i+1)
// 	}
// }

func BenchmarkCountVowels(b *testing.B) {
	s := GenerateStr(1000)

	// pprof.WriteHeapProfile(memProfile)

	for i := 0; i < b.N; i++ {
		CountVowels(s)
	}
	memProfile, _ := os.Create("memprofile.out")
	defer memProfile.Close()

	pprof.WriteHeapProfile(memProfile)
}

func BenchmarkCountVowelsOPT(b *testing.B) {
	s := GenerateStr(1000)

	for i := 0; i < b.N; i++ {
		CountVowelsOPT(s)
	}

}
