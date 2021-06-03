package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type Recurser map[string]interface{} // TODO: actual struct with fields

type RecurserDB interface {
	GetByUserID(ctx context.Context, userID string) (Recurser, error)
	Set(ctx context.Context, userID string, recurser Recurser) error
	Delete(ctx context.Context, userID string) error
	ListPairingTomorrow(ctx context.Context) ([]Recurser, error)
	ListSkippingTomorrow(ctx context.Context) ([]Recurser, error)
	SetSkippingTomorrow(ctx context.Context, userID string) error
}

type FirestoreRecurserDB struct {
	client *firestore.Client
}

func (f *FirestoreRecurserDB) GetByUserID(ctx context.Context, userID string) (Recurser, error) {
	return Recurser{}, nil
}

func (f *FirestoreRecurserDB) Set(ctx context.Context, userID string, recurser Recurser) error {
	return nil
}

func (f *FirestoreRecurserDB) Delete(ctx context.Context, userID string) error {
	return nil
}

func (f *FirestoreRecurserDB) ListPairingTomorrow(ctx context.Context) ([]Recurser, error) {
	return nil, nil
}

func (f *FirestoreRecurserDB) ListSkippingTomorrow(ctx context.Context) ([]Recurser, error) {
	return nil, nil
}

func (f *FirestoreRecurserDB) SetSkippingTomorrow(ctx context.Context, userID string) error {
	return nil
}

type APIAuthDB interface {
	GetAPIAuthKey(ctx context.Context) (string, error)
}

type FirestoreAPIAuthDB struct {
	client *firestore.Client
}

func (f *FirestoreAPIAuthDB) GetAPIAuthKey(ctx context.Context) (string, error) {
	doc, err := f.client.Collection("botauth").Doc("token").Get(ctx)
	if err != nil {
		log.Println("Something weird happened trying to read the auth token from the database")
		return "", err
	}

	token := doc.Data()
	return token["value"].(string), nil
}
