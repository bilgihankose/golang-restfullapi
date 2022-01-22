package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bilgihankose/restfullapi/dataloaders"
	"github.com/bilgihankose/restfullapi/models"
	"github.com/bilgihankose/restfullapi/utils"
	"net/http"
)

func Run()  {
	fmt.Println("Running server on :8080")
	http.HandleFunc("", Handler)
	utils.CheckError(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//page nesnesi
	page := models.Page{ID: 7, Name: "Kullanicilar", Description: "Kullanici Listesi", URI: "/users"}

	//dataloaders
	users := dataloaders.LoadUsers()
	interests := dataloaders.LoadInterest()
	interestMappings := dataloaders.LoadInterestMapping()
	//islem

	var newUsers []model.User

	for _, user := range users {

		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := models.UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))

}
