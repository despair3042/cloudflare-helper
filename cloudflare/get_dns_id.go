package cloudflare

import (
	"encoding/json"
	"net/http"

	"github.com/tdl3/cloudflare-helper/models"
	"github.com/tdl3/cloudflare-helper/models/resp"
	"github.com/tdl3/cloudflare-helper/utils"
	"go.uber.org/zap"
)

// This will get the dns_id from cloudflare, and store it in the cfg
func GetDNSId(cfg *models.Config) (StatusCode int) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", utils.AssembleUrl(*cfg, utils.GetDnsId), nil)
	if err != nil {
		zap.L().Fatal("Something went wrong", zap.Error(err))
	}
	request.Header.Add("Authorization", cfg.Cloudflare.BearerToken)
	response, err := client.Do(request)
	if err != nil {
		zap.L().Fatal("Can not resolve host", zap.Error(err))
	}
	defer response.Body.Close()
	var content resp.DNSId
	err = json.NewDecoder(response.Body).Decode(&content)
	if err != nil {
		zap.L().Info("Can not parse response", zap.Error(err))
	}
	zap.L().Debug(utils.PrettyJson(content))
	cfg.Cloudflare.DnsId = content.Result[0].Id
	return response.StatusCode
}
