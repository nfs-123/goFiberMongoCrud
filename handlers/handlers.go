package handlers

import (
	"context"
	"fmt"
	"goFiberMongoCrud/Config"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type College struct {
	Name string `bson:"collegeName,omitempty" json:"collegeName"`
	City string `json:"city,omitempty" bson:",omitempty"`
}

func GetColleges(c *fiber.Ctx) error {
	client, _ := db.GetMongoDbConnection()
	collection := client.Database("local").Collection("College")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("error in finding documengts", err)
	}
	var colleges []College
	if err := cursor.All(context.Background(), &colleges); err != nil {
		log.Fatal("error in finding documents", err)
	}

	return c.JSON(fiber.Map{
		"success": "true",
		"data":    colleges,
	})

}
func AddCollege(c *fiber.Ctx) error {
	college := College{}
	if err := c.BodyParser(&college); err != nil {
		log.Fatal("error in posting", err)
	}
	fmt.Println("log",college.Name)
	client,_ := db.GetMongoDbConnection()
	collection := client.Database("local").Collection("College")
	resultId,err := collection.InsertOne(context.Background(),college)
	if err != nil {
		log.Fatal("Issue in Inserting",err)
	}

	
	return c.JSON(fiber.Map{
		"success": "true",
		"id":resultId,
		"data":college,
		
	})
}
func DeleteCollege(c *fiber.Ctx) error{
	city := c.Params("name")
	fmt.Println(city)
	client,_ := db.GetMongoDbConnection()
	collection := client.Database("local").Collection("College")
	resultId,err := collection.DeleteOne(context.Background(),bson.M{"City":city})
	if err!=nil {
		fmt.Println("issue in deleting",err)
	}
	return c.JSON(fiber.Map{
		"result":"true",
		"id":resultId,
	})




}
