/* Here Behaviour of each method is defined. */

package ports

import (
	"context"

	"github.com/ashwin-pf9/pipeline-shared/domain"
	"github.com/google/uuid"
)

// What it is: This Interface defines behaviour of a process which is responsible for Creating stages, running stages, retrieving status of each stage
// The AddStage method will be used for adding a new stage to pipeline.
type PipelineOrchestratorInterface interface {
	// Add a new stage to the pipeline
	AddStage(stage domain.Stage) error

	// Execute the entire pipeline
	Execute(ctx context.Context, input interface{}) (interface{}, error)

	// Get pipeline execution status
	GetStatus(pipelineID uuid.UUID) (domain.Status, error)

	// Cancel pipeline execution
	Cancel(pipelineID uuid.UUID) error
}
