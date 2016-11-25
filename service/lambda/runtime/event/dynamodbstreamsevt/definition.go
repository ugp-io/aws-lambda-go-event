//
// Copyright 2016 Alsanium, SAS. or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dynamodbstreamsevt

import (
	"encoding/json"
	"time"
)

// Timestamp represents an Amazon DynamoDB Streams event record epoch Timestamp
// with millisecond precision.
type Timestamp struct {
	time.Time
}

// AttributeValue represents the data for an attribute. One, and only one, of
// the elements is set.
//
// Each attribute in an item is a name-value pair. An attribute can be
// single-valued or multi-valued set. For example, a book item can have title
// and authors attributes. Each book has one title but can have many authors.
// The multi-valued attribute is a set; duplicate values are not allowed.
//
// See http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBMapper.DataTypes.html
type AttributeValue struct {
	// A Binary data type.
	//
	// B is automatically base64 encoded/decoded by the SDK.
	B []byte

	// A Boolean data type.
	BOOL bool

	// A Binary Set data type.
	BS [][]byte

	// A List of attribute values.
	L []AttributeValue

	// A Map of attribute values.
	M map[string]AttributeValue

	// A Number data type.
	N string

	// A Number Set data type.
	NS []string

	// A Null data type.
	NULL bool

	// A String data type.
	S string

	// A String Set data type.
	SS []string
}

// Record is a description of a single data modification that was performed on
// an item in a DynamoDB table.
type Record struct {
	// The unique identifier of the record in the stream.
	// See also http://docs.aws.amazon.com/streams/latest/dev/kinesis-record-processor-duplicates.html ;)
	SequenceNumber string

	// The type of data from the modified DynamoDB item that was captured in
	// this stream record:
	//   * KEYS_ONLY - only the key attributes of the modified item.
	//   * NEW_IMAGE - the entire item, as it appeared after it was modified.
	//   * OLD_IMAGE - the entire item, as it appeared before it was modified.
	//   * NEW_AND_OLD_IMAGES - both the new and the old item images of the item.
	StreamViewType string

	// The approximate date and time when the stream record was created.
	ApproximateCreationDateTime Timestamp

	// The primary key attribute(s) for the DynamoDB item that was modified.
	Keys map[string]*AttributeValue

	// The item in the DynamoDB table as it appeared before it was modified.
	OldImage map[string]*AttributeValue

	// The item in the DynamoDB table as it appeared after it was modified.
	NewImage map[string]*AttributeValue

	// The size of the stream record, in bytes.
	SizeBytes int64
}

// EventRecord provides contextual information about an Amazon DynamoDB Streams
// record.
type EventRecord struct {
	// A globally unique identifier for the event that was recorded in this
	// stream record.
	EventID string

	// The type of data modification that was performed on the DynamoDB table:
	//   * INSERT - a new item was added to the table.
	//   * MODIFY - one or more of an existing item's attributes were modified.
	//   * REMOVE - the item was deleted from the table.
	EventName string

	// The version number of the stream record format. This number is updated
	// whenever the structure of Record is modified.
	//
	// Client applications must not assume that eventVersion will remain at a
	// particular value, as this number is subject to change at any time. In
	// general, eventVersion will only increase as the low-level Amazon DynamoDB
	// Streams API evolves.
	EventVersion string

	// The AWS service from which the stream record originated.
	// For Amazon DynamoDB Streams, this is aws:dynamodb.
	EventSource string

	// The region in which the GetRecords request was received.
	AWSRegion string

	// The main body of the stream record, containing all of the
	// DynamoDB-specific fields.
	DynamoDB *Record
}

// Event represents an Amazon DynamoDB Streams event.
type Event struct {
	// The list of Amazon DynamoDB Streams event records.
	Records []*EventRecord
}

// String returns the string representation
func (e Event) String() string {
	s, _ := json.MarshalIndent(e, "", "  ")
	return string(s)
}

// GoString returns the string representation
func (e Event) GoString() string {
	return e.String()
}