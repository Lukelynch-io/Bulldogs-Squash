package auth

type Claim string
type ClaimArray []Claim
type ClaimMap map[Claim]Claim

func (claimMap ClaimMap) IntoArray() ClaimArray {
	var array []Claim
	for _, claim := range claimMap {
		array = append(array, claim)
	}
	return array
}

func (claimArray ClaimArray) IntoMap() ClaimMap {
	claimMap := make(map[Claim]Claim)
	for _, claim := range claimArray {
		claimMap[claim] = claim
	}
	return claimMap
}

const CREATE_POST Claim = "CREATE_POST"
const DELETE_POST Claim = "DELETE_POST"

const ADMIN Claim = "Admin"
