package cluster

import (
	"context"
	"fmt"
	"terraform-provider-celerdatabyoc/celerdata-sdk/client"
	"terraform-provider-celerdatabyoc/celerdata-sdk/version"
)

type IClusterAPI interface {
	Deploy(ctx context.Context, req *DeployReq) (*DeployResp, error)
	Get(ctx context.Context, req *GetReq) (*GetResp, error)
	Release(ctx context.Context, req *ReleaseReq) (*ReleaseResp, error)
	Suspend(ctx context.Context, req *SuspendReq) (*SuspendResp, error)
	Resume(ctx context.Context, req *ResumeReq) (*ResumeResp, error)
	GetState(ctx context.Context, req *GetStateReq) (*GetStateResp, error)
	ScaleIn(ctx context.Context, req *ScaleInReq) (*ScaleInResp, error)
	ScaleOut(ctx context.Context, req *ScaleOutReq) (*ScaleOutResp, error)
	ScaleUp(ctx context.Context, req *ScaleUpReq) (*ScaleUpResp, error)
	IncrStorageSize(ctx context.Context, req *IncrStorageSizeReq) (*IncrStorageSizeResp, error)
	UnlockFreeTier(ctx context.Context, clusterID string) error
}

func NewClustersAPI(cli *client.CelerdataClient) IClusterAPI {
	return &clusterAPI{cli: cli, apiVersion: version.API_1_0}
}

type clusterAPI struct {
	cli        *client.CelerdataClient
	apiVersion version.ApiVersion
}

func (c *clusterAPI) Deploy(ctx context.Context, req *DeployReq) (*DeployResp, error) {
	resp := &DeployResp{}
	err := c.cli.Post(ctx, fmt.Sprintf("/api/%s/clusters", c.apiVersion), req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) Get(ctx context.Context, req *GetReq) (*GetResp, error) {
	resp := &GetResp{}
	err := c.cli.Get(ctx, fmt.Sprintf("/api/%s/clusters/%s", c.apiVersion, req.ClusterID), nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) Release(ctx context.Context, req *ReleaseReq) (*ReleaseResp, error) {
	resp := &ReleaseResp{}
	err := c.cli.Delete(ctx, fmt.Sprintf("/api/%s/clusters/%s", c.apiVersion, req.ClusterID), nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) GetState(ctx context.Context, req *GetStateReq) (*GetStateResp, error) {
	resp := &GetStateResp{}
	err := c.cli.Get(ctx, fmt.Sprintf("/api/%s/clusters/%s/state", c.apiVersion, req.ClusterID), map[string]string{"action_id": req.ActionID}, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) Suspend(ctx context.Context, req *SuspendReq) (*SuspendResp, error) {
	resp := &SuspendResp{}
	err := c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/suspend", c.apiVersion, req.ClusterID), nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) Resume(ctx context.Context, req *ResumeReq) (*ResumeResp, error) {
	resp := &ResumeResp{}
	err := c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/resume", c.apiVersion, req.ClusterID), nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) UnlockFreeTier(ctx context.Context, clusterID string) error {
	return c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/unlock-free-tier", c.apiVersion, clusterID), nil, nil)
}

func (c *clusterAPI) ScaleIn(ctx context.Context, req *ScaleInReq) (*ScaleInResp, error) {
	resp := &ScaleInResp{}
	err := c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/scale-in", c.apiVersion, req.ClusterId), req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) ScaleOut(ctx context.Context, req *ScaleOutReq) (*ScaleOutResp, error) {
	resp := &ScaleOutResp{}
	err := c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/scale-out", c.apiVersion, req.ClusterId), req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) ScaleUp(ctx context.Context, req *ScaleUpReq) (*ScaleUpResp, error) {
	resp := &ScaleUpResp{}
	err := c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/scale-up", c.apiVersion, req.ClusterId), req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clusterAPI) IncrStorageSize(ctx context.Context, req *IncrStorageSizeReq) (*IncrStorageSizeResp, error) {
	resp := &IncrStorageSizeResp{}
	err := c.cli.Patch(ctx, fmt.Sprintf("/api/%s/clusters/%s/incr-storage-size", c.apiVersion, req.ClusterId), req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
