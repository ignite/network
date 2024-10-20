package types

func GetTierFromID(tierList []Tier, tierID uint64) (Tier, bool) {
	for _, tier := range tierList {
		if tier.TierId == tierID {
			return tier, true
		}
	}

	return Tier{}, false
}
