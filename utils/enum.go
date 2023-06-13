package utils

import "database/sql/driver"

// Define the MilitaryRank enum type
type MilitaryRank int

// Define the constants for the possible MilitaryRank values
const (
	軍旗 MilitaryRank = iota + 1
	司令
	軍長
	師長
	旅長
	團長
	營長
	連長
	排長
	工兵
)

func (m MilitaryRank) String() string {
	switch m {
	case 軍旗:
		return "軍旗"
	case 司令:
		return "司令"
	case 軍長:
		return "軍長"
	// case 師長:
	// 	return "師長"
	// case 旅長:
	// 	return "旅長"
	// case 團長:
	// 	return "團長"
	// case 營長:
	// 	return "營長"
	// case 連長:
	// 	return "連長"
	// case 排長:
	// 	return "排長"
	// case 工兵:
	// 	return "工兵"
	default:
		return "Unknown"
	}
}

func (ct *MilitaryRank) Scan(value interface{}) error {
	*ct = MilitaryRank(value.(int))
	return nil
}

func (ct MilitaryRank) Value() (driver.Value, error) {
	return int(ct), nil
}
