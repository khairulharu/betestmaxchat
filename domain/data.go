package domain

import "beexamxchat/dto"

var AllData = []Data{
	{"rv1", "Rover #1", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua", "Mobil", "Kecerdasan Buatan, Mobil", "Progress"},
	{"tfx", "Transformer X", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua", "Transformation", "Mobil, Robot", "Active"},
	{"px1", "Pacifista 1", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua", "Humanoid", "AI, Robot", "Active"},
	{"ax8", "Android #8", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua", "Humanoid", "Cyborg", "Active"},
	{"ax7", "Android #7", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua", "Humanoid", "Cyborg", "Inactive"},
}

type Data struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Model       string `json:"model"`
	Tech        string `json:"tech"`
	Status      string `json:"status"`
}

type DataRepository interface {
	Create(data Data) Data
	ReadAll() []Data
	ReadByCode(code string) (*Data, error)
	Update(code string, updatedData Data) error
	Delete(code string) error
	FilterBy(model string, tech string) []Data
}

type DataService interface {
	CreateData(data Data) dto.Response
	DeleteData(code string) dto.Response
	GetAllData() dto.Response
	GetDataByCode(code string) dto.Response
	UpdateData(code string, updatedData Data) dto.Response
	FilterData(model string, tech string) dto.Response
}
