package repository

import (
	"context"
	"time"

	"backend/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type movieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) domain.MovieRepository {
	return &movieRepository{
		collection: db.Collection("movies"),
	}
}

func (r *movieRepository) GetByID(ctx context.Context, id string) (*domain.Movie, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var movie domain.Movie
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *movieRepository) GetByTMDBID(ctx context.Context, tmdbID int) (*domain.Movie, error) {
	var movie domain.Movie
	err := r.collection.FindOne(ctx, bson.M{"tmdbMovieId": tmdbID}).Decode(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *movieRepository) Create(ctx context.Context, movie *domain.Movie) error {
	movie.ID = primitive.NewObjectID()
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

func (r *movieRepository) Update(ctx context.Context, movie *domain.Movie) error {
	movie.UpdatedAt = time.Now()

	filter := bson.M{"_id": movie.ID}
	update := bson.M{"$set": movie}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *movieRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (r *movieRepository) List(ctx context.Context, limit, offset int) ([]*domain.Movie, int64, error) {
	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []*domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (r *movieRepository) Search(ctx context.Context, query string, limit, offset int) ([]*domain.Movie, int64, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": query, "$options": "i"}},
			{"overview": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []*domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (r *movieRepository) GetByGenre(ctx context.Context, genreID int, limit, offset int) ([]*domain.Movie, int64, error) {
	filter := bson.M{"genres.id": genreID}

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []*domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (r *movieRepository) GetPopular(ctx context.Context, limit, offset int) ([]*domain.Movie, int64, error) {
	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(offset)).
		SetSort(bson.M{"voteAverage": -1, "voteCount": -1})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var movies []*domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}
