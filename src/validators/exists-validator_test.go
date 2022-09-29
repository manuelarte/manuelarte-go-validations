package validators

//import (
//	"fmt"
//	"github.com/go-playground/validator/v10"
//	"github.com/stretchr/testify/assert"
//	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
//	"gorm.io/gorm"
//	"testing"
//)
//
//func TestExistsValidator_EntityExists(t *testing.T) {
//
//	type testEntityStruct struct {
//		ID   uint
//		Name string
//	}
//	type testRequestStruct struct {
//		EntityID uint `validate:"exists=test_entity_structs.id"`
//	}
//
//	tests := map[string]struct {
//		Entity   testEntityStruct
//		entityID uint
//	}{
//		"testing finding an existing entity": {
//			Entity:   testEntityStruct{ID: 1, Name: "Manuel"},
//			entityID: 1,
//		},
//	}
//
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			database, err := NewInMemoryDatabase(t)
//			if err != nil {
//				assert.Failf(t, "Can't create the database: %s", err.Error())
//			}
//			err = database.AutoMigrate(testEntityStruct{})
//			if err != nil {
//				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
//			}
//			tx := database.Debug().Model(tc.Entity).Create(&tc.Entity)
//			if tx.Error != nil {
//				assert.Failf(t, "Can't create entity", "Can't create entity: %v", tc.Entity)
//			}
//			validate := validator.New()
//			existsValidator := NewExistsValidator(database)
//			err = existsValidator.RegisterExistsValidator(validate)
//			if err != nil {
//				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
//			}
//			err = validate.Struct(testRequestStruct{tc.entityID})
//			assert.Nil(t, err)
//		})
//	}
//}
//
//func TestExistsValidator_ArrayEntityExists(t *testing.T) {
//
//	type testEntityStruct struct {
//		ID   uint
//		Name string
//	}
//	type testRequestStruct struct {
//		EntityIDs []uint `validate:"exists=test_entity_structs.id"`
//	}
//
//	tests := map[string]struct {
//		entitiesToMigrate []*testEntityStruct
//		entityIDs         []uint
//	}{
//		"testing finding an existing entity": {
//			entitiesToMigrate: []*testEntityStruct{{ID: 1, Name: "Manuel"}},
//			entityIDs:         []uint{1},
//		},
//		"testing finding existing and non existing entity": {
//			entitiesToMigrate: []*testEntityStruct{{ID: 1, Name: "Manuel"}, {ID: 2, Name: "Elisa"}},
//			entityIDs:         []uint{1, 2},
//		},
//	}
//
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			database, err := NewInMemoryDatabase(t)
//			if err != nil {
//				assert.Failf(t, "Can't create the database: %s", err.Error())
//			}
//			err = database.AutoMigrate(testEntityStruct{})
//			if err != nil {
//				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
//			}
//			tx := database.Debug().Model(tc.entitiesToMigrate[0]).CreateInBatches(&tc.entitiesToMigrate, len(tc.entitiesToMigrate))
//			if tx.Error != nil {
//				assert.Failf(t, "Can't create entity", "Can't create entity: %v", tc.entitiesToMigrate)
//			}
//			validate := validator.New()
//			existsValidator := NewExistsValidator(database)
//			err = existsValidator.RegisterExistsValidator(validate)
//			if err != nil {
//				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
//			}
//			err = validate.Struct(testRequestStruct{tc.entityIDs})
//			assert.Nil(t, err)
//		})
//	}
//}
//
//func TestExistsValidator_EntityDoesNotExists(t *testing.T) {
//
//	type testEntityStruct struct {
//		ID   uint
//		Name string
//	}
//	type testRequestStruct struct {
//		EntityID uint `validate:"exists=test_entity_structs.id"`
//	}
//
//	tests := map[string]struct {
//		Entity           testEntityStruct
//		entityID         uint
//		expectedErrorMsg string
//	}{
//		"testing finding a non existing entity": {
//			Entity:           testEntityStruct{ID: 1, Name: "Manuel"},
//			entityID:         2,
//			expectedErrorMsg: "Key: 'testRequestStruct.EntityID' Error:Field validation for 'EntityID' failed on the 'exists' tag",
//		},
//	}
//
//	for name, tc := range tests {
//		t.Run(name, func(t *testing.T) {
//			database, err := NewInMemoryDatabase(t)
//			if err != nil {
//				assert.Failf(t, "Can't create the database: %s", err.Error())
//			}
//			err = database.AutoMigrate(testEntityStruct{})
//			if err != nil {
//				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
//			}
//			tx := database.Debug().Model(tc.Entity).Create(&tc.Entity)
//			if tx.Error != nil {
//				assert.Failf(t, "Can't create entity", "Can't create entity: %v", tc.Entity)
//			}
//			validate := validator.New()
//			existsValidator := NewExistsValidator(database)
//			err = existsValidator.RegisterExistsValidator(validate)
//			if err != nil {
//				assert.Failf(t, "Can't migrate entity", "Can't auto migrate entity: %s", err.Error())
//			}
//			err = validate.Struct(testRequestStruct{tc.entityID})
//			assert.NotNil(t, err)
//			assert.Equal(t, tc.expectedErrorMsg, err.Error())
//		})
//	}
//}
//
//func NewInMemoryDatabase(t *testing.T) (database *gorm.DB, err error) {
//	connectionString := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
//	database, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
//	return
//}
