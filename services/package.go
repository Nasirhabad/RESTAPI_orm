package services

import (
	"my-orm/database"
	"my-orm/models"
)

type PackageService struct{}

func (ps *PackageService) GetAll() ([]models.Package, error) {

	var packages []models.Package
	if err := database.DB.Find(&packages).Error; err != nil {
		return []models.Package{}, err

	}
	return packages, nil
}

func (ps *PackageService) GetByID(id string) (models.Package, error) {
	var pkg models.Package
	if err := database.DB.First(&pkg, "id = ?", id).Error; err != nil {
		return models.Package{}, err
	}
	return pkg, nil
}

func (ps *PackageService) Create(pkgReq models.PackageRequest) (models.Package, error) {
	var pkg models.Package = models.Package{
		Name:             pkgReq.Name,
		Sender:           pkgReq.Sender,
		Receiver:         pkgReq.Receiver,
		Senderlocation:   pkgReq.Senderlocation,
		Receiverlocation: pkgReq.Receiverlocation,
		Fee:              pkgReq.Fee,
		Weight:           pkgReq.Weight,
	}
	result := database.DB.Create(&pkg)
	if err := result.Error; err != nil {
		return models.Package{}, err
	}
	if err := result.Last(&pkg).Error; err != nil {
		return models.Package{}, err
	}
	return pkg, nil
}

func (ps *PackageService) Update(id string, pkgReq models.PackageRequest) (models.Package, error) {
	pkg, err := ps.GetByID(id)
	if err != nil {
		return models.Package{}, err
	}
	pkg.Name = pkgReq.Name
	pkg.Sender = pkgReq.Sender
	pkg.Receiver = pkgReq.Receiver
	pkg.Senderlocation = pkgReq.Senderlocation
	pkg.Receiverlocation = pkgReq.Receiverlocation
	pkg.Fee = pkgReq.Fee
	pkg.Weight = pkgReq.Weight
	if err := database.DB.Save(&pkg).Error; err != nil {
		return models.Package{}, err
	}

	return pkg, nil
}

func (ps *PackageService) Delete(id string) error {
	pkg, err := ps.GetByID(id)

	if err != nil {
		return err
	}
	if err := database.DB.Delete(&pkg).Error; err != nil {
		return err
	}
	return nil
}
