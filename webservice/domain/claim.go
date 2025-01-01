package domain

type Claim string

func IntoArray(claimMap map[Claim]Claim) []Claim {
	var array []Claim
	for _, claim := range claimMap {
		array = append(array, claim)
	}
	return array
}

func IntoMap(claimArray []Claim) map[Claim]Claim {
	claimMap := make(map[Claim]Claim)
	for _, claim := range claimArray {
		claimMap[claim] = claim
	}
	return claimMap
}
