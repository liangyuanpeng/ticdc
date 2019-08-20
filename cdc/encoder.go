package cdc

import (
	"fmt"

	"github.com/pingcap/errors"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/tidb/types"
)

// encodeRow holds all the pieces necessary to encode a row change into a key or
// value.
type encodeRow struct {
	// datums is the new value of a changed table row.
	datums []types.Datum

	commitTS uint64

	// deleted is true if row is a deletion. In this case, only the primary
	// key columns are guaranteed to be set in `datums`.
	deleted bool

	tableInfo *model.TableInfo
}

type Encoder interface {
	// EncodeKey encodes the primary key of the given row.
	EncodeKey(encodeRow) ([]byte, error)

	// EncodeValue encodes the primary key of the given row.
	EncodeValue(encodeRow) ([]byte, error)

	// EncodeResolvedTimestamp encodes a resolved timestamp payload
	EncodeResolvedTimestamp(uint64) ([]byte, error)
}

func getEncoder(opts map[string]string) (Encoder, error) {
	switch formatType(opts[optFormat]) {
	case "", optFormatJSON:
		return newJSONEncoder(opts)
	default:
		return nil, errors.Errorf("unknow format: %s", opts[optFormat])
	}
}

type jsonEncoder struct {
}

var _ Encoder = &jsonEncoder{}

func newJSONEncoder(opts map[string]string) (*jsonEncoder, error) {
	// TODO
	return &jsonEncoder{}, nil
}

func (e *jsonEncoder) EncodeKey(row encodeRow) ([]byte, error) {
	// TODO
	str := fmt.Sprintf("%+v", row)
	return []byte(str), nil
}

func (e *jsonEncoder) EncodeValue(row encodeRow) ([]byte, error) {
	// TODO
	str := fmt.Sprintf("%+v", row)
	return []byte(str), nil
}

func (e *jsonEncoder) EncodeResolvedTimestamp(ts uint64) ([]byte, error) {
	// TODO
	str := fmt.Sprintf("ts: %v", ts)
	return []byte(str), nil
}
