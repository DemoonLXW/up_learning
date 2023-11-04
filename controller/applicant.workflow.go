package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/DemoonLXW/up_learning/database/ent"
	"go.uber.org/cadence"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

func (cont *ApplicantController) ReviewProjectWorkflow(ctx workflow.Context, reviewProjectID uint32, applicantID uint32) error {
	logger := workflow.GetLogger(ctx)
	// info := workflow.GetInfo(ctx)
	logger.Info("result", zap.Uint32("projectID", reviewProjectID), zap.Uint32("applicant", applicantID))
	info := workflow.GetInfo(ctx)
	expiration := time.Duration(info.ExecutionStartToCloseTimeoutSeconds) * time.Second
	retryPolicy := &cadence.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2,
		MaximumInterval:    10 * time.Second,
		ExpirationInterval: expiration,
		MaximumAttempts:    10,
	}
	ao := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 15 * 24 * time.Hour,
		HeartbeatTimeout:       time.Minute,
		RetryPolicy:            retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	DBctx := context.Background()
	client := cont.Applicant.DB

	logger.Info("Start Activity")
	future := workflow.ExecuteActivity(ctx, cont.Applicant.FindUserWithTeacherOrStudentById, DBctx, client, applicantID)
	user := ent.User{}
	err := future.Get(ctx, &user)
	// user, err := cont.Applicant.FindUserWithTeacherOrStudentById(DBctx, client, applicantID)
	if err != nil {
		logger.Error("Activity err" + err.Error())
		return fmt.Errorf("review project find user failed: %w", err)
	}
	logger.Info("Find user", zap.String("account", user.Account))
	// if user.Edges.Teacher == nil && user.Edges.Student == nil {
	// 	return fmt.Errorf("review project user is neither teacher nor student")
	// }

	return nil
}
