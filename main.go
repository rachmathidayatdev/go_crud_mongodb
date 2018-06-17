package main

import (
	"fmt"
	"time"
	"github.com/rachmathidayatdev/go_crud_mongodb/config"
	"github.com/rachmathidayatdev/go_crud_mongodb/src/modules/profile/model"
	"github.com/rachmathidayatdev/go_crud_mongodb/src/modules/profile/repository"
)

//main
func main() {
	fmt.Println("Bikin crud di go nih")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfileById("U2", profileRepository)
	getAllProfiles(profileRepository)
}

//saveProfile
func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile

	p.ID = "U1"
	p.FirstName = "dadang"
	p.LastName = "dudung"
	p.Email = "dadang@dudung.com"
	p.Password = "dadangdudung"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil save...")
	}
}

//updateProfile
func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile

	p.ID = "U1"
	p.FirstName = "dadang"
	p.LastName = "dudung"
	p.Email = "dudung@dadang.com"
	p.Password = "dadangdudung"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil update...")
	}
}

//deleteProfile
func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil delete...")
	}
}

//getProfileById
func getProfileById(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	} 
	
	fmt.Println(profile)
}

//getAllProfiles
func getAllProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	} 
	
	for _, profile := range profiles {
		fmt.Println(profile)
	}
}
