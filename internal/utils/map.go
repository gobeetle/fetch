package utils

import (
	"fmt"
	"math"
	"strconv"
)

// CategorizeMap takes a map of strings to any and organizes them into type-specific maps
func CategorizeMap(claimsMap map[string]any) (
	map[string]int32,
	map[string]int64,
	map[string]float64,
	map[string]bool,
	map[string]string,
) {
	int32Map := make(map[string]int32)
	int64Map := make(map[string]int64)
	float64Map := make(map[string]float64)
	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)

	for key, value := range claimsMap {
		switch v := value.(type) {
		case string:
			stringMap[key] = v
			// Try to parse as int64
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				if i >= math.MinInt32 && i <= math.MaxInt32 {
					int32Map[key] = int32(i)
				}
				int64Map[key] = i
			}
			// Try to parse as float64
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				float64Map[key] = f
			}
			// Try to parse as bool
			if b, err := strconv.ParseBool(v); err == nil {
				boolMap[key] = b
			}

		case float64:
			float64Map[key] = v
			// Check if it's an integer value that fits in int64
			if v == float64(int64(v)) {
				int64Map[key] = int64(v)
				// Check if it also fits in int32
				if v >= math.MinInt32 && v <= math.MaxInt32 {
					int32Map[key] = int32(v)
				}
			}

		case bool:
			boolMap[key] = v

		case int:
			int64Map[key] = int64(v)
			if v >= math.MinInt32 && v <= math.MaxInt32 {
				int32Map[key] = int32(v)
			}

		case int32:
			int32Map[key] = v
			int64Map[key] = int64(v)

		case int64:
			int64Map[key] = v
			if v >= math.MinInt32 && v <= math.MaxInt32 {
				int32Map[key] = int32(v)
			}

		case float32:
			float64Map[key] = float64(v)

		default:
			// For other types, convert to string representation
			stringMap[key] = fmt.Sprintf("%v", v)
		}
	}

	return int32Map, int64Map, float64Map, boolMap, stringMap
}
