package claim

type Claim string

func ClaimMapIntoArray(arrayOfClaims map[Claim]Claim) []Claim {
	var array []Claim
	for _, claim := range arrayOfClaims {
		array = append(array, claim)
	}
	return array
}
