package domain

/* Data Representation */

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uint   `gorm:"primaryKey" json:"id"`                  //will automatically map "uint" with "serial type" in the postgres
	Name        string `gorm:"type:text;unique;not null" json:"name"` // e.g., "admin", "user", "moderator"
	Description string `gorm:"type:text" json:"description"`
}

type Permission struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"` // e.g., "add_pipeline", "view_dashboard", "manage_users",
}

type RolePermission struct {
	ID           uint       `gorm:"primaryKey"`
	RoleID       uint       `gorm:"not null"`
	PermissionID uint       `gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE" json:"-"`
}

// User represents a user in the system
type Profile struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;references:auth.users(id);onDelete:CASCADE" json:"id"`
	Name   string    `gorm:"type:text" json:"name"`
	Email  string    `gorm:"type:text;unique;not null" json:"email"`
	RoleID uint      `gorm:"not null" json:"role_id"`
	Role   Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"-"` //user is allowed to exist even if role is deleted
}

// type Pipeline struct {
// 	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
// 	Name   string    `gorm:"type:text;not null" json:"name"`
// 	UserID uuid.UUID `gorm:"type:uuid;not null;constraint:foreignKey:UserID;references:users(ID)" json:"user_id"` //foreign key for connecting pipeline with user
// 	//Not sure about how this Multi-Valued attribute will be translated: [This should be removed and a foreign key should be added in the Stages table referencing to Pipelines tables primary key]
// 	Stages []Stage `json:"stages"`
// }

// Pipeline represents a user's pipeline
type Pipeline struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name    string    `gorm:"not null" json:"name"`
	UserID  uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Profile Profile   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"` //if user entry is delete then delete all his pipelines
	// Stages []Stage   `json:"stages"` // Relationship
}

// type Stage struct {
// 	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
// 	Name       string    `gorm:"type:text" json:"name"`
// 	PipelineID uuid.UUID `gorm:"type:uuid;not null;constraint:foreignKey:PipelineID:pipelines(ID)" json:"pipeline_id"` //Foreign key for connecting user with stage
// }

// Stage represents a stage in a pipeline
type Stage struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	PipelineID uuid.UUID `gorm:"type:uuid;not null" json:"pipeline_id"`
	Pipeline   Pipeline  `gorm:"foreignKey:PipelineID;constraint:OnDelete:CASCADE" json:"-"`
}

// PipelineExecution tracks a pipeline run
type PipelineExecution struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PipelineID uuid.UUID `gorm:"type:uuid;not null" json:"pipeline_id"`
	Pipeline   Pipeline  `gorm:"foreignKey:PipelineID;constraint:OnDelete:CASCADE" json:"-"`
	Status     string    `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	StartedAt  time.Time `gorm:"default:now()" json:"started_at"`
	EndedAt    time.Time `json:"ended_at,omitempty"`
}

// StageExecution tracks the execution of a stage
type StageExecution struct {
	ID           uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StageID      uuid.UUID         `gorm:"type:uuid;not null" json:"stage_id"`
	ExecutionID  uuid.UUID         `gorm:"type:uuid;not null" json:"execution_id"`
	Stage        Stage             `gorm:"foreignKey:StageID;constraint:OnDelete:CASCADE" json:"-"`
	Execution    PipelineExecution `gorm:"foreignKey:ExecutionID" json:"-"`
	Status       string            `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	StartedAt    time.Time         `gorm:"default:now()" json:"started_at"`
	EndedAt      *time.Time        `json:"ended_at,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
}

// Status Type for telling the status of pipeline //
type Status string

const (
	Pending   Status = "Pending"
	Running   Status = "Running"
	Failed    Status = "Failed"
	Success   Status = "Success"
	Completed Status = "Completed"
	Unknown   Status = "Unknown" //Error occured before fetching status
)
