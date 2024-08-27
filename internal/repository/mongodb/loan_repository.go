package mongodb

import (
	"context"
	"time"

	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewLoanRepository(collection *mongo.Collection) *LoanRepositoryMongo {
	return &LoanRepositoryMongo{
		Collection: collection,
	}
}

func (r *LoanRepositoryMongo) CreateLoan(loan domain.Loan) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, loan)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *LoanRepositoryMongo) FindLoanByID(id string) (*domain.Loan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var loan domain.Loan
	err := r.Collection.FindOne(ctx, filter).Decode(&loan)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func (r *LoanRepositoryMongo) FindAllLoans(status string, order string) ([]domain.Loan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}

	opts := options.Find()
	if order != "" {
		opts.SetSort(bson.M{"createdAt": order})
	}

	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var loans []domain.Loan
	for cursor.Next(ctx) {
		var loan domain.Loan
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil
}

func (r *LoanRepositoryMongo) UpdateLoanStatus(id string, status string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *LoanRepositoryMongo) DeleteLoan(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *LoanRepositoryMongo) ViewSystemLogs() ([]domain.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []domain.Log
	for cursor.Next(ctx) {
		var log domain.Log
		err := cursor.Decode(&log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}
func (r *LoanRepositoryMongo) FindAllLogs() ([]domain.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []domain.Log
	for cursor.Next(ctx) {
		var log domain.Log
		err := cursor.Decode(&log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}