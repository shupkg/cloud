package p

//string
//int
//int8
//int16
//int32
//int64
//uint
//uint8
//uint16
//uint32
//uint64
//float32
//float64
//bool

func StringP(value string) *string {
	return &value
}

func IntP(value int) *int {
	return &value
}
func Int8P(value int8) *int8 {
	return &value
}
func Int16P(value int16) *int16 {
	return &value
}
func Int32P(value int32) *int32 {
	return &value
}
func Int64P(value int64) *int64 {
	return &value
}
func UintP(value uint) *uint {
	return &value
}
func Uint8P(value uint8) *uint8 {
	return &value
}
func Uint16P(value uint16) *uint16 {
	return &value
}
func Uint32P(value uint32) *uint32 {
	return &value
}
func Uint64P(value uint64) *uint64 {
	return &value
}
func Float32P(value float32) *float32 {
	return &value
}
func Float64P(value float64) *float64 {
	return &value
}
func BoolP(value bool) *bool {
	return &value
}

func StringPs(values []string) []*string {
	ps := make([]*string, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func IntPs(values []int) []*int {
	ps := make([]*int, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Int8Ps(values []int8) []*int8 {
	ps := make([]*int8, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Int16Ps(values []int16) []*int16 {
	ps := make([]*int16, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Int32Ps(values []int32) []*int32 {
	ps := make([]*int32, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Int64Ps(values []int64) []*int64 {
	ps := make([]*int64, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func UintPs(values []uint) []*uint {
	ps := make([]*uint, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Uint8Ps(values []uint8) []*uint8 {
	ps := make([]*uint8, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Uint16Ps(values []uint16) []*uint16 {
	ps := make([]*uint16, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Uint32Ps(values []uint32) []*uint32 {
	ps := make([]*uint32, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Uint64Ps(values []uint64) []*uint64 {
	ps := make([]*uint64, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Float32Ps(values []float32) []*float32 {
	ps := make([]*float32, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func Float64Ps(values []float64) []*float64 {
	ps := make([]*float64, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}

func BoolPs(values []bool) []*bool {
	ps := make([]*bool, len(values))
	for i := range ps {
		ps[i] = &values[i]
	}
	return ps
}
