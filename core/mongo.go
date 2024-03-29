package core

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

const dbName = "BOKapp"

type classic_entry struct {
	ID   int64  `bson:"id,omitempty"`
	Text string `bson:"text,omitempty"`
	Tag  string `bson:"tag,omitempty"`
}

type classic_category struct {
	TypeOfCategory string `bson:"TypeOfCategory,omitempty"`
	Category       string `bson:"Category,omitempty"`
	Entries        []int  `bson:"entries"`
}

type classic_category_list struct {
	Categories []string `bson:"categories"`
}

func getAllCategories(client *mongo.Client, categoryListCollection string) []string {
	var result classic_category_list
	var search_criteria = primitive.E{}
	record := readFromMongoDB(client, search_criteria, categoryListCollection)
	err := record.Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.Categories
}

func getAllEntriesWithCategory(client *mongo.Client, category string, categoryCollection string, entryCollection string) []classic_entry {
	allEntryIDWithCategory := getAllEntryIDWithCategory(client, category, categoryCollection)
	var AllEntriesWithCategory []classic_entry
	for _, v := range allEntryIDWithCategory {
		entry := getEntryByEntryID(client, v, entryCollection)
		AllEntriesWithCategory = append(AllEntriesWithCategory, entry)
	}
	return AllEntriesWithCategory
}

func getAllEntryIDWithCategory(client *mongo.Client, category string, categoryCollection string) []int {
	var result classic_category
	var search_criteria = primitive.E{Key: "Category", Value: category}
	record := readFromMongoDB(client, search_criteria, categoryCollection)
	err := record.Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.Entries
}

func getRandomEntry(client *mongo.Client, entryCollection string) classic_entry {
	coll := client.Database(dbName).Collection(entryCollection)
	filter := bson.D{}
	size, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(int(size)) + 1
	return getEntryByEntryID(client, id, entryCollection)
}

func getEntryByEntryID(client *mongo.Client, id int, entryCollection string) classic_entry {
	var result classic_entry
	var search_criteria = primitive.E{Key: "id", Value: id}
	record := readFromMongoDB(client, search_criteria, entryCollection)
	err := record.Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func insertEntry(client *mongo.Client, text string, entryCollection string) (bool, int) {
	coll := client.Database(dbName).Collection(entryCollection)
	filter := bson.D{}
	size, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	new_ID := size + 1
	entryToInsert := classic_entry{ID: new_ID, Text: text, Tag: "basic"}
	return insertIntoMongo(client, entryCollection, entryToInsert), int(new_ID)
}

func insertCategory(client *mongo.Client, category string, entryID int, categoryCollection string, categoryListCollection string) bool {
	categoryToInsert := classic_category{TypeOfCategory: "basic", Entries: []int{entryID}, Category: category}
	success := insertIntoMongo(client, categoryCollection, categoryToInsert)
	success2 := updateCategoryList(client, categoryListCollection, category)
	return success && success2
}

func updateCategoryList(client *mongo.Client, categoryListCollection string, category string) bool {
	coll := client.Database(dbName).Collection(categoryListCollection)
	allCategories := getAllCategories(client, categoryListCollection)
	allCategories = append(allCategories, category)
	categoryListUpdate := classic_category_list{Categories: allCategories}

	result, err := coll.ReplaceOne(context.TODO(), bson.D{}, categoryListUpdate)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if result.ModifiedCount != 1 {
		return false
	}
	return true
}

func updateCategory(client *mongo.Client, category string, entryID int, categoryCollection string) bool {
	coll := client.Database(dbName).Collection(categoryCollection)
	filter := primitive.E{Key: "Category", Value: category}
	allEntryIDWithCategory := getAllEntryIDWithCategory(client, category, categoryCollection)
	allEntryIDWithCategory = append(allEntryIDWithCategory, entryID)
	categoryToUpdate := classic_category{TypeOfCategory: "basic", Entries: allEntryIDWithCategory, Category: category}
	result, err := coll.ReplaceOne(context.TODO(), bson.D{filter}, categoryToUpdate)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if result.MatchedCount == 0 {
		return false
	}
	return true
}

func readFromMongoDB(client *mongo.Client, search_criteria primitive.E, collection string) *mongo.SingleResult {
	coll := client.Database(dbName).Collection(collection)
	result := coll.FindOne(context.TODO(), bson.D{search_criteria})
	return result
}

func insertIntoMongo(client *mongo.Client, collection string, toInsert interface{}) bool {
	coll := client.Database(dbName).Collection(collection)
	_, err := coll.InsertOne(context.TODO(), toInsert)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Connect() *mongo.Client {
	uri := "mongodb://knightofhonour:Inlussion123@ac-adysxti-shard-00-00.duplhsn.mongodb.net:27017,ac-adysxti-shard-00-01.duplhsn.mongodb.net:27017,ac-adysxti-shard-00-02.duplhsn.mongodb.net:27017/?ssl=true&replicaSet=atlas-wt2oe1-shard-0&authSource=admin&retryWrites=true&w=majority"
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

// func main() {
// 	client := connect()
// 	entry := getEntryByEntryID(client, 1)
// 	fmt.Println(entry)
// 	entry = getRandomEntry(client)
// 	fmt.Println(entry)
// 	tagEntries := getAllEntriesWithCategory(client, "test1")
// 	fmt.Println(tagEntries)
// }
