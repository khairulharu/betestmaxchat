package service

import (
	"beexamxchat/domain"
	"beexamxchat/dto"
)

type DataService struct {
	dataRepository domain.DataRepository
}

func NewDataService(dataRepository domain.DataRepository) domain.DataService {
	return &DataService{
		dataRepository: dataRepository,
	}
}

func (s *DataService) CreateData(data domain.Data) dto.Response {

	insertedData := s.dataRepository.Create(data)

	return dto.Response{
		Code:    200,
		Message: "Succecfuly Create Data",
		Data:    insertedData,
	}
}

func (s *DataService) DeleteData(code string) dto.Response {
	if err := s.dataRepository.Delete(code); err != nil {
		return dto.Response{
			Code:    400,
			Message: "Bad Request check code",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "Succesfuly Deleted Data",
	}
}

func (s *DataService) GetAllData() dto.Response {
	var allDataResponse []dto.AllData

	allDataFromRepository := s.dataRepository.ReadAll()

	for _, data := range allDataFromRepository {
		allDataResponse = append(allDataResponse, dto.AllData{
			Code:   data.Code,
			Name:   data.Name,
			Model:  data.Model,
			Tech:   data.Tech,
			Status: data.Status,
		})
	}

	return dto.Response{
		Code:    200,
		Message: "Succesfuly Get All data",
		Data:    allDataResponse,
	}
}

func (s *DataService) GetDataByCode(code string) dto.Response {
	dataFromRepository, err := s.dataRepository.ReadByCode(code)

	if err != nil {
		return dto.Response{
			Code:    400,
			Message: "Bad Request Check Code",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    200,
		Message: "Succesfuly Get Data",
		Data:    dataFromRepository,
	}
}

func (s *DataService) UpdateData(code string, updatedData domain.Data) dto.Response {

	if err := s.dataRepository.Update(code, updatedData); err != nil {
		return dto.Response{
			Code:    400,
			Message: "Bad Request Check Request Body",
			Error:   err.Error(),
		}
	}

	singleDataUpdated, _ := s.dataRepository.ReadByCode(code)

	return dto.Response{
		Code:    200,
		Message: "Succesfuly Updated Data",
		Data:    singleDataUpdated,
	}
}

func (s *DataService) FilterData(model string, tech string) dto.Response {
	filteredData := s.dataRepository.FilterBy(model, tech)

	if len(filteredData) == 0 {
		return dto.Response{
			Code:    404,
			Message: "Filtered Data Not Found",
			Data:    filteredData,
		}
	}

	return dto.Response{
		Code:    200,
		Message: "Succesfuly Filtered Data",
		Data:    filteredData,
	}
}
