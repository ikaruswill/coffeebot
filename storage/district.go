package storage

import (
	"github.com/ikaruswill/koffea/koffea"
	"gorm.io/gorm/clause"
)

type Location struct {
	Name       string
	DistrictID uint
}

type District struct {
	Name      string
	Locations []Location
}

type DistrictService struct {
	Storage Storage
}

func (d *DistrictService) List() (*[]koffea.District, error) {
	panic("not implemented")
}

func (d *DistrictService) GetByName(name string) *koffea.District {
	panic("not implemented")
}

func (d *DistrictService) initDB() {
	districts := []District{
		{
			Name: "Marina Area",
			Locations: []Location{
				{Name: "Boat Quay"},
				{Name: "Chinatown"},
				{Name: "Havelock Road"},
				{Name: "Marina Square"},
				{Name: "Raffles Place"},
				{Name: "Suntec City"},
			},
		},
		{
			Name: "CBD",
			Locations: []Location{
				{Name: "Anson Road"},
				{Name: "Chinatown"},
				{Name: "Neil Road"},
				{Name: "Raffles Place"},
				{Name: "Shenton Way"},
				{Name: "Tanjong Pagar"},
			},
		},
		{
			Name: "Central South",
			Locations: []Location{
				{Name: "Alexandra Road"},
				{Name: "Tiong Bahru"},
				{Name: "Queenstown"},
			},
		},
		{
			Name: "Keppel",
			Locations: []Location{
				{Name: "Keppel"},
				{Name: "Mount Faber"},
				{Name: "Sentosa"},
				{Name: "Telok Blangah"},
			},
		},
		{
			Name: "South West",
			Locations: []Location{
				{Name: "Buona Vista"},
				{Name: "Dover"},
				{Name: "Pasir Panjang"},
				{Name: "West Coast"},
			},
		},
		{
			Name: "City Hall",
			Locations: []Location{
				{Name: "City Hall"},
				{Name: "High Street"},
				{Name: "North Bridge Road"},
			},
		},
		{
			Name: "Beach Road",
			Locations: []Location{
				{Name: "Beach Road"},
				{Name: "Bencoolen Road"},
				{Name: "Bugis"},
				{Name: "Rochor"},
			},
		},
		{
			Name: "Little India",
			Locations: []Location{
				{Name: "Little India"},
				{Name: "Farrer Park"},
				{Name: "Serangoon Road"},
			},
		},
		{
			Name: "Orchard",
			Locations: []Location{
				{Name: "Cairnhill"},
				{Name: "Killiney"},
				{Name: "Leonie Hill"},
				{Name: "Orchard"},
				{Name: "Oxley"},
			},
		},
		{
			Name: "Tanglin",
			Locations: []Location{
				{Name: "Balmoral"},
				{Name: "Bukit Timah"},
				{Name: "Grange Road"},
				{Name: "Holland"},
				{Name: "Orchard Boulevard"},
				{Name: "River Valley"},
				{Name: "Tanglin Road"},
			},
		},
		{
			Name: "Newton",
			Locations: []Location{
				{Name: "Chancery"},
				{Name: "Bukit Timah"},
				{Name: "Dunearn Road"},
				{Name: "Newton"},
			},
		},
		{
			Name: "Toa Payoh",
			Locations: []Location{
				{Name: "Balestier"},
				{Name: "Moulmein"},
				{Name: "Novena"},
				{Name: "Toa Payoh"},
			},
		},
		{
			Name: "Central East",
			Locations: []Location{
				{Name: "Potong Pasir"},
				{Name: "Macpherson"},
			},
		},
		{
			Name: "Eunos",
			Locations: []Location{
				{Name: "Eunos"},
				{Name: "Geylang"},
				{Name: "Kembangan"},
				{Name: "Paya Lebar"},
			},
		},
		{
			Name: "East Coast",
			Locations: []Location{
				{Name: "Katong"},
				{Name: "Marine Parade"},
				{Name: "Siglap"},
				{Name: "Tanjong Rhu"},
			},
		},
		{
			Name: "Central West",
			Locations: []Location{
				{Name: "Clementi"},
				{Name: "Upper Bukit Timah"},
				{Name: "Hume Avenue"},
			},
		},
	}
	d.Storage.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&districts)
}

func NewDistrictService(storage Storage) *DistrictService {
	s := DistrictService{
		Storage: storage,
	}
	s.Storage.AutoMigrate(
		&District{},
		&Location{},
	)
	s.initDB()
	return &s
}
