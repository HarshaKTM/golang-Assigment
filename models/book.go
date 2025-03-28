package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book represents the book entity
type Book struct {
	ID              primitive.ObjectID `json:"bookId,omitempty" bson:"_id,omitempty"`
	AuthorID        string             `json:"authorId" bson:"authorId"`
	PublisherID     string             `json:"publisherId" bson:"publisherId"`
	Title           string             `json:"title" bson:"title"`
	PublicationDate string             `json:"publicationDate" bson:"publicationDate"`
	ISBN            string             `json:"isbn" bson:"isbn"`
	Pages           int                `json:"pages" bson:"pages"`
	Genre           string             `json:"genre" bson:"genre"`
	Description     string             `json:"description" bson:"description"`
	Price           float64            `json:"price" bson:"price"`
	Quantity        int                `json:"quantity" bson:"quantity"`
}
