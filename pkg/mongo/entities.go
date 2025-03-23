package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Todo struct {
	ID          bson.ObjectID `bson:"_id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Completed   bool          `bson:"completed"`
	Priority    uint8         `bson:"priority_level"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

type Board struct {
	ID          bson.ObjectID `bson:"_id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Completed   bool          `bson:"completed"`
	Priority    uint8         `bson:"priority_level"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}
