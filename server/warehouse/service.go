//go:generate protoc ./warehouse.proto --go_out=plugins=grpc:../api

package warehouse

import (
	"context"
	"fmt"
	"time"

	api "github.com/craigericlewis/warehouse/server/api"
	"github.com/craigericlewis/warehouse/server/utils"
)

type Service interface {
	GetPackageInfo(ctx context.Context, packageName, packageSize string, language api.Language) (*api.PackageInfo, error)
	getPythonPackageInfo(packageName string) (*api.PackageInfo, error)
}

type warehouseService struct {
}

type pythonPackageInfo struct {
	Name        string `json:"string"`
	Description string `json:"description"`
	PackageURL  string `json:"package_url"`
	HomePage    string `json:"home_page"`
	ProjectURLs struct {
		Code string `json:"Code"`
	} `json:"project_urls"`
	RequiresDist []string `json:"requires_dist"`
	Version      string   `json:"version"`
	Releases     map[string][]struct {
		UploadTime string `json:"upload_time"`
		ZippedSize int64  `json:"size"`
	} `json:"releases"`
}

func NewService() Service {
	return &warehouseService{}
}

func (s warehouseService) GetPackageInfo(
	ctx context.Context,
	packageName,
	packageSize string,
	language api.Language,
) (*api.PackageInfo, error) {
	if language == api.Language_PYTHON3 {
		packageInfo, err := s.getPythonPackageInfo(packageName)

		if err != nil {
			return nil, err
		}

	}
}

func (s warehouseService) getPythonPackageInfo(
	packageName string,
) (*api.PackageInfo, error) {
	packageInfo := &pythonPackageInfo{}
	url := fmt.Sprintf("https://pypi.python.org/pypi/%s/json", packageName)
	err := utils.GetJson(url, 5*time.Minute, packageInfo)

	if err != nil {
		return nil, err
	}

	resp := &api.PackageInfo{
		Name:        packageName,
		Description: packageInfo.Description,
		Urls: &api.Urls{
			GithubLink: packageInfo.ProjectURLs.Code,
			PyPiLink:   packageInfo.PackageURL,
		},
	}

	return resp, nil
}
