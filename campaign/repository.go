package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]campaign, error)
	FindByUserID(userID int) ([]campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]campaign, error) {

	var campaigns []campaign

	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err

	}
	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]campaign, error) {
	var campaigns []campaign

	err := r.db.Where("user_id=?", userID).Preload("CampaignImages", "campaign_images.isprimary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByID(ID int) ([]campaign, error) {
	var campaigns []campaign

	err := r.db.Preload("User").Preload("CampaignImages.is_primary =1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
