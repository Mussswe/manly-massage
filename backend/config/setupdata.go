package config

import (
	"backend/entity"
	"fmt"
	"time"
)

// SetupData สร้างข้อมูลเริ่มต้นสำหรับ Course, Period, CoursePeriod
func SetupData() {
	// เช็คว่ามี data อยู่แล้วหรือไม่
	var count int64
	DB.Model(&entity.Course{}).Count(&count)
	if count > 0 {
		fmt.Println("⚠️ Data already exists, skipping setup")
		return
	}

	// 1️⃣ สร้าง Courses
	courses := []entity.Course{
		{Name: "Foot Reflexology"},
		{Name: "Thai massage (Deep)"},
		{Name: "Deep Tissue massage"},
		{Name: "Sport Boxing oil massage"},
		{Name: "Aromatherapy oil massage"},
		{Name: "Coconut oil massage"},
		{Name: "Remedial massage"},
		{Name: "Remedial with Health"},
		{Name: "Pregnancy massage"},
		{Name: "Body scrub"},

		{Name: "Relaxation massage"},
		{Name: "Thai massage (Soft)"},

		{Name: "Hot stone massage"},

		{Name: "Cupping Therapy Vaccum"},
		{Name: "Remedial with Cupping"},

		{Name: "Neck, Should"},

		{Name: "Remedial massage with Coconut oil"},
		{Name: "Deep Tissue massage with Coconut oil"},
	}

	for _, c := range courses {
		DB.Create(&c)
	}

	// 2️⃣ สร้าง Periods
	periods := []entity.Period{
		{Duration: 0},   //1
		{Duration: 10},  //2
		{Duration: 15},  //3
		{Duration: 20},  //4
		{Duration: 30},  //5
		{Duration: 45},  //6
		{Duration: 60},  //7
		{Duration: 75},  //8
		{Duration: 90},  //9
		{Duration: 120}, //10
	}

	for _, p := range periods {
		DB.Create(&p)
	}

	// 3️⃣ สร้าง CoursePeriod (ราคาต่างกัน)
	coursePeriods := []entity.CoursePeriod{
		{CourseID: 1, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 1, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 1, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 1, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 1, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 2, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 2, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 2, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 2, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 2, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 3, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 3, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 3, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 3, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 3, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 4, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 4, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 4, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 4, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 4, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 5, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 5, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 5, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 5, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 5, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 6, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 6, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 6, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 6, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 6, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 7, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 7, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 7, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 7, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 7, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 8, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 8, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 8, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 8, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 8, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 9, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 9, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 9, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 9, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 9, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 10, PeriodID: 5, CusPrice: 70, EmpPrice: 30},   //30
		{CourseID: 10, PeriodID: 6, CusPrice: 95, EmpPrice: 39},   //45
		{CourseID: 10, PeriodID: 7, CusPrice: 115, EmpPrice: 50},  //60
		{CourseID: 10, PeriodID: 9, CusPrice: 165, EmpPrice: 75},  //90
		{CourseID: 10, PeriodID: 10, CusPrice: 230, EmpPrice: 95}, //120

		{CourseID: 11, PeriodID: 5, CusPrice: 65, EmpPrice: 29},   //30
		{CourseID: 11, PeriodID: 6, CusPrice: 85, EmpPrice: 39},   //45
		{CourseID: 11, PeriodID: 7, CusPrice: 105, EmpPrice: 46},  //60
		{CourseID: 11, PeriodID: 9, CusPrice: 155, EmpPrice: 71},  //90
		{CourseID: 11, PeriodID: 10, CusPrice: 210, EmpPrice: 92}, //120

		{CourseID: 12, PeriodID: 5, CusPrice: 65, EmpPrice: 29},   //30
		{CourseID: 12, PeriodID: 6, CusPrice: 85, EmpPrice: 39},   //45
		{CourseID: 12, PeriodID: 7, CusPrice: 105, EmpPrice: 46},  //60
		{CourseID: 12, PeriodID: 9, CusPrice: 155, EmpPrice: 71},  //90
		{CourseID: 12, PeriodID: 10, CusPrice: 210, EmpPrice: 92}, //120

		{CourseID: 13, PeriodID: 7, CusPrice: 120, EmpPrice: 50},   //60
		{CourseID: 13, PeriodID: 9, CusPrice: 170, EmpPrice: 75},   //90
		{CourseID: 13, PeriodID: 10, CusPrice: 235, EmpPrice: 100}, //120

		{CourseID: 14, PeriodID: 4, CusPrice: 49, EmpPrice: 22},  //20
		{CourseID: 14, PeriodID: 8, CusPrice: 139, EmpPrice: 60}, //75

		{CourseID: 15, PeriodID: 2, CusPrice: 20, EmpPrice: 10},   //10
		{CourseID: 15, PeriodID: 3, CusPrice: 29, EmpPrice: 14.5}, //15
		{CourseID: 15, PeriodID: 4, CusPrice: 39, EmpPrice: 19.5}, //20

		{CourseID: 16, PeriodID: 5, CusPrice: 75, EmpPrice: 30},   //30
		{CourseID: 16, PeriodID: 6, CusPrice: 100, EmpPrice: 39},  //45
		{CourseID: 16, PeriodID: 7, CusPrice: 120, EmpPrice: 50},  //60
		{CourseID: 16, PeriodID: 9, CusPrice: 170, EmpPrice: 75},  //90
		{CourseID: 16, PeriodID: 10, CusPrice: 235, EmpPrice: 95}, //120

		{CourseID: 17, PeriodID: 5, CusPrice: 75, EmpPrice: 30},   //30
		{CourseID: 17, PeriodID: 6, CusPrice: 100, EmpPrice: 39},  //45
		{CourseID: 17, PeriodID: 7, CusPrice: 120, EmpPrice: 50},  //60
		{CourseID: 17, PeriodID: 9, CusPrice: 170, EmpPrice: 75},  //90
		{CourseID: 17, PeriodID: 10, CusPrice: 235, EmpPrice: 95}, //120
	}

	for _, cp := range coursePeriods {
		DB.Create(&cp)
	}

	fmt.Println("✅ SetupData completed successfully")
}

