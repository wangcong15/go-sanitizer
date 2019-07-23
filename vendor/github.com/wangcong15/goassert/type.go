package goassert

func allToFloat64(source interface{}) float64 {
	switch source.(type) {
	case int:
		return float64(source.(int))
	case uint8:
		return float64(source.(uint8))
	case uint16:
		return float64(source.(uint16))
	case uint32:
		return float64(source.(uint32))
	case uint64:
		return float64(source.(uint64))
	case int8:
		return float64(source.(int8))
	case int16:
		return float64(source.(int16))
	case int32:
		return float64(source.(int32))
	case float32:
		return float64(source.(float32))
	case float64:
		return source.(float64)
	case uintptr:
		return float64(source.(uintptr))
	case uint:
		return float64(source.(uint))
	default:
		return 0
	}
}
