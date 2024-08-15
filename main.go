package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Yeffian/school_management_api/models/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
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

	teachersDb := sqlite.TeacherModel{
		DB: db,
	}

	classDb := sqlite.ClassModel{
		DB: db,
	}

	PORT := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/api/classes", func(c *fiber.Ctx) error {
		rows, err := classDb.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/students", func(c *fiber.Ctx) error {
		rows, err := studentsDb.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
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

	app.Get("/api/students/classes", func(c *fiber.Ctx) error {
		classCode := c.Query("classCode")

		rows, err := classDb.StudentsByClass(classCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/students/lastName/:lastName", func(c *fiber.Ctx) error {
		lastName := c.Params("lastName")
		student, err := studentsDb.FromLastName(lastName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(student)
	})

	app.Get("/api/teachers", func(c *fiber.Ctx) error {
		rows, err := teachersDb.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/teachers/firstName/:firstName", func(c *fiber.Ctx) error {
		firstName := c.Params("firstName")
		teacher, err := teachersDb.FromFirstName(firstName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(teacher)
	})

	app.Get("/api/teachers/lastName/:lastName", func(c *fiber.Ctx) error {
		lastName := c.Params("lastName")
		teacher, err := teachersDb.FromLastName(lastName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(teacher)
	})

	app.Get("/api/teachers/subjects/:subject", func(c *fiber.Ctx) error {
		subject := c.Params("subject")
		rows, err := teachersDb.FromSubject(subject)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/teachers/classes", func(c *fiber.Ctx) error {
		classCode := c.Query("classCode")

		rows, err := classDb.TeachersByClass(classCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	log.Fatal(app.Listen(":" + PORT))
}
