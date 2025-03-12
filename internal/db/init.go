package db

import (
	"accounts/internal/autogen"
	"accounts/internal/config"
	"accounts/internal/models"
	"database/sql"
	"errors"
	"slices"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalPermissions map[autogen.GlobalPermission]models.Permission = map[autogen.GlobalPermission]models.Permission{}

func InitDB() (*sql.DB, *gorm.DB, error) {

    dsn := config.GetConfig().DbConfig.Dsn
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        logrus.Fatalf("open db error: %v", err)
        return nil, nil, err
    }

    // Get the underlying sql.DB object to close the connection later
    sqlDB, err := db.DB()
    if err != nil {
        logrus.Errorf("Error getting database: %v", err)
        return nil, nil, err
    }

    // Ping the database to check if the connection is successful
    err = sqlDB.Ping()
    if err != nil {
        logrus.Errorf("Error pinging database: %v", err)
        return nil, nil, errors.New("ping db error")
    }

    logrus.Info("Database connection successful")

    // Perform auto migration
    err = db.AutoMigrate(
        models.Permission{},
        models.Role{},
        models.User{},
    )
    if err != nil {
        logrus.Errorf("Error auto migrating database : %v", err)
        return nil, nil, err
    }
    logrus.Info("Database auto migration completed")

    err = populateDatabase(db)
    if err != nil {
        logrus.Errorf("Error populating database : %v", err)
        return nil, nil, err
    }

	return sqlDB, db, nil
}

// Populate database with default roles and permissions
func populateDatabase(db *gorm.DB) error {
    initPermissions(db)
    initDefaultRoles(db)
    return nil
}

// Initialize all permissions
func initPermissions(db *gorm.DB) {

    logrus.Debug("Populating permissions")

    for _, name := range(autogen.GlobalPermission_values) {

        var perm models.Permission;
        res := db.Where(models.Permission{Name: name}).Limit(1).Find(&perm)

        if res.Error != nil || res.RowsAffected == 0 {
            perm.Name = name;
            db.Create(&perm)
            logrus.Infof("Added permission %s to database", name)
        }

        GlobalPermissions[name] = perm;
    }

    logrus.Debug("Populated permissions successfully")
}

// Create a default role or update its permissions
func createOrUpdateRole(db *gorm.DB, role *models.Role) {

    var currentRole models.Role;
    result := db.Preload("Permissions").Where(models.Role{Name: role.Name}).Limit(1).Find(&currentRole);

    // Create role if it does not exist
    if result.Error != nil || result.RowsAffected == 0 {
        db.Create(role)
        logrus.Infof("Created %s role with id %d", role.Name, role.ID)
        return;
    }

    needs_update := false;

    // Check for permissions to remove
    for _, perm := range(currentRole.Permissions) {
        if !slices.ContainsFunc(role.Permissions, func(p models.Permission) bool {return perm.ID == p.ID;}) {
            needs_update = true;
            break;
        }
    }

    // Check for permissions to add
    for _, perm := range(role.Permissions) {
        if !slices.ContainsFunc(currentRole.Permissions, func(p models.Permission) bool {return perm.ID == p.ID;}) {
            needs_update = true;
            break;
        }
    }

    if needs_update {
        currentRole.Permissions = role.Permissions;
        db.Save(&currentRole);

        perm_names := make([]string, len(role.Permissions))
        for i, p := range(role.Permissions) {
            perm_names[i] = string(p.Name)
        }
        logrus.Infof("Updated %s role's permissions to %v", role.Name, perm_names)
    }
}

// Initialize default roles
func initDefaultRoles(db *gorm.DB) {

    logrus.Debug("Populating default roles");

    // Admin
    admin := &models.Role{
        Name: "Admin",
        IsDefault: true,
        Permissions: []models.Permission{
            GlobalPermissions[autogen.GlobalPermissionADMIN],
        },
    }
    createOrUpdateRole(db, admin)

    // BDE
    bde := &models.Role{
        Name: "BDE",
        IsDefault: true,
        Permissions: []models.Permission {
            GlobalPermissions[autogen.GlobalPermissionMANAGEUSERSINFO],
            GlobalPermissions[autogen.GlobalPermissionMANAGECLUBS],
        },
    }
    createOrUpdateRole(db, bde);

    // Bureau Restreint BDE
    bde_br := &models.Role{
        Name: "Bureau Restreint BDE",
        IsDefault: true,
        Permissions: []models.Permission {
            GlobalPermissions[autogen.GlobalPermissionMANAGEROLES],
            GlobalPermissions[autogen.GlobalPermissionMANAGEUSERSROLES],
        },
    }
    createOrUpdateRole(db, bde_br);

    logrus.Debug("Populated default roles successfully")
}
