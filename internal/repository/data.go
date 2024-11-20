package repository

import (
	"beexamxchat/domain"
	"errors"
	"strings"
)

type DataRepository struct {
	allData []domain.Data
}

func NewDataRepository(allData []domain.Data) domain.DataRepository {
	return &DataRepository{
		allData: allData,
	}
}

func (d *DataRepository) Create(data domain.Data) domain.Data {
	d.allData = append(d.allData, data)
	return data
}

func (d *DataRepository) Delete(code string) error {
	for i, data := range d.allData {
		if data.Code == code {
			d.allData = append(d.allData[:i], d.allData[i+1:]...)
			return nil
		}
	}
	return errors.New("data not found")
}

func (d *DataRepository) ReadAll() []domain.Data {
	return d.allData
}

func (d *DataRepository) ReadByCode(code string) (*domain.Data, error) {
	for _, data := range d.allData {
		if data.Code == code {
			return &data, nil
		}
	}
	return nil, errors.New("data not found")
}

func (d *DataRepository) Update(code string, updatedData domain.Data) error {
	for i, data := range d.allData {
		if data.Code == code {
			d.allData[i] = updatedData
			return nil
		}
	}
	return errors.New("data not found")
}

func (d *DataRepository) FilterBy(model string, tech string) []domain.Data {
	var filteredData []domain.Data

	techFilters := strings.Split(tech, ",")

	for _, d := range d.allData {

		modelMatch := model == "" || strings.EqualFold(d.Model, model)

		techMatch := true
		if tech != "" {
			techMatch = false
			for _, filter := range techFilters {
				if strings.Contains(strings.ToLower(d.Tech), strings.TrimSpace(strings.ToLower(filter))) {
					techMatch = true
					break
				}
			}
		}

		if modelMatch && techMatch {
			filteredData = append(filteredData, d)
		}
	}

	return filteredData
}
