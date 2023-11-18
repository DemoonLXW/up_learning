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

func (serv *WorkflowService) StartAProcessInstance(reqBody map[string]interface{}) error {

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
func (serv *WorkflowService) QueryForHistoricProcessInstances(reqBody map[string]interface{}) (map[string]interface{}, error) {
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("query for historic process instances failed: %w", err)
	}
	reqBodyBuffer := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"query/historic-process-instances", reqBodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("query for historic process instances failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("query for historic process instances failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query for historic process instances failed: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("query for historic process instances failed: %w", err)
	}
	resp.Body.Close()

	var respBody map[string]interface{}
	err = json.Unmarshal(b, &respBody)
	if err != nil {
		return nil, fmt.Errorf("query for historic process instances failed: %w", err)
	}

	return respBody, nil
}

func (serv *WorkflowService) QueryForHistoricTaskInstances(reqBody map[string]interface{}) (map[string]interface{}, error) {
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	}
	reqBodyBuffer := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"query/historic-task-instances", reqBodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query for historic task instances failed: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	}
	resp.Body.Close()

	var respBody map[string]interface{}
	err = json.Unmarshal(b, &respBody)
	if err != nil {
		return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	}

	return respBody, nil
}
