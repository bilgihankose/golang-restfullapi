package dataloaders

import (
	"encoding/json"
	"github.com/bilgihankose/restfullapi/models"
	"github.com/bilgihankose/restfullapi/utils"
)

func LoadUsers() []models.User{
 bytes, _ := utils.ReadFile("../json/users.json")
 var data []models.User
 utils.CheckError(json.Unmarshal([]byte(bytes), &data))
 return data
}

func LoadInterest() []models.Interest {
	bytes, _ := utils.ReadFile("../json/interest.json")
	var data []models.Interest
	utils.CheckError(json.Unmarshal([]byte(bytes), &data))
	return data
}

func LoadInterestModels() []models.InterestMapping {
	bytes, _ := utils.ReadFile("../json/userInterestMappings.json")
	var data []models.InterestMapping
	utils.CheckError(json.Unmarshal([]byte(bytes), &data))
	return data
}