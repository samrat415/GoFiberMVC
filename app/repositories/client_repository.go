package repositories

import (
	"GoFiberMVC/app/initializers"
	"GoFiberMVC/app/models"
	"fmt"
)

// ClientRepository represents the repository for managing clients
type ClientRepository struct {
}

// Save saves an client to the database
func (r *ClientRepository) Save(client *models.Client) (*models.Client, error) {
	db := initializers.Db
	result := db.Create(&client)
	if result.Error != nil {
		return nil, result.Error
	}
	return client, nil
}

// Find finds an client by ID
func (r *ClientRepository) Find(id uint) (*models.Client, error) {
	db := initializers.Db
	client := &models.Client{}
	result := db.Preload("Provider").First(client, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return client, nil
}

func (r *ClientRepository) LoadAll() ([]*models.Client, error) {
	db := initializers.Db
	var clients []*models.Client
	result := db.Debug().Find(&clients)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to load clients: %s", result.Error.Error())
	}
	return clients, nil
}

// FindAll retrieves all clients based on the provided query parameters
func (r *ClientRepository) FindAll(queryParams map[string]interface{}) ([]*models.Client, error) {
	db := initializers.Db
	var clients []*models.Client
	// Create a query builder
	queryBuilder := db.Model(&models.Client{})
	// Apply pagination
	if perPage, ok := queryParams["perPage"].(int); ok && perPage > 0 {
		queryBuilder = queryBuilder.Limit(perPage)
	}
	// Apply Page
	if page, ok := queryParams["page"].(int); ok && page > 0 {
		queryBuilder = queryBuilder.Offset(page)
	}
	// Apply filtering conditions
	if name, ok := queryParams["name"].(string); ok && name != "" {
		queryBuilder = queryBuilder.Where("name LIKE ?", "%"+name+"%")
	}
	queryBuilder.Where("is_active", true)
	if isActive, ok := queryParams["is_active"].(bool); ok {
		queryBuilder = queryBuilder.Where("is_active = ?", isActive)
	}

	// Apply sorting
	if sort, ok := queryParams["sort"].(string); ok && sort != "" {
		queryBuilder = queryBuilder.Order(sort)
	}
	// Retrieve the clients
	if err := queryBuilder.Preload("Provider").Find(&clients).Error; err != nil {
		return nil, fmt.Errorf("failed to load clients: %s", err.Error())
	}
	return clients, nil
}
