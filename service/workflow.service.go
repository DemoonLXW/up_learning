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

func (serv *WorkflowService) doRequest(method string, url string, reqBody map[string]interface{}, header map[string]string) (*http.Response, error) {
	var reqBodyBuffer *bytes.Buffer
	if reqBody != nil {
		reqBodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("do request failed: %w", err)
		}
		reqBodyBuffer = bytes.NewBuffer(reqBodyBytes)
	}

	req, err := http.NewRequest(method, url, reqBodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	c := http.Client{}
	return c.Do(req)
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
func (serv *WorkflowService) QueryForHistoricProcessInstances(reqBody map[string]interface{}) ([]byte, error) {
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

	// var respBody map[string]interface{}
	// err = json.Unmarshal(b, &respBody)
	// if err != nil {
	// 	return nil, fmt.Errorf("query for historic process instances failed: %w", err)
	// }

	return b, nil
}

func (serv *WorkflowService) QueryForHistoricTaskInstances(reqBody map[string]interface{}) ([]byte, error) {
	// reqBodyBytes, err := json.Marshal(reqBody)
	// if err != nil {
	// 	return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	// }
	// reqBodyBuffer := bytes.NewBuffer(reqBodyBytes)

	// req, err := http.NewRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"query/historic-task-instances", reqBodyBuffer)
	// if err != nil {
	// 	return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	// }
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// c := http.Client{}
	// resp, err := c.Do(req)
	// if err != nil {
	// 	return nil, fmt.Errorf("query for historic task instances failed: %w", err)
	// }
	header := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
	resp, err := serv.doRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"query/historic-task-instances", reqBody, header)
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

	return b, nil
}

func (serv *WorkflowService) QueryForTasks(reqBody map[string]interface{}) ([]byte, error) {
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("query for tasks failed: %w", err)
	}
	reqBodyBuffer := bytes.NewBuffer(reqBodyBytes)

	req, err := http.NewRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"query/tasks", reqBodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("query for tasks failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("query for tasks failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query for tasks failed: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("query for tasks failed: %w", err)
	}
	resp.Body.Close()

	return b, nil
}

// func (serv *WorkflowService) GetATask(taskID string) ([]byte, error) {
// 	header := map[string]string{
// 		"Content-Type": "application/json; charset=utf-8",
// 	}

// 	resp, err := serv.doRequest(http.MethodGet, nil, *serv.FlowableConfig.URL+"/"+"runtime/tasks/"+taskID, header)
// 	if err != nil {
// 		return nil, fmt.Errorf("query a task failed: %w", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("query a task failed: %d", resp.StatusCode)
// 	}

// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("query a task failed: %w", err)
// 	}
// 	resp.Body.Close()

// 	return b, nil
// }

func (serv *WorkflowService) TaskActions(taskID string, reqBody map[string]interface{}) error {
	header := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
	resp, err := serv.doRequest(http.MethodPost, *serv.FlowableConfig.URL+"/"+"runtime/tasks/"+taskID, reqBody, header)
	if err != nil {
		return fmt.Errorf("task actions failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("task actions failed: %w", err)
		}
		resp.Body.Close()
		return fmt.Errorf("task actions failed: %d\n%s", resp.StatusCode, b)
	}

	return nil

}
