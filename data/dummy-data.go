package data

import "my-fiber-app/models"

var Home = models.Home{
	Title: "Hello, I’m Brooklyn Gilbert",
	Description: "I'm a Freelance UI/UX Designer and Developer based in London, England. I strives to build immersive and beautiful web applications through carefully crafted code and user-centric design.",
	Experience: []models.Exp{
		{Name: "Experience", Value: "15 Y."},
		{Name: "Projects Completed", Value: "250+"},
		{Name: "Happy Clients", Value: "58"},
	},
	Img: "/data/img/person-CqOZwXV1.png",
}

var About = models.About{
	Img: "/data/img/person2-DvYIsw9m.png",
	Title: "I am Professional User Experience Designer",
	Description: "I design and develop services for customers specializing creating stylish, modern websites, web services and online stores. My passion is to design digital user experiences.I design and develop services for customers specializing creating stylish, modern websites, web services.",
	Links: []models.Link{
		{Name: "CV", Value: "#"},
		{Name: "Facebook", Value: "facebook.com"},
		{Name: "Instagram", Value: "instagram.com"},
		{Name: "LinkedIn", Value: "linkedin.com"},
	},
}

var Process = models.Process{
	Title: "Work Process",
	Description: "Driven by design and powered by code, I create digital interfaces that feel intuitive and perform seamlessly. Every layout, animation, and component is crafted with intention — merging usability with visual clarity, I blend clean design with efficient code to build engaging, user-friendly web experiences that stand out.",
	Details: []models.Detail{
		{Name: "1. Research", Description: "Design meets function in every pixel, blending clarity with intuitive motion."},
		{Name: "2. Analyze", Description: "Crafting clean, thoughtful interfaces where form flows seamlessly into function and clarity."},
		{Name: "3. Design", Description: "I design seamless digital experiences with precision, purpose, and a touch of elegance."},
		{Name: "4. Launch", Description: "I craft digital products where thoughtful design meets performance-driven, responsive development."},
	},
}

var Portfolio = models.PortfolioList{
	Portfolio: []models.Portfolio{
		{Img:"/data/img/card-1-D83uj-qZ.png",Title:"Product Admin Dashboard	", Description: "I focus on crafting smooth, responsive interfaces that balance aesthetic appeal with practical functionality.", Category: "UI-UX DESIGN" },
		{Img:"/data/img/card-2-BJ8-9N8h.png",Title:"Product Admin Dashboard	", Description: "Designed an intuitive dashboard for product management, emphasizing clarity and user efficiency.", Category: "UI-UX DESIGN" },
		{Img:"/data/img/card-3-Ka1ll87R.png",Title:"Product Admin Dashboard	", Description: "Developed a modern admin panel with a focus on usability and seamless navigation for end users and so on.", Category: "UI-UX DESIGN" },
		{Img:"/data/img/card-4-1lvzje-u.png",Title:"Product Admin Dashboard	", Description: "Created a responsive dashboard layout that adapts smoothly across devices and screen sizes and so on.", Category: "UI-UX DESIGN" },
		{Img:"/data/img/card-5-TrbU7d8r.png",Title:"Product Admin Dashboard	", Description: "Implemented interactive charts and widgets to visualize product data effectively for stakeholders.", Category: "UI-UX DESIGN" },
		{Img:"/data/img/card-6-CMi2awuo.png",Title:"Product Admin Dashboard	", Description: "Enhanced user experience by streamlining workflows and optimizing interface components and so on.", Category: "UI-UX DESIGN" },
	},
}

