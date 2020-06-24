package warehouse

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/craigericlewis/warehouse/server/api"
	"google.golang.org/api/iterator"
)

type Firestore interface {
	Close() error
	GetPackageInfo(ctx context.Context, packageName, packageVersion string, language api.Language) (*PackageInfo, error)
	AddPackageInfo(ctx context.Context, packageName, packageVersion, packageSize string, language api.Language) error
	getCollectionName(language api.Language) string
}

type firestoreRepository struct {
	client *firestore.Client
}

type PackageInfo struct {
	Name    string `firestore:"name,omitempty"`
	Version string `firestore:"version,omitempty"`
	Size    string `firestore:"size,omitempty"`
}

func NewFirestoreRepository(ctx context.Context, projectID string) (Firestore, error) {
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		return nil, err
	}

	return firestoreRepository{client}, nil
}

func (f *firestoreRepository) Close() error {
	return f.client.Close()
}

func (f *firestoreRepository) GetPackageInfo(ctx context.Context, packageName, packageVersion string, language api.Language) (*PackageInfo, error) {
	collectionName := f.getCollectionName(language)
	iter := f.client.Collection(collectionName).Where("name", "==", packageName).Where("version", "==", packageVersion).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

	}
}

func (f *firestoreRepository) AddPackageInfo(ctx context.Context, packageName, packageVersion, packageSize string, language api.Language) error {
	collectionName := f.getCollectionName(language)
	_, _, err := f.client.Collection(collectionName).Add(ctx, PackageInfo{
		Name:    packageName,
		Version: packageVersion,
		Size:    packageSize,
	})
	return err
}

func (f *firestoreRepository) getCollectionName(language api.Language) string {
	if language == api.Language_PYTHON3 {
		return "python3-packages"
	}
	return "python3-packages"
}
