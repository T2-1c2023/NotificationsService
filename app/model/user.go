package model

type User struct {
	Id              int
	Fullname        string
	Mail            string
	Phone_Number    string
	Is_Athlete      bool
	Is_Trainer      bool
	Is_Verified     bool
	Photo_Id        string
	Expo_Push_Token string
}
