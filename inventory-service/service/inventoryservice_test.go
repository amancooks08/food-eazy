package service

import (
	"testing"

	"inventory-service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthServiceTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *AuthServiceTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	models.InitInventoryModels(db)
}

func (suite *AuthServiceTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&models.Item{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (suite *AuthServiceTestSuite) TestAddItem() {
	t := suite.T()
	type args struct {
		name        string
		description string
		price       float64
		quantity    uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Add item successfully",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       100,
				quantity:    100,
			},
			wantErr: false,
		},
		{
			name: "Add item with empty name",
			args: args{
				name:        "",
				description: "testitem1.desc",
				price:       100,
				quantity:    100,
			},
			wantErr: true,
		},
		{
			name: "Add item with empty description",
			args: args{
				name:        "testitem1",
				description: "",
				price:       100,
				quantity:    100,
			},
			wantErr: true,
		},
		{
			name: "Add item with zero price",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       0,
				quantity:    100,
			},
			wantErr: true,
		},
		{
			name: "Add item with zero quantity",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       100,
				quantity:    0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := AddItem(tt.args.name, tt.args.description, tt.args.price, tt.args.quantity)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, item)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, item)
				assert.Equal(t, tt.args.name, item.Name)
				assert.Equal(t, tt.args.description, item.Description)
				assert.Equal(t, tt.args.price, item.Price)
				assert.Equal(t, tt.args.quantity, item.Quantity)
			}
		})
	}
}
