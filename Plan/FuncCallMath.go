package Plan

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/xitongsys/guery/Util"
)

func NewAbsFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "ABS",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			if len(es) < 1 {
				return Util.UNKNOWNTYPE, fmt.Errorf("not enough parameters in Abs")
			}
			return es[0].GetType(md)
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return nil, fmt.Errorf("not enough parameters in Abs")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return nil, err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return nil, fmt.Errorf("type cann't use ABS function")
			case Util.FLOAT64:
				v := tmp.(float64)
				if v < 0 {
					v *= -1
				}
				return v, nil
			case Util.FLOAT32:
				v := tmp.(float32)
				if v < 0 {
					v *= -1
				}
				return v, nil
			case Util.INT64:
				v := tmp.(int64)
				if v < 0 {
					v *= -1
				}
				return v, nil
			case Util.INT32:
				v := tmp.(int32)
				if v < 0 {
					v *= -1
				}
				return v, nil
			default:
				return nil, fmt.Errorf("unknown type")
			}
		},
	}
	return res
}

func NewSqrtFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "SQRT",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in SQRT")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return nil, err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return nil, fmt.Errorf("type cann't use SQRT function")

			default:
				return math.Sqrt(Util.ToFloat64(tmp)), nil
			}
		},
	}
	return res
}

func NewPowFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "POW",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 2 {
				return false
			}

			if es[0].IsAggregate() || es[1].IsAggregate() {
				return true
			}
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 2 {
				return float64(0), fmt.Errorf("not enough parameters in POW")
			}
			var (
				err        error
				tmp1, tmp2 interface{}
				t1         *ExpressionNode = Expressions[0]
				t2         *ExpressionNode = Expressions[1]
			)

			if tmp1, err = t1.Result(input); err != nil {
				return float64(0), err
			}
			if tmp2, err = t2.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp1) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use POW function")

			default:
				switch Util.TypeOf(tmp2) {
				case Util.STRING, Util.BOOL, Util.TIMESTAMP:
					return float64(0), fmt.Errorf("type cann't use POW function")
				}
				v1, v2 := Util.ToFloat64(tmp1), Util.ToFloat64(tmp2)
				return math.Pow(v1, v2), nil
			}
		},
	}
	return res
}

func NewLogFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "LOG",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 2 {
				return false
			}
			if es[0].IsAggregate() || es[1].IsAggregate() {
				return true
			}
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 2 {
				return float64(0), fmt.Errorf("not enough parameters in LOG")
			}
			var (
				err        error
				tmp1, tmp2 interface{}
				t1         *ExpressionNode = Expressions[0]
				t2         *ExpressionNode = Expressions[1]
			)

			if tmp1, err = t1.Result(input); err != nil {
				return float64(0), err
			}
			if tmp2, err = t2.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp1) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use LOG function")

			default:
				switch Util.TypeOf(tmp2) {
				case Util.STRING, Util.BOOL, Util.TIMESTAMP:
					return float64(0), fmt.Errorf("type cann't use LOG function")
				}
				v1, v2 := Util.ToFloat64(tmp1), Util.ToFloat64(tmp2)
				return math.Log(v1) / math.Log(v2), nil
			}
		},
	}
	return res
}

func NewLog10Func() *GueryFunc {
	res := &GueryFunc{
		Name: "LOG",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in LOG10")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use LOG10 function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Log10(v), nil
			}
		},
	}
	return res
}

func NewLog2Func() *GueryFunc {
	res := &GueryFunc{
		Name: "LOG",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in LOG10")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use LOG10 function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Log2(v), nil
			}
		},
	}
	return res
}

func NewLnFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "LN",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in LOG10")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use LOG10 function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Log(v), nil
			}
		},
	}
	return res
}

func NewCeilFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "CEIL",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in CEIL")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use CEIL function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Ceil(v), nil
			}
		},
	}
	return res
}

func NewFloorFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "FLOOR",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in FLOOR")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use FLOOR function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Floor(v), nil
			}
		},
	}
	return res
}

func NewRoundFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "ROUND",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in ROUND")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use ROUND function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Round(v), nil
			}
		},
	}
	return res
}

func NewSinFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "SIN",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in SIN")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use SIN function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Sin(v), nil
			}
		},
	}
	return res
}

func NewCosFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "COS",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in COS")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use COS function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Cos(v), nil
			}
		},
	}
	return res
}

func NewTanFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "TAN",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in TAN")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use TAN function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Tan(v), nil
			}
		},
	}
	return res
}

func NewASinFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "ASIN",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in ASIN")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use ASIN function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Asin(v), nil
			}
		},
	}
	return res
}

func NewACosFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "ACOS",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in ACOS")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use ACOS function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Acos(v), nil
			}
		},
	}
	return res
}

func NewATanFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "ATAN",
		IsAggregate: func(es []*ExpressionNode) bool {
			if len(es) < 1 {
				return false
			}
			return es[0].IsAggregate()
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			if len(Expressions) < 1 {
				return float64(0), fmt.Errorf("not enough parameters in ATAN")
			}
			var (
				err error
				tmp interface{}
				t   *ExpressionNode = Expressions[0]
			)

			if tmp, err = t.Result(input); err != nil {
				return float64(0), err
			}

			switch Util.TypeOf(tmp) {
			case Util.STRING, Util.BOOL, Util.TIMESTAMP:
				return float64(0), fmt.Errorf("type cann't use ATAN function")

			default:
				v := Util.ToFloat64(tmp)
				return math.Atan(v), nil
			}
		},
	}
	return res
}

func NewRandomFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "RANDOM",
		IsAggregate: func(es []*ExpressionNode) bool {
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			return rand.Float64(), nil
		},
	}
	return res
}

func NewEFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "E",
		IsAggregate: func(es []*ExpressionNode) bool {
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			return math.E, nil
		},
	}
	return res
}

func NewPiFunc() *GueryFunc {
	res := &GueryFunc{
		Name: "PI",
		IsAggregate: func(es []*ExpressionNode) bool {
			return false
		},

		GetType: func(md *Util.Metadata, es []*ExpressionNode) (Util.Type, error) {
			return Util.FLOAT64, nil
		},

		Result: func(input *Util.RowsBuffer, Expressions []*ExpressionNode) (interface{}, error) {
			return math.Pi, nil
		},
	}
	return res
}
