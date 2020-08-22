package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var recipesCollection *mongo.Collection

// RecipeItem recorded in the database
type RecipeItem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Content string             `bson:"content"`
	Title   string             `bson:"title"`
}

// GetOne recipe by id
func GetOne(recipeID primitive.ObjectID) (*RecipeItem, error) {
	data := &RecipeItem{}
	res := recipesCollection.FindOne(context.Background(), bson.M{"_id": recipeID})
	if err := res.Decode(data); err != nil {
		return nil, fmt.Errorf("Can't find recipe: %s", err)
	}

	return data, nil
}

// GetAll recipes
func GetAll() ([]*RecipeItem, error) {
	var resp []*RecipeItem
	cur, err := recipesCollection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("error in find")
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &RecipeItem{}
		if err := cur.Decode(data); err != nil {
			return nil, fmt.Errorf("Can't find recipe: %s", err)
		}
		resp = append(resp, data)
	}

	fmt.Printf("%v\n", resp)
	return resp, nil
}

// Create new recipe
func Create(data *RecipeItem) (*RecipeItem, error) {
	res, err := recipesCollection.InsertOne(context.Background(), bson.D{
		{Key: "title", Value: data.Title},
		{Key: "content", Value: data.Content},
	})
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("Cannot convert to OID")
	}

	data.ID = oid
	return data, nil
}

// Update recipe ite mwith input values
func Update(id primitive.ObjectID, data *RecipeItem, rewrite bool) (bson.M, error) {
	updateValues := MapUpdateData(*data, rewrite) // so apparently the *data converts pointer to value

	update := bson.D{{Key: "$set", Value: updateValues}}
	fmt.Println(update)
	resp := recipesCollection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	err := resp.Err()
	if err != nil {
		return nil, err
	}

	doc := bson.M{}
	decodeErr := resp.Decode(doc)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return doc, nil
}

// Delete recipe
func Delete(recipeID primitive.ObjectID) error {
	_, err := recipesCollection.DeleteOne(context.Background(), bson.M{"_id": recipeID})
	if err != nil {
		return errors.New("Cannot convert to OID")
	}
	return nil
}
