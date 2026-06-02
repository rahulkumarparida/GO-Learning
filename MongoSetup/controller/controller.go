package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongosetup/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


const connectionString string = "your-mongo-url"
const dbName string = "NetfilxLite"
const collectionName string = "watchlist"

var  collection *mongo.Collection

func init()  {
	clientOption := options.Client().ApplyURI(connectionString)

	client , err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo COnnnection successful")

	collection = client.Database(dbName).Collection(collectionName)
	
	fmt.Println("collection created sucesssfully")
}




// Mongo Helpers
func insertMovie(movie model.Netfilx) *mongo.InsertOneResult{

	inserted , err := collection.InsertOne(context.Background(),movie)

	if err != nil {
		log.Fatal(err)
	}

	return  inserted
}


func updateMovie(movieId string) *mongo.UpdateResult{
	id , err := bson.ObjectIDFromHex(movieId)
	
	if err != nil {
		log.Fatal(err)
	}

	// updated , err := collection.UpdateByID(context.Background(), id) 
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"isWatched":true}}
	
	updated , err := collection.UpdateOne(context.Background(), filter , update)
	
	fmt.Println("Movie is updated:", updated)	

	return  updated
}


func deleteMovie(movieId string) int64{
	id ,err := bson.ObjectIDFromHex(movieId)
	
	if err != nil {
		log.Fatal(err)
	}
	
	filter := bson.M{"_id":id}
	
	deleted , err := collection.DeleteOne(context.Background(), filter)
	
	fmt.Println("Deleted Id:", id," Successfully responded: ", deleted.DeletedCount)
	
	return  deleted.DeletedCount
}

func deleteMovies() int64{
	filter := bson.D{{}}
	
	deleted , err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	return deleted.DeletedCount
}



func  getAllMovies()  []bson.M{
	cur , err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cur.Next(context.Background()){
		var movieObj bson.M
		if err := cur.Decode(&movieObj); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movieObj)
		fmt.Println("Movie: ", movieObj)
	}

	defer cur.Close(context.Background())

	fmt.Println("movies list:", movies)
	return  movies
	
}





//  Mongo Controllers
func GetAllMovies(w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
	defer r.Body.Close()
	return
}

func CreateMovie(w http.ResponseWriter , r *http.Request)  {
			w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Allow-Control-AlloW-Methods", "POST")

		var movie model.Netfilx
		_ = json.NewDecoder(r.Body).Decode(&movie)

		createdmovie := insertMovie(movie)

		json.NewEncoder(w).Encode(createdmovie)
	
}


func MarkMovieWatched(w http.ResponseWriter , r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Allow-Control-AlloW-Methods", "PATCH")

		params := mux.Vars(r)

		update :=updateMovie(params["id"])

		json.NewEncoder(w).Encode(update)
}

func DeleteMovie(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Allow-Control-AlloW-Methods", "DELETE")

		params := mux.Vars(r)

		deletecount := deleteMovie(params["id"])


		json.NewEncoder(w).Encode(deletecount)

}


func DeleteMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-AlloW-Methods", "DELETE")

	deletecount := deleteMovies()

	json.NewEncoder(w).Encode(deletecount)
}