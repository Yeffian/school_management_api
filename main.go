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

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	dbProvider := sqlite.CreateDatabaseProvider(db)

	PORT := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/api/classes", func(c *fiber.Ctx) error {
		rows, err := dbProvider.Class.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Post("/api/classes/new", func(c *fiber.Ctx) error {
		cl := &models.Class{}

		if err := c.BodyParser(cl); err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
		}

		cl, err = dbProvider.Class.New(cl.ClassCode, cl.Subject)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(cl)
	})

	app.Get("/api/students", func(c *fiber.Ctx) error {
		rows, err := dbProvider.Student.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Post("/api/students/new", func(c *fiber.Ctx) error {
		s := &models.Student{}

		if err := c.BodyParser(s); err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
		}

		err = dbProvider.Student.New(*s)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(s)
	})

	app.Get("/api/students/firstName/:firstName", func(c *fiber.Ctx) error {
		firstName := c.Params("firstName")
		student, err := dbProvider.Student.ByFirstName(firstName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(student)
	})

	app.Get("/api/students/classes", func(c *fiber.Ctx) error {
		classCode := c.Query("classCode")

		rows, err := dbProvider.Class.StudentsByClass(classCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/students/lastName/:lastName", func(c *fiber.Ctx) error {
		lastName := c.Params("lastName")
		student, err := dbProvider.Student.ByLastName(lastName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(student)
	})

	app.Get("/api/teachers", func(c *fiber.Ctx) error {
		rows, err := dbProvider.Teacher.All()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Post("/api/teachers/new", func(c *fiber.Ctx) error {
		t := &models.Teacher{}

		if err := c.BodyParser(t); err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
		}

		err = dbProvider.Teacher.New(*t)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(t)
	})

	app.Get("/api/teachers/firstName/:firstName", func(c *fiber.Ctx) error {
		firstName := c.Params("firstName")
		teacher, err := dbProvider.Teacher.ByFirstName(firstName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(teacher)
	})

	app.Get("/api/teachers/lastName/:lastName", func(c *fiber.Ctx) error {
		lastName := c.Params("lastName")
		teacher, err := dbProvider.Teacher.ByLastName(lastName)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": "No student found with that name."})
		}

		return c.Status(200).JSON(teacher)
	})

	app.Get("/api/teachers/subjects/:subject", func(c *fiber.Ctx) error {
		subject := c.Params("subject")
		rows, err := dbProvider.Teacher.BySubject(subject)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	app.Get("/api/teachers/classes", func(c *fiber.Ctx) error {
		classCode := c.Query("classCode")

		rows, err := dbProvider.Class.TeachersByClass(classCode)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
		}

		return c.Status(200).JSON(rows)
	})

	log.Fatal(app.Listen(":" + PORT))
}
