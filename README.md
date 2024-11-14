# server_kulunvalvonta
Älykkäät laitteet serveri puoli

# luo uusi entiteetti ohjeet
1. Luo model
```go
package models

import "context"

type Person struct {
	ID         int    `json:"id"`
	PersonID   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	RoomID     int    `json:"room_id"`
}

type PersonRepository interface {
	CreatePerson(Data *Person, ctx context.Context) error
	ReadOnePerson(id int, ctx context.Context) (*Person, error)
	//ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*Data, error)
	//Update(data *Data, ctx context.Context) (int64, error)
	//Delete(data *Data, ctx context.Context) (int64, error)
}
```
