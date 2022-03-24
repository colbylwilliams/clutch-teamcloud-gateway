package teamcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "strings"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"

	teamcloudconfigv1 "github.com/colbylwilliams/clutch-teamcloud-gateway/backend/api/config/service/teamcloud/v1"
	teamcloudv1 "github.com/colbylwilliams/clutch-teamcloud-gateway/backend/api/teamcloud/v1"
	sdk "github.com/colbylwilliams/teamcloud-go/teamcloud"
	"github.com/lyft/clutch/backend/service"
)

const Name = "colbylwilliams.service.teamcloud"

// const apiHost = "https://teamclouddemo-api.azurewebsites.net"

func New(cfg *any.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	config := &teamcloudconfigv1.Config{}
	if err := cfg.UnmarshalTo(config); err != nil {
		return nil, err
	}

	tc := sdk.NewClient(nil)

	return &client{
		http:     &http.Client{},
		tc:       tc,
		hostUrl:  config.Host,
		tenantId: config.TenantId,
		clientId: config.ClientId,
		log:      logger,
	}, nil
}

type Client interface {
	// GetTeamCloudInfo(ctx context.Context, name string) (*teamcloudv1.GetTeamCloudInfoResponse, error)
	GetTeamCloudInfo(ctx context.Context) (*teamcloudv1.GetTeamCloudInfoResponse, error)
	GetInfo(ctx context.Context) (*teamcloudv1.TeamCloudInformationDataResult, error)
	GetAdapters(ctx context.Context) (*teamcloudv1.AdapterInformationListDataResult, error)
	CancelComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error)
	CreateComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error)
	GetComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error)
	GetComponentTasks(ctx context.Context) (*teamcloudv1.ComponentTaskListDataResult, error)
	ReRunComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error)
	GetComponentTemplate(ctx context.Context) (*teamcloudv1.ComponentTemplateDataResult, error)
	GetComponentTemplates(ctx context.Context) (*teamcloudv1.ComponentTemplateListDataResult, error)
	CreateComponent(ctx context.Context) (*teamcloudv1.ComponentDataResult, error)
	DeleteComponent(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetComponent(ctx context.Context) (*teamcloudv1.ComponentDataResult, error)
	GetComponents(ctx context.Context) (*teamcloudv1.ComponentListDataResult, error)
	InitializeAuthorization(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error)
	CreateDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error)
	DeleteDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error)
	GetDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error)
	GetDeploymentScopes(ctx context.Context) (*teamcloudv1.DeploymentScopeListDataResult, error)
	UpdateDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error)
	// NegotiateSignalR (ctx context.Context (p*teamcloudv1.rotobuf, error.Empty)
	GetAuditCommands(ctx context.Context) (*teamcloudv1.StringListDataResult, error)
	GetAuditEntries(ctx context.Context) (*teamcloudv1.CommandAuditEntityListDataResult, error)
	GetAuditEntry(ctx context.Context) (*teamcloudv1.CommandAuditEntityDataResult, error)
	CreateOrganizationUser(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	DeleteOrganizationUser(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetOrganizationUser(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	GetOrganizationUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	GetOrganizationUsers(ctx context.Context) (*teamcloudv1.UserListDataResult, error)
	UpdateOrganizationUser(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	UpdateOrganizationUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	CreateOrganization(ctx context.Context) (*teamcloudv1.OrganizationDataResult, error)
	DeleteOrganization(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetOrganization(ctx context.Context) (*teamcloudv1.OrganizationDataResult, error)
	GetOrganizations(ctx context.Context) (*teamcloudv1.OrganizationListDataResult, error)
	CreateProjectIdentity(ctx context.Context) (*teamcloudv1.ProjectIdentityDataResult, error)
	DeleteProjectIdentity(ctx context.Context) (*teamcloudv1.ProjectIdentityDataResult, error)
	GetProjectIdentities(ctx context.Context) (*teamcloudv1.ProjectIdentityListDataResult, error)
	GetProjectIdentity(ctx context.Context) (*teamcloudv1.ProjectIdentityDataResult, error)
	UpdateProjectIdentity(ctx context.Context) (*teamcloudv1.StatusResult, error)
	CreateProject(ctx context.Context) (*teamcloudv1.ProjectDataResult, error)
	DeleteProject(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetProject(ctx context.Context) (*teamcloudv1.ProjectDataResult, error)
	GetProjects(ctx context.Context) (*teamcloudv1.ProjectListDataResult, error)
	CreateProjectTag(ctx context.Context) (*teamcloudv1.StatusResult, error)
	DeleteProjectTag(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetProjectTagByKey(ctx context.Context) (*teamcloudv1.StringStringDictionaryDataResult, error)
	GetProjectTags(ctx context.Context) (*teamcloudv1.StringStringDictionaryDataResult, error)
	UpdateProjectTag(ctx context.Context) (*teamcloudv1.StatusResult, error)
	CreateProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error)
	DeleteProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error)
	GetProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error)
	GetProjectTemplates(ctx context.Context) (*teamcloudv1.ProjectTemplateListDataResult, error)
	UpdateProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error)
	CreateProjectUser(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	DeleteProjectUser(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetProjectUser(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	GetProjectUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	GetProjectUsers(ctx context.Context) (*teamcloudv1.UserListDataResult, error)
	UpdateProjectUser(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	UpdateProjectUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error)
	CreateSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error)
	GetSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error)
	GetSchedules(ctx context.Context) (*teamcloudv1.ScheduleListDataResult, error)
	RunSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error)
	UpdateSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error)
	GetProjectStatus(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetStatus(ctx context.Context) (*teamcloudv1.StatusResult, error)
	GetUserProjects(ctx context.Context) (*teamcloudv1.ProjectListDataResult, error)
	GetUserProjectsMe(ctx context.Context) (*teamcloudv1.ProjectListDataResult, error)
}

type client struct {
	http     *http.Client
	tc       *sdk.Client
	hostUrl  string
	clientId string
	tenantId string
	log      *zap.Logger
}

// type RawResponse struct {
// 	Amiibo *RawTeamCloudInfo `json:"amiibo"`
// }

type RawTeamCloudInfo struct {
	Code   int32                `json:"code"`
	Status string               `json:"status"`
	Data   RawTeamCloudInfoData `json:"data"`
}

type RawTeamCloudInfoData struct {
	ImageVersion    string `json:"imageVersion"`
	TemplateVersion string `json:"templateVersion"`
}

func (r RawTeamCloudInfo) toProto() *teamcloudv1.GetTeamCloudInfoResponse {
	return &teamcloudv1.GetTeamCloudInfoResponse{
		Code:   r.Code,
		Status: r.Status,
		Data:   r.Data.toProto(),
	}
}

func (r RawTeamCloudInfoData) toProto() *teamcloudv1.TeamCloudInfoData {
	return &teamcloudv1.TeamCloudInfoData{
		ImageVersion:    r.ImageVersion,
		TemplateVersion: r.TemplateVersion,
	}
}

func charactersFromJSON(data []byte) (*teamcloudv1.GetTeamCloudInfoResponse, error) {
	raw := &RawTeamCloudInfo{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	ret := raw.toProto()
	return ret, nil
}

// func (c *client) GetTeamCloudInfo(ctx context.Context, name string) ([]*amiibov1.Amiibo, error) {
func (c *client) GetTeamCloudInfo(ctx context.Context) (*teamcloudv1.GetTeamCloudInfoResponse, error) {
	// url := fmt.Sprintf("%s/api/amiibo?character=%s", apiHost, name)
	url := fmt.Sprintf("%s/", c.hostUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, status.Error(service.CodeFromHTTPStatus(resp.StatusCode), string(body))
	}

	tcinfo, tcerr := c.tc.GetInfo(ctx, &sdk.ClientGetInfoOptions{})

	if tcerr != nil {
		return nil, tcerr
	}

	c.log.Info("teamcloud info", zap.Any("info", tcinfo))

	return charactersFromJSON(body)
}

func (c *client) GetInfo(ctx context.Context) (*teamcloudv1.TeamCloudInformationDataResult, error) {
	thing := teamcloudv1.TeamCloudInformationDataResult{}
	return &thing, nil
}

func (c *client) GetAdapters(ctx context.Context) (*teamcloudv1.AdapterInformationListDataResult, error) {
	thing := teamcloudv1.AdapterInformationListDataResult{}
	return &thing, nil
}

func (c *client) CancelComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error) {
	thing := teamcloudv1.ComponentTaskDataResult{}
	return &thing, nil
}

func (c *client) CreateComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error) {
	thing := teamcloudv1.ComponentTaskDataResult{}
	return &thing, nil
}

func (c *client) GetComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error) {
	thing := teamcloudv1.ComponentTaskDataResult{}
	return &thing, nil
}

