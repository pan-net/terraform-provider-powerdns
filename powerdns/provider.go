package powerdns

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_API_KEY", nil),
				Description: "REST API authentication key",
			},
			"server_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_SERVER_URL", nil),
				Description: "Location of PowerDNS server",
			},
			"skip_tls_verify": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_SKIP_TLS_VERIFY", false),
				Description: "Disable verification of the PowerDNS server's TLS certificate",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"powerdns_zone":   resourcePDNSZone(),
			"powerdns_record": resourcePDNSRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiKey:        data.Get("api_key").(string),
		ServerUrl:     data.Get("server_url").(string),
		SkipTLSVerify: data.Get("skip_tls_verify").(bool),
	}

	return config.Client()
}
