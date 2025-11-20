package main

// This is script for generate database table to struct
// run "go run technical_service/orm/scaffolder.go"

import (
	"Go_OOP/Models"
	"fmt"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type foreignKeyInfo struct {
	TableName        string
	ColumnName       string
	ReferencedTable  string
	ReferencedColumn string
}

type Constraint struct {
	Table            string
	ForeignKeys      []string
	ReferenceTable   string
	ReferenceColumns []string
}

func getForeignKeys(db *gorm.DB) ([]foreignKeyInfo, error) {
	var results []foreignKeyInfo
	query := `
		SELECT 
			fk_tab.name AS TableName,
			fk_col.name AS ColumnName,
			pk_tab.name AS ReferencedTable,
			pk_col.name AS ReferencedColumn
		FROM sys.foreign_key_columns fkc
		INNER JOIN sys.tables fk_tab ON fk_tab.object_id = fkc.parent_object_id
		INNER JOIN sys.columns fk_col ON fk_col.column_id = fkc.parent_column_id AND fk_col.object_id = fk_tab.object_id
		INNER JOIN sys.tables pk_tab ON pk_tab.object_id = fkc.referenced_object_id
		INNER JOIN sys.columns pk_col ON pk_col.column_id = fkc.referenced_column_id AND pk_col.object_id = pk_tab.object_id;
	`
	err := db.Raw(query).Scan(&results).Error
	return results, err
}

// Scaffold generates models with relationships automatically
func main() {
	var adapter *Models.Adapter
	adapter = adapter.GetAdapterIntance()
	db := adapter.GetGormInstance()

	// Configure GORM Generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./Technical_Service/Entity",
		ModelPkgPath: "Entity/EntityStruct",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)

	// Step 1: Get all tables from DB
	allTables, err := db.Migrator().GetTables()
	if err != nil {
		panic(fmt.Errorf("failed to list tables: %w", err))
	}

	// Step 2: Generate basic models (no relations yet)
	models := make(map[string]any)
	for _, table := range allTables {
		models[table] = g.GenerateModel(table)
	}

	// Step 3: Discover foreign key relationships
	fks, err := getForeignKeys(db)
	if err != nil {
		fmt.Println("⚠️  Warning: failed to load foreign keys:", err)
	}

	// Step 4: Attach relationships (HasMany / BelongsTo)
	for _, fk := range fks {
		parent := g.GenerateModel(fk.ReferencedTable)
		child := g.GenerateModel(fk.TableName)

		// Add BelongsTo on child
		childRel := gen.FieldRelate(
			field.BelongsTo,
			fk.ReferencedTable, // relationship name
			parent,
			&field.RelateConfig{
				RelatePointer: true,
				GORMTag: field.GormTag{}.
					Set("foreignKey", fk.ColumnName).
					Set("references", fk.ReferencedColumn),
			},
		)

		// Add HasMany on parent
		parentRel := gen.FieldRelate(
			field.HasMany,
			fk.TableName,
			child,
			&field.RelateConfig{
				RelateSlicePointer: true,
				GORMTag: field.GormTag{}.
					Set("foreignKey", fk.ColumnName).
					Set("references", fk.ReferencedColumn),
			},
		)

		// Update models with relationships
		models[fk.TableName] = g.GenerateModelAs(fk.TableName, fk.TableName, childRel)
		models[fk.ReferencedTable] = g.GenerateModelAs(fk.ReferencedTable, fk.ReferencedTable, parentRel)
	}

	// Step 5: Apply all models and execute
	var all []any
	for _, m := range models {
		all = append(all, m)
	}

	g.ApplyBasic(all...)
	g.Execute()

	fmt.Println("✅ Scaffold generation complete!")
}
