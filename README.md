# Introduction

This go library has some simple "general purpose" validations to be used with [Go Playground Validator](https://github.com/go-playground/validator)

# Validators

## IsValid Validator

Exists validator allows you to validate a struct by using an inside function isValid.

### Example

```go
type MyStruct struct {
	MyName Name `validate:"is-valid"`
}

type Name string

func (n Name) IsValid() {
	return len(n) > 5 && len(n) < 40
}
```

## Exists Validator

The exists validator queries the database based on the table name and field in the validation param.

### Example

#### Example with primitive field

```go
type MyStruct struct {
	FriendID int `binding:"exists:my-structs.id"`
}
```

The validator will query the table `my-structs` get the column `id` and check that there is a record.

#### Example with slice field

```go
type MyStruct struct {
	FriendIDs []int `binding:"exists:my-structs.id"`
}
```

The validator will query the table `my-structs` get the column `id` and check that there are as many records as the length of FriendsIDs.