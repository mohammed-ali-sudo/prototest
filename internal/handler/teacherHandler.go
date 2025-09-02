package handler

import (
	"context"
	"fmt"
	"reflect"

	"gapi/internal/model"
	gapi "gapi/proto/gen"

	"github.com/kamva/mgm/v3"
)

func (s *Server) AddTeacher(ctx context.Context, req *gapi.Teachers) (*gapi.Teachers, error) {
	// Convert gRPC teachers into slice of interface{}
	docs := make([]interface{}, len(req.Teachers))
	for i, t := range req.Teachers {
		docs[i] = &model.Teacher{
			FirstName: t.FirstName,
			LastName:  t.LastName,
			Email:     t.Email,
			Class:     t.Class,
			Subject:   t.Subject,
		}
		pbval := reflect.ValueOf(t).Elem()
		fmt.Println(pbval)
	}

	// Use InsertMany on the teachers collection
	coll := mgm.Coll(&model.Teacher{})
	_, err := coll.InsertMany(ctx, docs)
	if err != nil {
		_ = ErrorHandler(err, "❌ Failed to insert teachers")
		return nil, fmt.Errorf("failed to insert teachers: %w", err)
	}

	// req is already *gapi.Teachers
	return req, nil
}
