package base

// ClientQuery is the client query data.
type ClientQuery struct {
	// OnlyActive filters the records in Status Active.
	OnlyActive bool
	// RQL is the query data is considered to be a RQL json as defined in https://github.com/a8m/rql.
	//
	//	{
	//		"limit": 0,
	//		"offset": 0,
	//		"filter": {
	//			"$or": [
	//				{ "Name": "TLV" },
	//				{ "Quantity": { "$gte": 49800, "$lte": 57080 } }
	//			],
	//			"$and":[
	//				{"Email":{"$like":"asd@gmail.com%"}},
	//				{"Status":2}
	//			]
	//		},
	//		"sort": ["-Quantity"]
	//	}
	RQL []byte
}
