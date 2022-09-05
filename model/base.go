package model

import (
	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/satori/uuid"
	"time"
)

const (
	errAlreadyInitialized = "this entity is already initialized"
	errInvalidGuid        = "a valid uuid(v4) is required"
	errFutureCreated      = "DateCreatedUTCx is in the future"
	errFutureUpdated      = "DateUpdatedUTCx is in the future"
	errBeforeCreated      = "DateUpdatedUTCx before DateCreatedUTCx"
	errInvalidBool        = "must be true  or false"
)

//base is embedded into all entities to provide common meta data
type base struct {
	IDx             string    `json:"id,omitempty"`               //IDx uniquely identifies a domain entity
	DateCreatedUTCx time.Time `json:"date_created_utc,omitempty"` //DateCreatedUTCx is the date the base was created in UTC time
	DateUpdatedUTCx time.Time `json:"date_updated_utc,omitempty"` //DateUpdatedUTCx is the date the base was last updated in UTC time
	Deletedx        bool      `json:"deleted,omitempty"`          //Deletedx is if the base has been marked as Deletedx. Items marked for deletion should only be used for internal purposes.
	//tracerID        string    `json:"tracer_id,omitempty"`        //tracerID provided a tracer IDx for distributed event correlation
}

// ID gets the IDx
func (b *base) ID() string { return b.IDx }

// DateCreatedUTC gets the date the base was created in UTC format
func (b *base) DateCreatedUTC() time.Time { return b.DateCreatedUTCx }

// DateUpdatedUTC gets the date the base was last updated in UTC format
func (b *base) DateUpdatedUTC() time.Time { return b.DateUpdatedUTCx }

// Deleted gets the Deletedx status of the base
func (b *base) Deleted() bool { return b.Deletedx }

// ETag gets the etag value of the base
//func (b *base) TracerID() string { return b.tracerID }

// LoadBase is used when this struct is already loaded as an embedded struct into an existing model
func (b *base) LoadBase(id string, dateCreatedUTC time.Time, dateUpdatedUTC time.Time, deleted bool, tracerID string) *shared.CustomError {

	if uuid.FromStringOrNil(id) == uuid.Nil {
		return shared.NewCustomError(errInvalidGuid, shared.ErrorTypeSystem)
	}

	//if uuid.FromStringOrNil(tracerID) == uuid.Nil {
	//	return shared.NewCustomError(errInvalidGuid, shared.ErrorTypeSystem)
	//}

	if dateCreatedUTC.After(time.Now().UTC()) {
		return shared.NewCustomError(errFutureCreated, shared.ErrorTypeSystem)
	}

	if dateUpdatedUTC.After(time.Now().UTC()) {
		return shared.NewCustomError(errFutureUpdated, shared.ErrorTypeSystem)
	}

	if dateUpdatedUTC.Before(dateCreatedUTC) {
		return shared.NewCustomError(errBeforeCreated, shared.ErrorTypeSystem)
	}

	if deleted != true && deleted != false {
		return shared.NewCustomError(errInvalidBool, shared.ErrorTypeSystem)
	}

	b.IDx = id
	b.DateCreatedUTCx = dateCreatedUTC
	b.DateUpdatedUTCx = dateUpdatedUTC
	b.Deletedx = deleted
	//b.tracerID = tracerID
	return nil
}

// LoadExistingBase returns an existing base model
func LoadExistingBase(id string, dateCreatedUTC time.Time, dateUpdatedUTC time.Time, deleted bool) (*base, *shared.CustomError) {
	if uuid.FromStringOrNil(id) == uuid.Nil {
		return nil, shared.NewCustomError(errInvalidGuid, shared.ErrorTypeSystem)
	}

	if dateCreatedUTC.After(time.Now().UTC()) {
		return nil, shared.NewCustomError(errFutureCreated, shared.ErrorTypeSystem)
	}

	if dateUpdatedUTC.After(time.Now().UTC()) {
		return nil, shared.NewCustomError(errFutureUpdated, shared.ErrorTypeSystem)
	}

	if dateUpdatedUTC.Before(dateCreatedUTC) {
		return nil, shared.NewCustomError(errBeforeCreated, shared.ErrorTypeSystem)
	}

	if deleted != true && deleted != false {
		return nil, shared.NewCustomError(errInvalidBool, shared.ErrorTypeSystem)
	}
	return &base{
		IDx:             id,
		DateCreatedUTCx: dateCreatedUTC,
		DateUpdatedUTCx: dateUpdatedUTC,
		Deletedx:        deleted,
		//tracerID:        eTag,
	}, nil
}

//typeAssert will return nil if the value does not exist otherwise when  type casting it will blow up
func checkNil(v interface{}) interface{} {
	//todo: look at this janky code for a better way to do this
	switch v.(type) {
	case nil:
		return nil // does not exist or if nil already then passthrough
	default:
		return v // it has a value
	}
}

// newBase returns a new base model
func newBase() *base {
	return &base{shared.NewUUID(), time.Now().UTC(), time.Now().UTC(), false /*, shared.NewUUID()*/}
}
