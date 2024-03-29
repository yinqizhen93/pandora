// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CasbinRulesColumns holds the columns for the "casbin_rules" table.
	CasbinRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ptype", Type: field.TypeString, Default: ""},
		{Name: "v0", Type: field.TypeString, Default: ""},
		{Name: "v1", Type: field.TypeString, Default: ""},
		{Name: "v2", Type: field.TypeString, Default: ""},
		{Name: "v3", Type: field.TypeString, Default: ""},
		{Name: "v4", Type: field.TypeString, Default: ""},
		{Name: "v5", Type: field.TypeString, Default: ""},
	}
	// CasbinRulesTable holds the schema information for the "casbin_rules" table.
	CasbinRulesTable = &schema.Table{
		Name:       "casbin_rules",
		Columns:    CasbinRulesColumns,
		PrimaryKey: []*schema.Column{CasbinRulesColumns[0]},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeInt},
		{Name: "is_deleted", Type: field.TypeInt8, Default: 0},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "updated_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "department_children1", Type: field.TypeInt, Nullable: true},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "departments_departments_children1",
				Columns:    []*schema.Column{DepartmentsColumns[9]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MaterialsColumns holds the columns for the "materials" table.
	MaterialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "describe", Type: field.TypeString, Size: 2147483647},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "buy_date", Type: field.TypeTime},
	}
	// MaterialsTable holds the schema information for the "materials" table.
	MaterialsTable = &schema.Table{
		Name:       "materials",
		Columns:    MaterialsColumns,
		PrimaryKey: []*schema.Column{MaterialsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "descript", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "is_deleted", Type: field.TypeInt8, Default: 0},
		{Name: "access_api", Type: field.TypeString},
		{Name: "access_method", Type: field.TypeString, SchemaType: map[string]string{"mysql": "char(6)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// StocksColumns holds the columns for the "stocks" table.
	StocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "market", Type: field.TypeString, SchemaType: map[string]string{"mysql": "char(4)"}},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "open", Type: field.TypeFloat32},
		{Name: "close", Type: field.TypeFloat32},
		{Name: "high", Type: field.TypeFloat32},
		{Name: "low", Type: field.TypeFloat32},
		{Name: "volume", Type: field.TypeInt32},
		{Name: "outstanding_share", Type: field.TypeInt32},
		{Name: "turnover", Type: field.TypeFloat32},
	}
	// StocksTable holds the schema information for the "stocks" table.
	StocksTable = &schema.Table{
		Name:       "stocks",
		Columns:    StocksColumns,
		PrimaryKey: []*schema.Column{StocksColumns[0]},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"once", "periodical"}},
		{Name: "describe", Type: field.TypeString, Size: 2147483647},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "cost_time", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeInt},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "phone_number", Type: field.TypeString, Unique: true},
		{Name: "refresh_token", Type: field.TypeString, Default: ""},
		{Name: "is_active", Type: field.TypeInt8, Default: 1},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_department", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_departments_department",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CasbinRulesTable,
		DepartmentsTable,
		MaterialsTable,
		RolesTable,
		StocksTable,
		TasksTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	DepartmentsTable.ForeignKeys[0].RefTable = DepartmentsTable
	UsersTable.ForeignKeys[0].RefTable = DepartmentsTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}
