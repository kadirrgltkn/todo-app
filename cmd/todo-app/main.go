package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kadirrgltkn/todo-app/pkg/models/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	opts := options.Client().
		ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	mongoDBTodosCollection := client.Database("todo-app").Collection("todos")

	/** todoList := &todo.TodoList{
		Todos: []todo.Todo{{
			Completed: true,
			CreatedAt: time.Now(),
			Title:     "Şerafettin Hürtaşak",
			Gorev:     "Ekonomi Bakanı",
		}, {
			Completed: false,
			CreatedAt: time.Now(),
			Title:     "Damat Nizamettin Hürtaşak",
			Gorev:     "Damat, Ekonomi Bakanı Yardımcısı",
		},
			{
				Completed: true,
				CreatedAt: time.Now(),
				Title:     "Fatih Sultan Mehmed",
				Gorev:     "1453 - İstanbul'un Fethi",
			}},
	}

	for _, todo := range todoList.Todos {
		if _, err = mongoDBTodosCollection.InsertOne(context.TODO(), todo); err != nil {
			println("Failed to save To-Do. %s", err)
		}

		println("To-Do Saved to database.")
	}

	println("All todos saved to database successfully.")
	*/

	serafettininyarragi := http.NewServeMux()

	serafettininyarragi.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		var todos []*todo.Todo

		filter := bson.D{{Key: "completed", Value: true}}

		cursor, err := mongoDBTodosCollection.Find(context.TODO(), filter, nil)

		if err != nil {
			println("Failed to retrieve todos from database.")
		}

		for cursor.Next(context.TODO()) {
			var todo *todo.Todo

			if err = cursor.Decode(&todo); err != nil {
				println("Hata verdi ..... düzgün yaz.")
			}

			todos = append(todos, todo)
		}

		w.Header().Add("Content-Type", "application/json")

		if err = json.NewEncoder(w).Encode(todos); err != nil {
			http.Error(w, "As denied server.", http.StatusInternalServerError)
		}
	})

	http.ListenAndServe("127.0.0.1:3131", serafettininyarragi)

}
