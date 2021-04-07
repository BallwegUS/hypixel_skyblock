package api

type ProfileReturn struct {
	Success bool
	Profile Profile
}

type Profile struct {
	ID                string `json:"profile_id"`
	Members           []interface{}
	CommunityUpgrades []interface{} `json:"community_upgrades"`
}
