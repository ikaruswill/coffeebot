package koffea

type District struct {
	Name      string
	Locations []Location
}

type Location struct {
	Name string
}

type DistrictService interface {
	GetByName(name string) *District
	List() (*[]District, error)
}
