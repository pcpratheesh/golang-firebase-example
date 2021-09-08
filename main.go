package main

import (
	"context"
	"flag"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var (
	client *db.Client
	ctx    context.Context

	add, update, delete *bool
	firebaseDBUrl       *string
)

type Products struct {
	Name  string  `json:"product_name"`
	Price float32 `json:"product_value"`
}

func init() {
	add = flag.Bool("add", false, "add new data to database")
	update = flag.Bool("update", false, "update entry in database")
	delete = flag.Bool("delete", false, "delete entry from database")

	firebaseDBUrl = flag.String("db-url", "db url", "firebase db url")

	flag.Parse()

	ctx = context.Background()

	conf := &firebase.Config{
		DatabaseURL: *firebaseDBUrl,
	}
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

}

func main() {
	// Connect to the structure
	ref := client.NewRef("products")

	// add new entry to database
	switch {
	case *add:
		err := ref.Set(ctx, map[string]Products{
			"camera": Products{
				Name:  "Canon 5D",
				Price: 1024,
			},
			"Nikon D5600": Products{
				Name:  "Nikon D5600",
				Price: 65500,
			},
			"Nikon D7500": Products{
				Name:  "Nikon D5600",
				Price: 84000,
			},
		})

		if err != nil {
			log.Fatalln("unable to insert products:", err)
			return
		}

		log.Printf("successfully added products")
	case *update:
		err := ref.Child("camera").Set(ctx, Products{
			Name:  "Canon 5D",
			Price: 38000,
		})
		if err != nil {
			log.Fatalln("unable to update product:", err)
			return
		}
		log.Printf("successfully updated product")

	case *delete:
		err := ref.Child("camera").Delete(ctx)
		if err != nil {
			log.Fatalln("unable to delete product:", err)
			return
		}
		log.Printf("successfully deleted product")
	default:
		log.Println("Please provide any of input add,update,delete")
	}
}
