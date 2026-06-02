package model

import "go.mongodb.org/mongo-driver/v2/bson" 

type Netfilx struct {
	ID bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string `json:"movie"`
	IsWatched bool `json:"isWatched"`
	CreatedAt bson.DateTime `bson:"created_at"`
}

func (n *Netfilx) IsEmpty() bool{
	
	return n.Movie == "" 
}