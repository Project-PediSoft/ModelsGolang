package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/99designs/gqlgen/graphql"
)

// GetHistogram is the resolver for the getHistogram field.
func (r *queryResolver) GetHistogram(ctx context.Context, file graphql.Upload) (string, error) {
	fileContent, err := ioutil.ReadAll(file.File)

	if err != nil {
		log.Printf("Error al leer el contenido del archivo DICOM: %v", err)
		return "error", err
	}
	histogram := make(map[string]int)
	for _, value := range fileContent {
		if value >= 0 && value <= 255 {
			key := fmt.Sprintf("%d", value)
			histogram[key]++
		}
	}
	histogramJSON, err := json.Marshal(histogram)
	if err != nil {
		log.Printf("Error al convertir el histograma a JSON: %v", err)
		return "error", err
	}

	return string(histogramJSON), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
