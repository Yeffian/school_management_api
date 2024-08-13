package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Yeffian/school_management_api/models"
	"github.com/Yeffian/school_management_api/models/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	adit  = models.NewStudent("Adit", "Charakbroty", "2056@someschool.edu")
	john  = models.NewStudent("John", "Doe", "2187@someschool.edu")
	sarah = models.NewStudent("Sarah", "Smith", "1780@someschool.edu")
	mike  = models.NewStudent("Mike", "Doe", "8176@someschool.edu")
	kevin = models.NewStudent("Kevin", "Mann", "9094@someschool.edu")
)

var (
	josh = models.NewTeacher("Josh", "Brooks", "1298@teachers.someschool.edu", "physics")
	jude = models.NewTeacher("Jude", "Smith", "1298@teachers.someschool.edu", "english")
	mary = models.NewTeacher("Mary", "Afton", "1298@teachers.someschool.edu", "math")
)

var (
	physicsClass = models.NewClass("phys101")
	mathClass    = models.NewClass("math101")
	englishClass = models.NewClass("eng201")
)

func main() {
	physicsClass.AddStudent(*adit)
	physicsClass.AddStudent(*john)
	physicsClass.AddStudent(*mike)

	mathClass.AddStudent(*john)
	mathClass.AddStudent(*mike)
	mathClass.AddStudent(*kevin)

	englishClass.AddStudent(*sarah)
	englishClass.AddStudent(*mike)
	englishClass.AddStudent(*adit)

	classes := [3]models.Class{*physicsClass, *mathClass, *englishClass}
	// students := [5]models.Student{*adit, *john, *sarah, *mike, *kevin}
	teachers := [3]models.Teacher{*josh, *jude, *mary}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	studentsDb := sqlite.StudentModel{
		DB: db,
	}

	PORT := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/api/classes/:code", func(c *fiber.Ctx) error {
		classCode := c.Params("code")

		for i := 0; i < 3; i++ {
			if classes[i].ClassCode == classCode {
				return c.Status(200).JSON(classes[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "No class found with that code."})
	})

	app.Get("/api/students", func(c *fiber.Ctx) error {
		rows, err := studentsDb.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "Error."})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/students/firstName/:firstName", func(c *fiber.Ctx) error {
		firstName := c.Params("firstName")
		student, err := studentsDb.FromFirstName(firstName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(student)
	})

	app.Get("/api/students/lastName/:lastName", func(c *fiber.Ctx) error {
		lastName := c.Params("lastName")
		student, err := studentsDb.FromLastName(lastName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(student)
	})

	app.Get("/api/teachers/firstName/:firstName", func(c *fiber.Ctx) error {
		firstName := c.Params("firstName")

		for i := 0; i < 5; i++ {
			if teachers[i].FirstName == firstName {
				return c.Status(200).JSON(teachers[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
	})

	app.Get("/api/teachers/lastName/:lastName", func(c *fiber.Ctx) error {
		lastName := c.Params("lastName")

		for i := 0; i < 5; i++ {
			if teachers[i].LastName == lastName {
				return c.Status(200).JSON(teachers[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
	})

	app.Get("/api/teachers/:subject", func(c *fiber.Ctx) error {
		subject := c.Params("subject")
		var result []models.Teacher

		for i := 0; i < 3; i++ {
			if teachers[i].Subject == subject {
				result = append(result, teachers[i])
			}
		}

		return c.Status(200).JSON(result)
	})

	log.Fatal(app.Listen(":" + PORT))
}
