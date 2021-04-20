package gerrors

import (
	errpb "github.com/srikrsna/errors/gen/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal         = status.Error(codes.Internal, "something seems wrong please try again after sometime")
	ErrNotFound         = status.Error(codes.NotFound, "the resource with the given id is not found")
	ErrFieldMask        = status.Error(codes.InvalidArgument, "field mask contains invalid root field paths")
	ErrInvalidPageToken = status.Error(codes.InvalidArgument, "pageToken is invalid")
)

var _errRevisionMismatch = status.New(codes.FailedPrecondition, "revision did not match the current revision")

func ErrRevisionMismatch(current, request int64) error {
	st, _ := _errRevisionMismatch.WithDetails(&errpb.RevisionMismatch{
		CurrrentRevision: current,
		RequestRevision:  request,
	})
	return st.Err()
}
