package utils

// Define the MilitaryRank enum type
type MilitaryRank int

// Define the constants for the possible MilitaryRank values
const (
	CommanderInChief MilitaryRank = iota + 1
	Commander
	General
	BrigadierGeneral
	Colonel
	LieutenantColonel
	Major
	Captain
	Lieutenant
	Soldier
)

func (m MilitaryRank) String() string {
	switch m {
	case CommanderInChief:
		return "軍旗"
	case Commander:
		return "司令"
	case General:
		return "軍長"
	case BrigadierGeneral:
		return "師長"
	case Colonel:
		return "旅長"
	case LieutenantColonel:
		return "團長"
	case Major:
		return "營長"
	case Captain:
		return "連長"
	case Lieutenant:
		return "排長"
	case Soldier:
		return "工兵"
	default:
		return "Unknown"
	}
}
