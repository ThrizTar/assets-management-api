package repository

import (
	"context"
	"asset-management/models"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

func SaveStocksToFirestore(app *firebase.App, stocks []models.Stock) error {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// ใช้ Batch เพื่อบันทึกข้อมูลหลายตัวพร้อมกัน (ประหยัดเวลาและ Resource)
	batch := client.Batch()
	for _, s := range stocks {
		docRef := client.Collection("assets").Doc(s.ID)
		batch.Set(docRef, s)
	}

	_, err = batch.Commit(ctx)
	return err
}