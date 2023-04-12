package model

import (
 "go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
 Id primitive.ObjectID
 Title string
 Article string
 Detail string
 EventDateFrom string
 EventDateTo string
 EventOrganizer string
 EventLocation string
 EventImage string
 CreatedBy int
 CreatedDate string
 UpdatedDate string
}