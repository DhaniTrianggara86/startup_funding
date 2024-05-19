package campaign

import "time"

type campaign struct {
	ID     int
	UserID int

	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BeckerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdateAt         time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdateAt   time.Time
}