var Blog = models.Blogs{
	Blogs: []models.Blog{
		{Img: "https://themewagon.github.io/picto/assets/blog-1-9asS__eh.jpg",Title: "Designing Engaging User Interfaces for M...", Date: "22 Oct, 2020", TotalComments:"246 Comments" },
		{Img: "https://themewagon.github.io/picto/assets/blog-2-Jn4pr14V.jpg",Title: "Tips for Effective Dashboard Layouts and...", Date: "22 Oct, 2020", TotalComments:"246 Comments" },
		{Img: "https://themewagon.github.io/picto/assets/blog-3-lkBU4Y-F.jpg",Title: "How to Visualize Data for Better Product...", Date: "22 Oct, 2020", TotalComments:"246 Comments" },
		{Img: "https://themewagon.github.io/picto/assets/blog-4-C9dcE4M3.jpg",Title: "Responsive Design: Adapting to All Devic...", Date: "22 Oct, 2020", TotalComments:"246 Comments" },
		{Img: "https://themewagon.github.io/picto/assets/blog-2-Jn4pr14V.jpg",Title: "Streamlining Workflows with UI/UX Best P...", Date: "22 Oct, 2020", TotalComments:"246 Comments" },
		{Img: "https://themewagon.github.io/picto/assets/blog-1-9asS__eh.jpg",Title: "Optimizing Interface Components for Perf...", Date: "22 Oct, 2020", TotalComments:"246 Comments" },
	},
}

var Services = models.Services{
	Title: "What I do?",
	Description: "I specialize in designing user experiences, crafting engaging interfaces, and building robust web applications that deliver value and usability. My approach combines creativity and technical expertise to deliver solutions that are both visually appealing and highly functional for users.",
	Services: []models.Service{
		{Title: "User Experience (UX)",Description: "I design intuitive and enjoyable experiences by understanding user needs, conducting research, and creating wireframes and prototypes that enhance usability."},
		{Title: "User Interface (UI)",Description: "I craft visually appealing and consistent interfaces, focusing on layout, color, and typography to ensure a seamless and engaging user journey."},
		{Title: "Web Development",Description: "I build responsive and high-performance web applications using modern technologies, ensuring accessibility, scalability, and maintainability."},
	},
}

var Testimonials = models.Testimonials{
	Testimonials: []models.Testimonial{
		{Name: "Esther Howard", Job: "Managing Director, ABC Company", Comment: "From the initial consultation to the final delivery, every step was handled professionally. The end result was a product that not only met our needs but also impressed our stakeholders.Highly recommended!", Slug: "Working with this team was a fantastic experience. Their attention to detail and commitment to quality exceeded our expectations."},
		{Name: "Ali Haider", Job: "COO, XYZ Company", Comment: "The team demonstrated a deep understanding of our requirements and delivered a solution that was both visually appealing and highly functional. Communication was clear throughout the project.", Slug: "Their expertise in UI/UX design helped us transform our digital presence and improve user engagement."},
		{Name: "Elon Max", Job: "Managing Director, KFC Company", Comment: "They delivered our project on time and went above and beyond to ensure our satisfaction. The new features have made a significant difference for our users. We look forward to working together again.", Slug: "Professional, reliable, and creative-everything you want in a development partner."},
	},
}

var Contact = models.Contact{
	Title: "Let’s discuss your Project",
	Description: "I'm available for freelance work. Drop me a line if you have a project you think I'd be a good fit for.",
	Details: []models.Detail{
		{Name: "Address:", Description:"New Mexico, 31134"},
		{Name: "My Email:", Description:"mymail@mail.com"},
		{Name: "Call Me Now:", Description:"00-123 00000"},
		{Name: "Facebook", Description:"facebook.com"},
		{Name: "Instagram", Description:"instagram.com"},
		{Name: "linkedin", Description:"linkedin"},
	},
}

// DummyRepository adalah versi pertama dari sumber data dummy
type DummyRepository struct{}

// Implementasi interface Repository
func (r DummyRepository) GetAbout() models.About {
	return About
}

func (r DummyRepository) GetHome() models.Home {
	return Home
}

func (r DummyRepository) GetProcess() models.Process {
	return Process
}

func (r DummyRepository) GetPortfolio() models.PortfolioList {
	return Portfolio
}

func (r DummyRepository) GetBlog() models.Blogs {
	return Blog
}

func (r DummyRepository) GetServices() models.Services {
	return Services
}

func (r DummyRepository) GetTestimonials() models.Testimonials {
	return Testimonials
}

func (r DummyRepository) GetContact() models.Contact {
	return Contact
}

