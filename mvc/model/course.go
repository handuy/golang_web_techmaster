package model

type CourseInfo struct {
	Name string
	Price int
}

var Courses = []CourseInfo{
	{
		Name: "Xây dựng web báng hàng bằng Golang",
		Price: 2400000,
	},
	{
		Name: "Web cơ bản HTML5, CSS3, Javascript",
		Price: 600000,
	},
	{
		Name: "Thiết kế UI/UX",
		Price: 2400000,
	},
	{
		Name: "Python phân tích dữ liệu",
		Price: 6000000,
	},
	{
		Name: "Lộ trình đào tạo mobile 6 tháng cam kết việc làm",
		Price: 2400000,
	},
}