func (c *client) GetComponentTasks(ctx context.Context) (*teamcloudv1.ComponentTaskListDataResult, error) {
	thing := teamcloudv1.ComponentTaskListDataResult{}
	return &thing, nil
}

func (c *client) ReRunComponentTask(ctx context.Context) (*teamcloudv1.ComponentTaskDataResult, error) {
	thing := teamcloudv1.ComponentTaskDataResult{}
	return &thing, nil
}

func (c *client) GetComponentTemplate(ctx context.Context) (*teamcloudv1.ComponentTemplateDataResult, error) {
	thing := teamcloudv1.ComponentTemplateDataResult{}
	return &thing, nil
}

func (c *client) GetComponentTemplates(ctx context.Context) (*teamcloudv1.ComponentTemplateListDataResult, error) {
	thing := teamcloudv1.ComponentTemplateListDataResult{}
	return &thing, nil
}

func (c *client) CreateComponent(ctx context.Context) (*teamcloudv1.ComponentDataResult, error) {
	thing := teamcloudv1.ComponentDataResult{}
	return &thing, nil
}

func (c *client) DeleteComponent(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetComponent(ctx context.Context) (*teamcloudv1.ComponentDataResult, error) {
	thing := teamcloudv1.ComponentDataResult{}
	return &thing, nil
}

func (c *client) GetComponents(ctx context.Context) (*teamcloudv1.ComponentListDataResult, error) {
	thing := teamcloudv1.ComponentListDataResult{}
	return &thing, nil
}

func (c *client) InitializeAuthorization(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error) {
	thing := teamcloudv1.DeploymentScopeDataResult{}
	return &thing, nil
}

func (c *client) CreateDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error) {
	thing := teamcloudv1.DeploymentScopeDataResult{}
	return &thing, nil
}

