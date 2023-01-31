package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://knightofhonour:Inlussion123@ac-adysxti-shard-00-00.duplhsn.mongodb.net:27017,ac-adysxti-shard-00-01.duplhsn.mongodb.net:27017,ac-adysxti-shard-00-02.duplhsn.mongodb.net:27017/?ssl=true&replicaSet=atlas-wt2oe1-shard-0&authSource=admin&retryWrites=true&w=majority"
const dbName = "BOKapp"

type classic_entry struct {
	ID   int64  `bson:"id,omitempty"`
	Text string `bson:"text,omitempty"`
	Tag  string `bson:"tag,omitempty"`
}

type classic_category struct {
	TypeOfCategory string `bson:"TypeOfTag,omitempty"`
	Category       string `bson:"Category,omitempty"`
	Entries        []int  `bson:"entries"`
}

func getAllEntriesWithCategory(client *mongo.Client, category string) []classic_entry {
	allEntryIDWithCategory := getAllEntryIDWithCategory(client, category)
	var AllEntriesWithCategory []classic_entry
	for _, v := range allEntryIDWithCategory {
		entry := getEntryByEntryID(client, v)
		AllEntriesWithCategory = append(AllEntriesWithCategory, entry)
	}
	return AllEntriesWithCategory
}

func getAllEntryIDWithCategory(client *mongo.Client, category string) []int {
	collection := "category"
	var result classic_category
	var search_criteria = primitive.E{Key: "Category", Value: category}
	record := readFromMongoDB(client, search_criteria, collection)
	err := record.Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.Entries
}

func getRandomEntry(client *mongo.Client) classic_entry {
	coll := client.Database(dbName).Collection("entry")
	filter := bson.D{}
	size, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(int(size)) + 1
	return getEntryByEntryID(client, id)
}

func getEntryByEntryID(client *mongo.Client, id int) classic_entry {
	collection := "entry"
	var result classic_entry
	var search_criteria = primitive.E{Key: "id", Value: id}
	record := readFromMongoDB(client, search_criteria, collection)
	err := record.Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func readFromMongoDB(client *mongo.Client, search_criteria primitive.E, collection string) *mongo.SingleResult {
	coll := client.Database(dbName).Collection(collection)
	err := coll.FindOne(context.TODO(), bson.D{search_criteria})
	return err
}

func connect() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	return client
}

func main() {
	client := connect()
	entry := getEntryByEntryID(client, 1)
	fmt.Println(entry)
	entry = getRandomEntry(client)
	fmt.Println(entry)
	tagEntries := getAllEntriesWithCategory(client, "test1")
	fmt.Println(tagEntries)
}
