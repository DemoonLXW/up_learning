package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DemoonLXW/up_learning/config"
)

type WorkflowService struct {
	*config.FlowableConfig
}

func (serv *WorkflowService) startAProcessInstance(reqBody map[string]interface{}) error {

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("start a processInstance failed: %w", err)
	}
	reqBodyBuffer := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"runtime/process-instances", reqBodyBuffer)
	if err != nil {
		return fmt.Errorf("start a processInstance failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("start a processInstance failed: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return fmt.Errorf("start a processInstance failed: %d\n%s", resp.StatusCode, b)
	}

	return nil
}
