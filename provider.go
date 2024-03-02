package ddnsman

import (
	"encoding/json"
	"fmt"

	"github.com/libdns/acmeproxy"
	"github.com/libdns/alidns"
	"github.com/libdns/azure"
	"github.com/libdns/bunny"
	"github.com/libdns/civo"
	"github.com/libdns/cloudflare"
	"github.com/libdns/ddnss"
	"github.com/libdns/desec"
	"github.com/libdns/digitalocean"
	"github.com/libdns/dinahosting"
	"github.com/libdns/dnspod"
	"github.com/libdns/dnsupdate"
	"github.com/libdns/duckdns"
	"github.com/libdns/dynv6"
	"github.com/libdns/gandi"
	"github.com/libdns/godaddy"
	"github.com/libdns/googleclouddns"
	"github.com/libdns/hetzner"
	"github.com/libdns/hexonet"
	"github.com/libdns/ionos"
	"github.com/libdns/leaseweb"
	"github.com/libdns/libdns"
	"github.com/libdns/linode"
	"github.com/libdns/loopia"
	"github.com/libdns/mailinabox"
	"github.com/libdns/metaname"
	"github.com/libdns/mythicbeasts"
	"github.com/libdns/namecheap"
	"github.com/libdns/namedotcom"
	"github.com/libdns/njalla"
	designate "github.com/libdns/openstack-designate"
	"github.com/libdns/ovh"
	"github.com/libdns/powerdns"
	"github.com/libdns/rfc2136"
	"github.com/libdns/route53"
	"github.com/libdns/scaleway"
	"github.com/libdns/tencentcloud"
	"github.com/libdns/totaluptime"
	"github.com/libdns/transip"
	"github.com/libdns/vercel"
	"github.com/libdns/vultr"
)

// Providers allows get and set records to DNS provider.
type Provider interface {
	libdns.RecordGetter
	libdns.RecordSetter
}

func newProvider(providerName string, data json.RawMessage) (Provider, error) {
	switch providerName {
	case "acmeproxy":
		return readProvider[acmeproxy.Provider](data)
	case "alidns":
		return readProvider[alidns.Provider](data)
	case "azure":
		return readProvider[azure.Provider](data)
	case "bunny":
		return readProvider[bunny.Provider](data)
	case "civo":
		return readProvider[civo.Provider](data)
	case "cloudflare":
		return readProvider[cloudflare.Provider](data)
	case "ddnss":
		return readProvider[ddnss.Provider](data)
	case "desec":
		return readProvider[desec.Provider](data)
	case "digitalocean":
		return readProvider[digitalocean.Provider](data)
	case "dinahosting":
		return readProvider[dinahosting.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "directadmin":
	// 	return readProvider[directadmin.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "dnsmadeeasy":
	// 	return readProvider[dnsmadeeasy.Provider](data)
	case "dnspod":
		return readProvider[dnspod.Provider](data)
	case "dnsupdate":
		return readProvider[dnsupdate.Provider](data)
	case "duckdns":
		return readProvider[duckdns.Provider](data)
	case "dynv6":
		return readProvider[dynv6.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "easydns":
	// 	return readProvider[easydns.Provider](data)
	case "gandi":
		return readProvider[gandi.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "glesys":
	// 	return readProvider[glesys.Provider](data)
	case "godaddy":
		return readProvider[godaddy.Provider](data)
	case "googleclouddns":
		return readProvider[googleclouddns.Provider](data)
	case "hetzner":
		return readProvider[hetzner.Provider](data)
	case "hexonet":
		return readProvider[hexonet.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "hosttech":
	// 	return readProvider[hosttech.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "infomaniak":
	// 	return readProvider[infomaniak.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "inwx":
	// 	return readProvider[inwx.Provider](data)
	case "ionos":
		return readProvider[ionos.Provider](data)
	case "leaseweb":
		return readProvider[leaseweb.Provider](data)
	case "linode":
		return readProvider[linode.Provider](data)
	case "loopia":
		return readProvider[loopia.Provider](data)
	case "mailinabox":
		return readProvider[mailinabox.Provider](data)
	case "metaname":
		return readProvider[metaname.Provider](data)
	case "mythicbeasts":
		return readProvider[mythicbeasts.Provider](data)
	case "namecheap":
		return readProvider[namecheap.Provider](data)
	case "namedotcom":
		return readProvider[namedotcom.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "namesilo":
	// 	return readProvider[namesilo.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "netcup":
	// 	return readProvider[netcup.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "netlify":
	// 	return readProvider[netlify.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "nicrudns":
	// 	return readProvider[nicrudns.Provider](data)
	case "njalla":
		return readProvider[njalla.Provider](data)
	case "openstack-designate":
		return readProvider[designate.Provider](data)
	case "ovh":
		return readProvider[ovh.Provider](data)
	// FIXME: Broken API due to int to uint conversion.
	// case "porkbun":
	// 	return readProvider[porkbun.Provider](data)
	case "powerdns":
		return readProvider[powerdns.Provider](data)
	case "rfc2136":
		return readProvider[rfc2136.Provider](data)
	case "route53":
		return readProvider[route53.Provider](data)
	case "scaleway":
		return readProvider[scaleway.Provider](data)
	case "tencentcloud":
		return readProvider[tencentcloud.Provider](data)
	case "totaluptime":
		return readProvider[totaluptime.Provider](data)
	case "transip":
		return readProvider[transip.Provider](data)
	case "vercel":
		return readProvider[vercel.Provider](data)
	case "vultr":
		return readProvider[vultr.Provider](data)
	default:
		return nil, fmt.Errorf("unknown provider %q", providerName)
	}
}

func readProvider[T any](data json.RawMessage) (*T, error) {
	provider := new(T)
	if err := json.Unmarshal(data, provider); err != nil {
		return nil, fmt.Errorf("unable to parse provider settings: %w", err)
	}
	return provider, nil
}
