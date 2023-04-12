package model

import (
 "go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
 Id primitive.ObjectID
 Title string
 Article string
}