func (c *client) DeleteDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error) {
	thing := teamcloudv1.DeploymentScopeDataResult{}
	return &thing, nil
}

func (c *client) GetDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error) {
	thing := teamcloudv1.DeploymentScopeDataResult{}
	return &thing, nil
}

func (c *client) GetDeploymentScopes(ctx context.Context) (*teamcloudv1.DeploymentScopeListDataResult, error) {
	thing := teamcloudv1.DeploymentScopeListDataResult{}
	return &thing, nil
}

func (c *client) UpdateDeploymentScope(ctx context.Context) (*teamcloudv1.DeploymentScopeDataResult, error) {
	thing := teamcloudv1.DeploymentScopeDataResult{}
	return &thing, nil
}

func (c *client) GetAuditCommands(ctx context.Context) (*teamcloudv1.StringListDataResult, error) {
	thing := teamcloudv1.StringListDataResult{}
	return &thing, nil
}

func (c *client) GetAuditEntries(ctx context.Context) (*teamcloudv1.CommandAuditEntityListDataResult, error) {
	thing := teamcloudv1.CommandAuditEntityListDataResult{}
	return &thing, nil
}

func (c *client) GetAuditEntry(ctx context.Context) (*teamcloudv1.CommandAuditEntityDataResult, error) {
	thing := teamcloudv1.CommandAuditEntityDataResult{}
	return &thing, nil
}

func (c *client) CreateOrganizationUser(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) DeleteOrganizationUser(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetOrganizationUser(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) GetOrganizationUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) GetOrganizationUsers(ctx context.Context) (*teamcloudv1.UserListDataResult, error) {
	thing := teamcloudv1.UserListDataResult{}
	return &thing, nil
}

func (c *client) UpdateOrganizationUser(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) UpdateOrganizationUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) CreateOrganization(ctx context.Context) (*teamcloudv1.OrganizationDataResult, error) {
	thing := teamcloudv1.OrganizationDataResult{}
	return &thing, nil
}

