package model

type Scope string
type Scopes []Scope
const (
	CreateAliasScope Scope = "alias.create"
	ReadAliasScope Scope = "alias.read"
	UpdateAliasScope Scope = "alias.update"
	DeleteAliasScope Scope = "alias.delete"
)

func (ss Scopes) Contains(s Scope) bool {
	for _, scope := range ss {
		if scope == s {
			return true
		}
	}

	return false
}

func (s1 Scopes) ContainsAll(s2 Scopes) bool {
	for _, scope2 := range s2 {
		if !s1.Contains(scope2) {
			return false
		}
	}

	return true
}