package ec_point

import (
	"math/big"
	"reflect"
	"testing"
)

func TestECPoint_AddECPoints(t *testing.T) {
	g := new(ECPoint)
	pointX := new(big.Int)
	pointY := new(big.Int)
	pointX, _ = pointX.SetString("17050763248295559795867271927207444725689095732948540793148880382012", 10)
	pointY, _ = pointY.SetString("23268242352825710724048039655494175553214710720127294858798603639245", 10)

	type args struct {
		a ECPoint
		b ECPoint
	}
	tests := []struct {
		name string
		args args
		want ECPoint
	}{
		{
			name: "Add the same base point",
			args: args{
				a: g.BasePointGGet(),
				b: g.BasePointGGet(),
			},
			want: g.DoubleECPoints(g.BasePointGGet()),
		},
		{
			name: "Multiplication",
			args: args{
				a: g.ScalarMult(g.BasePointGGet(), *big.NewInt(5)),
				b: g.ScalarMult(g.BasePointGGet(), *big.NewInt(10)),
			},
			want: g.ScalarMult(g.BasePointGGet(), *big.NewInt(15)),
		},
		{
			name: "Enter value",
			args: args{
				a: g.ScalarMult(g.BasePointGGet(), *big.NewInt(7)),
				b: g.ScalarMult(g.BasePointGGet(), *big.NewInt(12)),
			},
			want: g.ECPointGen(pointX, pointY),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.AddECPoints(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddECPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECPoint_BasePointGGet(t *testing.T) {
	g := new(ECPoint)
	pointX := new(big.Int)
	pointY := new(big.Int)
	pointX, _ = pointX.SetString("19277929113566293071110308034699488026831934219452440156649784352033", 10)
	pointY, _ = pointY.SetString("19926808758034470970197974370888749184205991990603949537637343198772", 10)

	tests := []struct {
		name string
		want ECPoint
	}{
		{
			name: "Get base point",
			want: g.ECPointGen(pointX, pointY),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.BasePointGGet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasePointGGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECPoint_DoubleECPoints(t *testing.T) {
	g := new(ECPoint)
	pointX := new(big.Int)
	pointY := new(big.Int)
	pointX, _ = pointX.SetString("18394223645928780833749682104922684895691658099010032856578518073597", 10)
	pointY, _ = pointY.SetString("6079806569424875768420363755065073818836240443816190687503610242399", 10)

	type args struct {
		a ECPoint
	}
	tests := []struct {
		name string
		args args
		want ECPoint
	}{
		{
			name: "Double point",
			args: args{
				a: g.ScalarMult(g.BasePointGGet(), *big.NewInt(5)),
			},
			want: g.ECPointGen(pointX, pointY),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.DoubleECPoints(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleECPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECPoint_ECPointToString(t *testing.T) {
	g := new(ECPoint)
	pointX := new(big.Int)
	pointY := new(big.Int)
	pointX, _ = pointX.SetString("123245456477546785468546854678456854864796505", 10)
	pointY, _ = pointY.SetString("487567984375845689087954967747554687569056998", 10)

	type args struct {
		point ECPoint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Point from value",
			args: args{
				point: g.ECPointGen(pointX, pointY),
			},
			want: "(123245456477546785468546854678456854864796505; 487567984375845689087954967747554687569056998)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := g.ECPointToString(tt.args.point); gotS != tt.want {
				t.Errorf("ECPointToString() = %v, want %v", gotS, tt.want)
			}
		})
	}
}

func TestECPoint_IsOnCurveCheck(t *testing.T) {
	g := new(ECPoint)
	pointX := new(big.Int)
	pointY := new(big.Int)
	pointX, _ = pointX.SetString("18394223645928780833749682104922684895691658099010032856578518073597", 10)
	pointY, _ = pointY.SetString("6079806569424875768420363755065073818836240443816190687503610242399", 10)

	type args struct {
		a ECPoint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "On curve",
			args: args{
				a: g.ECPointGen(pointX, pointY),
			},
			want: true,
		},
		{
			name: "Not on curve",
			args: args{
				a: g.ECPointGen(big.NewInt(54768457), big.NewInt(54768457)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.IsOnCurveCheck(tt.args.a); got != tt.want {
				t.Errorf("IsOnCurveCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECPoint_ScalarMult(t *testing.T) {
	g := new(ECPoint)
	k1 := big.NewInt(10)
	k2 := big.NewInt(100)
	pointX := new([6]*big.Int)

	for i := range pointX {
		pointX[i] = big.NewInt(0)
	}

	pointX[0], _ = pointX[0].SetString("18394223645928780833749682104922684895691658099010032856578518073597", 10)
	pointX[1], _ = pointX[1].SetString("6079806569424875768420363755065073818836240443816190687503610242399", 10)
	pointX[2], _ = pointX[2].SetString("3536454026327045433155366519977330985440981474337987928327401155591", 10)
	pointX[3], _ = pointX[3].SetString("23935576693874995528625878114524732944794631339874353685522377669282", 10)
	pointX[4], _ = pointX[4].SetString("6655078886200176356515534105337266884694580077117987390971476538482", 10)
	pointX[5], _ = pointX[5].SetString("11191284385278881117358119562496949647030741040745570630483052779103", 10)

	type args struct {
		a ECPoint
		k big.Int
	}
	tests := []struct {
		name string
		args args
		want ECPoint
	}{
		{
			name: "Multiply by 10",
			args: args{
				a: g.ECPointGen(pointX[0], pointX[1]),
				k: *k1,
			},
			want: g.ECPointGen(pointX[2], pointX[3]),
		},
		{
			name: "Multiply by 100",
			args: args{
				a: g.ECPointGen(pointX[0], pointX[1]),
				k: *k2,
			},
			want: g.ECPointGen(pointX[4], pointX[5]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.ScalarMult(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScalarMult() = %v, want %v", got, tt.want)
			}
		})
	}
}
