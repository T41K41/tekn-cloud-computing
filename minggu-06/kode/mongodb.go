package main

import (
  "context"
  "fmt"
  "log"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type test struct {
  NIM  string `bson:"nim"`
  Nama string `bson:"nama"`
}

func connect() (*mongo.Database, error) {
  clientOptions := options.Client()
  clientOptions.ApplyURI("mongodb://localhost:27017")
  client, err := mongo.NewClient(clientOptions)
  if err != nil {
    return nil, err
  }

  err = client.Connect(ctx)
  if err != nil {
    return nil, err
  }

  return client.Database("test"), nil
}

func find() {
  db, err := connect()
  if err != nil {
    log.Fatal(err.Error())
  }

  csr, err := db.Collection("test").Find(ctx, bson.M{"nama": "T41K41"})
  if err != nil {
    log.Fatal(err.Error())
  }
  defer csr.Close(ctx)

  result := make([]test, 0)
  for csr.Next(ctx) {
    var row test
    err := csr.Decode(&row)
    if err != nil {
      log.Fatal(err.Error())
    }

    result = append(result, row)
  }

  if len(result) > 0 {
    fmt.Println("NIM  :", result[0].NIM)
    fmt.Println("Nama :", result[0].Nama)
  }
}

func main() {
  find()
}