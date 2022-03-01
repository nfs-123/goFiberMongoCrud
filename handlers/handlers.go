package handlers

import (
	"context"
	"goFiberMongoCrud/Config"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)
type College struct {
	Name string `json:"collegeName" bson:"collegeName"`
	City string `json:"city" bson:"city"`
}


func GetColleges(c *fiber.Ctx) error {
	client,_ := db.GetMongoDbConnection()
	collection := client.Database("local").Collection("College")
	cursor,err := collection.Find(context.Background(),bson.M{})
	if err!=nil{
		log.Fatal("error in finding documengts",err)
	}
	var colleges []College
	if err := cursor.All(context.Background(),&colleges);err!=nil{
		log.Fatal("error in finding documents",err)
	}

	return c.JSON(fiber.Map{
		"success":"true",
		"data":colleges,
	})




}