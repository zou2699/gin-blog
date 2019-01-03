package model

type Auth struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) error {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		return err
	}
	return nil
}