func (c *client) DeleteOrganization(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetOrganization(ctx context.Context) (*teamcloudv1.OrganizationDataResult, error) {
	thing := teamcloudv1.OrganizationDataResult{}
	return &thing, nil
}

func (c *client) GetOrganizations(ctx context.Context) (*teamcloudv1.OrganizationListDataResult, error) {
	thing := teamcloudv1.OrganizationListDataResult{}
	return &thing, nil
}

func (c *client) CreateProjectIdentity(ctx context.Context) (*teamcloudv1.ProjectIdentityDataResult, error) {
	thing := teamcloudv1.ProjectIdentityDataResult{}
	return &thing, nil
}

func (c *client) DeleteProjectIdentity(ctx context.Context) (*teamcloudv1.ProjectIdentityDataResult, error) {
	thing := teamcloudv1.ProjectIdentityDataResult{}
	return &thing, nil
}

func (c *client) GetProjectIdentities(ctx context.Context) (*teamcloudv1.ProjectIdentityListDataResult, error) {
	thing := teamcloudv1.ProjectIdentityListDataResult{}
	return &thing, nil
}

func (c *client) GetProjectIdentity(ctx context.Context) (*teamcloudv1.ProjectIdentityDataResult, error) {
	thing := teamcloudv1.ProjectIdentityDataResult{}
	return &thing, nil
}

func (c *client) UpdateProjectIdentity(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) CreateProject(ctx context.Context) (*teamcloudv1.ProjectDataResult, error) {
	thing := teamcloudv1.ProjectDataResult{}
	return &thing, nil
}

func (c *client) DeleteProject(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetProject(ctx context.Context) (*teamcloudv1.ProjectDataResult, error) {
	thing := teamcloudv1.ProjectDataResult{}
	return &thing, nil
}

func (c *client) GetProjects(ctx context.Context) (*teamcloudv1.ProjectListDataResult, error) {
	thing := teamcloudv1.ProjectListDataResult{}
	return &thing, nil
}

func (c *client) CreateProjectTag(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) DeleteProjectTag(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetProjectTagByKey(ctx context.Context) (*teamcloudv1.StringStringDictionaryDataResult, error) {
	thing := teamcloudv1.StringStringDictionaryDataResult{}
	return &thing, nil
}

func (c *client) GetProjectTags(ctx context.Context) (*teamcloudv1.StringStringDictionaryDataResult, error) {
	thing := teamcloudv1.StringStringDictionaryDataResult{}
	return &thing, nil
}

func (c *client) UpdateProjectTag(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) CreateProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error) {
	thing := teamcloudv1.ProjectTemplateDataResult{}
	return &thing, nil
}

func (c *client) DeleteProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error) {
	thing := teamcloudv1.ProjectTemplateDataResult{}
	return &thing, nil
}

func (c *client) GetProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error) {
	thing := teamcloudv1.ProjectTemplateDataResult{}
	return &thing, nil
}

func (c *client) GetProjectTemplates(ctx context.Context) (*teamcloudv1.ProjectTemplateListDataResult, error) {
	thing := teamcloudv1.ProjectTemplateListDataResult{}
	return &thing, nil
}

func (c *client) UpdateProjectTemplate(ctx context.Context) (*teamcloudv1.ProjectTemplateDataResult, error) {
	thing := teamcloudv1.ProjectTemplateDataResult{}
	return &thing, nil
}

func (c *client) CreateProjectUser(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) DeleteProjectUser(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetProjectUser(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) GetProjectUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) GetProjectUsers(ctx context.Context) (*teamcloudv1.UserListDataResult, error) {
	thing := teamcloudv1.UserListDataResult{}
	return &thing, nil
}

func (c *client) UpdateProjectUser(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) UpdateProjectUserMe(ctx context.Context) (*teamcloudv1.UserDataResult, error) {
	thing := teamcloudv1.UserDataResult{}
	return &thing, nil
}

func (c *client) CreateSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error) {
	thing := teamcloudv1.ScheduleDataResult{}
	return &thing, nil
}

func (c *client) GetSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error) {
	thing := teamcloudv1.ScheduleDataResult{}
	return &thing, nil
}

func (c *client) GetSchedules(ctx context.Context) (*teamcloudv1.ScheduleListDataResult, error) {
	thing := teamcloudv1.ScheduleListDataResult{}
	return &thing, nil
}

func (c *client) RunSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error) {
	thing := teamcloudv1.ScheduleDataResult{}
	return &thing, nil
}

func (c *client) UpdateSchedule(ctx context.Context) (*teamcloudv1.ScheduleDataResult, error) {
	thing := teamcloudv1.ScheduleDataResult{}
	return &thing, nil
}

func (c *client) GetProjectStatus(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetStatus(ctx context.Context) (*teamcloudv1.StatusResult, error) {
	thing := teamcloudv1.StatusResult{}
	return &thing, nil
}

func (c *client) GetUserProjects(ctx context.Context) (*teamcloudv1.ProjectListDataResult, error) {
	thing := teamcloudv1.ProjectListDataResult{}
	return &thing, nil
}

func (c *client) GetUserProjectsMe(ctx context.Context) (*teamcloudv1.ProjectListDataResult, error) {
	thing := teamcloudv1.ProjectListDataResult{}
	return &thing, nil
}
