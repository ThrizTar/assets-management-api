package repository

import (
	"assets-management-api/models"
	"context"

	firebase "firebase.google.com/go/v4"
)

func SaveStocksToFirestore(app *firebase.App, stocks []models.Stock) error {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	bw := client.BulkWriter(ctx)
	for _, s := range stocks {
		docRef := client.Collection("stocks").Doc(s.Symbol)
		_, err := bw.Set(docRef, s)
		if err != nil {
			return err
		}
	}

	bw.Flush()
	return nil
}
