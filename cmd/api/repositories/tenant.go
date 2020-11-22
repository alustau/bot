package repositories

import (
	"database/sql"
	"github.com/cgauge/bot/cmd/api/models"
)

type TenantRepository struct {
	db *sql.DB
}

func NewTenantRepository(db *sql.DB) *TenantRepository {
	repository := &TenantRepository{}
	repository.db = db

	return repository
}

func (r *TenantRepository) GetAllTenants() (tenants []*models.Tenant, err error) {
	results, err := r.db.Query("select id, company_name from cgcore_companies")

	if err != nil {
		return nil, err
	}

	for results.Next() {
		var tenant models.Tenant
		err = results.Scan(&tenant.ID, &tenant.Name)

		if err != nil {
			return nil, err
		}

		tenants = append(tenants, &tenant)
	}

	return tenants, nil
}
