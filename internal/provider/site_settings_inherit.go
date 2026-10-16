package provider

import (
	"fmt"
	"net/url"
	"strings"

	cc "github.com/netascode/go-catalystcenter"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// siteParentId returns the ID of the parent site, or "" for Global and for sites that no longer exist.
func siteParentId(client *cc.Client, siteId string) (string, error) {
	res, err := client.Get("/dna/intent/api/v1/site?siteId=" + url.QueryEscape(siteId))
	if err != nil {
		if res.Get("response.errorCode").String() == "NCGR10008" {
			return "", nil
		}
		return "", err
	}
	return res.Get("response.parentId").String(), nil
}

// siteSettingsDeleteBody returns the PUT body that removes a site settings override.
// Global has no parent, so it gets globalBody, which unsets the setting; so does a site that no longer exists.
// Other sites get the parent's effective settings: Catalyst Center does not store an override
// equal to the inherited value, so this restores inheritance. A null body is no alternative,
// as time zone rejects it and AAA stores it as a copy of the parent's values.
func siteSettingsDeleteBody(client *cc.Client, endpoint, siteId, globalBody string) (string, error) {
	parentId, err := siteParentId(client, siteId)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve parent site of %s: %w", siteId, err)
	}
	if parentId == "" {
		return globalBody, nil
	}
	res, err := client.Get(fmt.Sprintf(endpoint, url.QueryEscape(parentId)) + "?_inherited=true")
	if err != nil {
		return "", fmt.Errorf("failed to retrieve settings of parent site %s: %w", parentId, err)
	}
	body := "{}"
	res.Get("response").ForEach(func(key, value gjson.Result) bool {
		raw := value.Raw
		if value.IsObject() {
			value.ForEach(func(field, _ gjson.Result) bool {
				if strings.HasPrefix(field.String(), "inherited") {
					raw, _ = sjson.Delete(raw, field.String())
				}
				return true
			})
		}
		body, _ = sjson.SetRaw(body, key.String(), raw)
		return true
	})
	return body, nil
}
