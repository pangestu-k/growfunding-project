package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdataeCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(campaignImageInput CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	// membuat slug
	slugString := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugString)

	newCamapaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCamapaign, err
	}

	return newCamapaign, nil
}

func (s *service) UpdataeCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("Campaign dengan id tersebut tidak ditemukan")
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("your not an owoner fo this campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return campaign, err
	}

	return updateCampaign, nil
}

func (s *service) SaveCampaignImage(campaignImageInput CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(campaignImageInput.CampaignID)

	if err != nil {
		return CampaignImage{}, errors.New("data campaign not Found")
	}

	if campaign.UserID != campaignImageInput.User.ID {
		return CampaignImage{}, errors.New("your not an owoner fo this campaign")
	}

	is_primary := 0
	fmt.Println(campaignImageInput.IsPrimary, "<- nilai nya")
	if campaignImageInput.IsPrimary {
		fmt.Println("kesini yak")
		is_primary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(campaignImageInput.CampaignID)

		if err != nil {
			return CampaignImage{}, err
		}
	} else {
		fmt.Println("gak kesini")
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = campaignImageInput.CampaignID
	campaignImage.FileName = fileLocation
	campaignImage.IsPrimary = is_primary

	newCampaign, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
