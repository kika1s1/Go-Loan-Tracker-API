package mongodb

import (
	"context"
	"time"

	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoanRepository struct {
	collection *mongo.Collection
}

func NewLoanRepository(database *mongo.Database) *LoanRepository {
	collection := database.Collection("loans")
	return &LoanRepository{
		collection: collection,
	}
}

func (r *LoanRepository) CreateLoan(loan domain.Loan) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, loan)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *LoanRepository) FindLoanByID(id string) (*domain.Loan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var loan domain.Loan
	err := r.collection.FindOne(ctx, filter).Decode(&loan)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func (r *LoanRepository) FindAllLoans(status string, order string) ([]domain.Loan, error) {
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

	cursor, err := r.collection.Find(ctx, filter, opts)
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

func (r *LoanRepository) UpdateLoanStatus(id string, status string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *LoanRepository) DeleteLoan(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}