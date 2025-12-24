package initializers

import (
	"github.com/majoalba20/redSocialPeriferiaIT/cmd/internal/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