func SeedData() {
	// Employees
	employees := []entity.Employee{
		{Name: "Nina", Revenue: 0},
		{Name: "Nat", Revenue: 0},
		{Name: "Yoko", Revenue: 0},
		{Name: "Rin", Revenue: 0},
		{Name: "Jenny", Revenue: 0},
		{Name: "Nile", Revenue: 0},
		{Name: "Rattie", Revenue: 0},
	}
	for _, e := range employees {
		DB.FirstOrCreate(&e, entity.Employee{Name: e.Name})
	}

	// Paymethods
	paymethods := []entity.Paymethod{
		{Name: "Cash"},
		{Name: "Credit Card"},
		{Name: "PromptPay"},
	}
	for _, p := range paymethods {
		DB.FirstOrCreate(&p, entity.Paymethod{Name: p.Name})
	}

	// HealthCares
	healthcares := []entity.HealthCare{
		{Name: "Bupa", Hicaps: 40},
		{Name: "NIB", Hicaps: 49},
		{Name: "GU Health", Hicaps: 50},
		{Name: "AHM", Hicaps: 30},
		{Name: "Australian unity", Hicaps: 35},
		{Name: "AAMI", Hicaps: 40},
		{Name: "ARHG", Hicaps: 25},
		{Name: "HCF", Hicaps: 10},
		{Name: "HBF", Hicaps: 40},
		{Name: "Doctors Health", Hicaps: 30},
	}
	for _, h := range healthcares {
		DB.FirstOrCreate(&h, entity.HealthCare{Name: h.Name})
	}

	fmt.Println("✅ SeedData completed successfully")
}

func SetHoliday() {
	publicHolidays := []entity.PublicHoliday{
		{
			HolidayDate: time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), // ตัวอย่างวัน Christmas
			Name:        "Christmas Day",
			State:       "NSW",
		},
		{
			HolidayDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), // New Year
			Name:        "New Year's Day",
			State:       "NSW",
		},
	}

	for _, h := range publicHolidays {
		DB.FirstOrCreate(&h, entity.PublicHoliday{HolidayDate: h.HolidayDate, State: h.State})
	}

	fmt.Println("✅ PublicHoliday seed completed successfully")
}
