# Greenhouse
[![CircleCI](https://circleci.com/gh/Hundred5/greenhouse/tree/master.svg?style=svg)](https://circleci.com/gh/Hundred5/greenhouse/tree/master)

## Greenhouse Ingestion API
https://developers.greenhouse.io/candidate-ingestion.html

## Installation
Install using the "go get" command:
```sh
go get github.com/hundred5/greenhouse/ingestion
```

### Example

```go
import "github.com/hundred5/greenhouse/ingestion"
```

Construct a new Ingestion Client and then work off of that. For example, to create a new candidate:
```go
client := ingestion.NewClient("accessToken", nil)
candidates, err := client.Candidates.Post([]ingestion.PostCandidate{
    ingestion.PostCandidate{
        ExternalID: "externalID",
        Addresses: []ingestion.Address{
            ingestion.Address{
                Address: "1600 Pennsylvania Ave NW, Washington, DC 20500, USA",
                Type:    ingestion.AddressTypeWork,
            },
            ingestion.Address{
                Address: "1 Infinite Loop, Cupertino, CA 95014, USA",
                Type:    ingestion.AddressTypeOther,
            },
        },
        //...
    },
})
```

errors.go provides some convenience functions for checking errors:
```go
if clientError, ok := ingestion.IsClientError(err); ok {
    // Client error - you messed up
    // Print status code & errors
    fmt.Printf("StatusCode: %d\n", clientError.StatusCode)
    for _, e := range clientError.Errors {
        fmt.Printf("Message: %s; Field: %s\n", e.Message, e.Field)
    }
}
if serverError, ok := ingestion.IsServerError(err); ok {
    // Server error - try again later
    // Print status code & errors
    fmt.Printf("StatusCode: %d\n", serverError.StatusCode)
    for _, e := range serverError.Errors {
        fmt.Printf("Message: %s; Field: %s\n", e.Message, e.Field)
    }
}
```