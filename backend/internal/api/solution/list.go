package solution

import (
	"context"
	"fmt"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/converter"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
	"go.uber.org/zap"
)

// List Solutions.
func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	logger.Info("List solutions...")
	solutions, err := i.solutionService.List(ctx)
	if err != nil {
		logger.Error("failed to list solutions", zap.Error(err))
		return nil, InternalError
	}

	var descSolutions []*desc.Solution
	for _, solution := range solutions {
		logger.Info(fmt.Sprintf("solution: %+v", solution))
		descSolutions = append(descSolutions, converter.ToSolutionFromService(solution))
	}

	return &desc.ListResponse{
		Solutions: descSolutions,
	}, nil
}
