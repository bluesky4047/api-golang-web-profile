package data

import "my-fiber-app/models"

// Interface agar bisa menukar sumber data
type Repository interface {
	GetHome() models.Home
	GetAbout() models.About
	GetProcess() models.Process
	GetPortfolio() models.PortfolioList
	GetBlog() models.Blogs
	GetServices() models.Services
	GetTestimonials() models.Testimonials
	GetContact() models.Contact
}

// Pilihan: DummyDataRepository atau DummyData1Repository
var activeRepo Repository = &DummyRepository{}

// Fungsi untuk mengganti repository aktif
func UseDummyData() {
	activeRepo = &DummyRepository{}
}

func UseDummyData1() {
	activeRepo = &Dummy1Repository{}
}

// Fungsi getter universal
func GetHome() models.Home               { return activeRepo.GetHome() }
func GetAbout() models.About             { return activeRepo.GetAbout() }
func GetProcess() models.Process         { return activeRepo.GetProcess() }
func GetPortfolio() models.PortfolioList { return activeRepo.GetPortfolio() }
func GetBlog() models.Blogs              { return activeRepo.GetBlog() }
func GetServices() models.Services       { return activeRepo.GetServices() }
func GetTestimonials() models.Testimonials { return activeRepo.GetTestimonials() }
func GetContact() models.Contact         { return activeRepo.GetContact() }
