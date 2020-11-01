package p

func IsNil(values ...interface{}) bool {
	for _, value := range values {
		if value == nil {
			return true
		}
	}
	return false
}

func String(value *string, defs ...string) string {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return ""
}

func Bool(value *bool, defs ...bool) bool {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return false
}

func Int(value *int, defs ...int) int {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Int8(value *int8, defs ...int8) int8 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Int16(value *int16, defs ...int16) int16 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Int32(value *int32, defs ...int32) int32 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Int64(value *int64, defs ...int64) int64 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Uint(value *uint, defs ...uint) uint {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Uint8(value *uint8, defs ...uint8) uint8 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Uint16(value *uint16, defs ...uint16) uint16 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Uint32(value *uint32, defs ...uint32) uint32 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Uint64(value *uint64, defs ...uint64) uint64 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Float32(value *float32, defs ...float32) float32 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}

func Float64(value *float64, defs ...float64) float64 {
	if value != nil {
		return *value
	}
	if len(defs) > 0 {
		return defs[0]
	}
	return 0
}
