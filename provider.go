package ddnsman

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/libdns/libdns"

	"github.com/libdns/acmedns"
	"github.com/libdns/acmeproxy"
	"github.com/libdns/alidns"
	"github.com/libdns/azure"
	"github.com/libdns/bunny"
	"github.com/libdns/cloudflare"
	"github.com/libdns/cloudns"
	"github.com/libdns/desec"
	"github.com/libdns/digitalocean"
	"github.com/libdns/directadmin"
	"github.com/libdns/dnsimple"
	"github.com/libdns/domainnameshop"
	"github.com/libdns/duckdns"
	"github.com/libdns/dynu"
	"github.com/libdns/edgeone"
	"github.com/libdns/gandi"
	"github.com/libdns/gcore"
	"github.com/libdns/glesys"
	"github.com/libdns/googleclouddns"
	"github.com/libdns/he"
	"github.com/libdns/hetzner"
	"github.com/libdns/huaweicloud"
	"github.com/libdns/infomaniak"
	"github.com/libdns/inwx"
	"github.com/libdns/ionos"
	"github.com/libdns/loopia"
	"github.com/libdns/luadns"
	"github.com/libdns/mailinabox"
	"github.com/libdns/metaname"
	"github.com/libdns/mijnhost"
	"github.com/libdns/namesilo"
	"github.com/libdns/netcup"
	"github.com/libdns/nfsn"
	"github.com/libdns/ovh"
	"github.com/libdns/porkbun"
	"github.com/libdns/regfish"
	"github.com/libdns/rfc2136"
	"github.com/libdns/scaleway"
	"github.com/libdns/simplydotcom"
	"github.com/libdns/tencentcloud"
	"github.com/libdns/westcn"
)

// Providers allows get and set records to DNS provider.
type Provider interface {
	libdns.RecordGetter
	libdns.RecordSetter
}

func newProvider(providerName string, data json.RawMessage) (Provider, error) {
	// Missing providers:
	// - dode - no [libdns.RecordGetter]/[libdns.RecordSetter] support.
	// - nanelo - no [libdns.RecordGetter]/[libdns.RecordSetter] support.
	// - civo – does not compatible with current libdns, public archive.

	switch providerName {
	case "acmedns":
		slog.Warn("acmedns is in a beta state. It may not work as expected.")
		return readProvider[acmedns.Provider](data)
	case "acmeproxy":
		return readProvider[acmeproxy.Provider](data)
	case "alidns":
		slog.Warn("alidns is in a beta state. It may not work as expected.")
		return readProvider[alidns.Provider](data)
	// FIXME: autodns is not supported yet: https://github.com/libdns/autodns/pull/8
	// case "autodns":
	// 	return readProvider[autodns.Provider](data)
	case "azure":
		return readProvider[azure.Provider](data)
	case "bunny":
		return readProvider[bunny.Provider](data)
	case "cloudflare":
		return readProvider[cloudflare.Provider](data)
	case "cloudns":
		return readProvider[cloudns.Provider](data)
	// FIXME: ddnss is not supported yet: https://github.com/libdns/ddnss/issues/2
	// case "ddnss":
	// 	return readProvider[ddnss.Provider](data)
	case "desec":
		return readProvider[desec.Provider](data)
	case "digitalocean":
		return readProvider[digitalocean.Provider](data)
	// FIXME: dinahosting is not supported yet: https://github.com/libdns/dinahosting/issues/1
	// case "dinahosting":
	// 	return readProvider[dinahosting.Provider](data)
	case "directadmin":
		return readProvider[directadmin.Provider](data)
	// FIXME: dnsexit is not supported yet.
	// case "dnsexit":
	// 	return readProvider[dnsexit.Provider](data)
	case "dnsimple":
		return readProvider[dnsimple.Provider](data)
	// FIXME: dnsmadeeasy is not supported yet: https://github.com/libdns/dnsmadeeasy/issues/8
	// case "dnsmadeeasy":
	// 	return readProvider[dnsmadeeasy.Provider](data)
	// FIXME: dnspod is not supported yet: https://github.com/libdns/dnspod/pull/9
	// case "dnspod":
	// 	return readProvider[dnspod.Provider](data)
	// FIXME: dnsupdate is not supported yet.
	// case "dnsupdate":
	// 	return readProvider[dnsupdate.Provider](data)
	case "domainnameshop":
		return readProvider[domainnameshop.Provider](data)
	// FIXME: dreamhost is not supported yet.
	// case "dreamhost":
	// 	return readProvider[dreamhost.Provider](data)
	case "duckdns":
		return readProvider[duckdns.Provider](data)
	case "dynu":
		return readProvider[dynu.Provider](data)
	// FIXME: dynv6 is not supported yet: https://github.com/libdns/dynv6/issues/1
	// case "dynv6":
	// 	return readProvider[dynv6.Provider](data)
	// FIXME: easydns is not supported yet: https://github.com/libdns/easydns/pull/2
	// case "easydns":
	// 	return readProvider[easydns.Provider](data)
	case "edgeone":
		return readProvider[edgeone.Provider](data)
	// FIXME: exoscale is not supported yet.
	// case "exoscale":
	// 	return readProvider[exoscale.Provider](data)
	case "gandi":
		return readProvider[gandi.Provider](data)
	case "gcore":
		return readProvider[gcore.Provider](data)
	case "glesys":
		return readProvider[glesys.Provider](data)
	// FIXME: godaddy is not supported yet: https://github.com/libdns/godaddy/issues/9
	// case "godaddy":
	// 	return readProvider[godaddy.Provider](data)
	case "googleclouddns":
		return readProvider[googleclouddns.Provider](data)
	case "he":
		return readProvider[he.Provider](data)
	case "hetzner":
		return readProvider[hetzner.Provider](data)
	// FIXME: hexonet is not supported yet: https://github.com/libdns/hexonet/issues/1
	// case "hexonet":
	// 	return readProvider[hexonet.Provider](data)
	// FIXME: hosttech is not supported yet: https://github.com/libdns/hosttech/pull/17
	// case "hosttech":
	// 	return readProvider[hosttech.Provider](data)
	case "huaweicloud":
		slog.Warn("huaweicloud is in a beta state. It may not work as expected.")
		return readProvider[huaweicloud.Provider](data)
	case "infomaniak":
		return readProvider[infomaniak.Provider](data)
	case "inwx":
		return readProvider[inwx.Provider](data)
	case "ionos":
		return readProvider[ionos.Provider](data)
	// FIXME: katapult is not supported yet.
	// case "katapult":
	// 	return readProvider[katapult.Provider](data)
	// FIXME: leaseweb is not supported yet.
	// case "leaseweb":
	// 	return readProvider[leaseweb.Provider](data)
	// FIXME: linode is not supported yet: https://github.com/libdns/linode/issues/19
	// case "linode":
	// 	return readProvider[linode.Provider](data)
	case "loopia":
		return readProvider[loopia.Provider](data)
	case "luadns":
		return readProvider[luadns.Provider](data)
	case "mailinabox":
		return readProvider[mailinabox.Provider](data)
	case "metaname":
		return readProvider[metaname.Provider](data)
	case "mijnhost":
		return readProvider[mijnhost.Provider](data)
	// FIXME: namecheap is not supported yet.
	// case "mythicbeasts":
	// 	return readProvider[mythicbeasts.Provider](data)
	// FIXME: namecheap is not supported yet: https://github.com/libdns/namecheap/pull/11
	// case "namecheap":
	// 	return readProvider[namecheap.Provider](data)
	// FIXME: namedotcom is not supported yet: https://github.com/libdns/namedotcom/issues/3
	// case "namedotcom":
	// 	return readProvider[namedotcom.Provider](data)
	case "namesilo":
		return readProvider[namesilo.Provider](data)
	// FIXME: neoserv is not supported yet.
	// case "neoserv":
	// 	return readProvider[neoserv.Provider](data)
	case "netcup":
		return readProvider[netcup.Provider](data)
	// FIXME: netlify is not supported yet: https://github.com/libdns/netlify/issues/6
	// case "netlify":
	// 	return readProvider[netlify.Provider](data)
	case "nfsn":
		return readProvider[nfsn.Provider](data)
	// FIXME: nicrudns is not supported yet.
	// case "nicrudns":
	// 	return readProvider[nicrudns.Provider](data)
	// FIXME: njalla is not supported yet: https://github.com/libdns/njalla/issues/1
	// case "njalla":
	// 	return readProvider[njalla.Provider](data)
	// FIXME: openstack-designate is not supported yet: https://github.com/libdns/openstack-designate/issues/2
	// case "openstack-designate":
	// 	return readProvider[openstack.Provider](data)
	case "ovh":
		return readProvider[ovh.Provider](data)
	case "porkbun":
		return readProvider[porkbun.Provider](data)
	// FIXME: powerdns is not supported yet: https://github.com/libdns/powerdns/issues/10
	// case "powerdns":
	// 	return readProvider[powerdns.Provider](data)
	// FIXME: regery is not supported yet.
	// case "regery":
	// 	return readProvider[regery.Provider](data)
	case "regfish":
		return readProvider[regfish.Provider](data)
	case "rfc2136":
		return readProvider[rfc2136.Provider](data)
	// FIXME: route53 is not supported yet: https://github.com/libdns/route53/issues/285
	// case "route53":
	// 	return readProvider[route53.Provider](data)
	case "scaleway":
		return readProvider[scaleway.Provider](data)
	// FIXME: timeweb is not supported yet.
	// case "selectel":
	// 	return readProvider[selectel.Provider](data)
	case "simplydotcom":
		return readProvider[simplydotcom.Provider](data)
	case "tencentcloud":
		return readProvider[tencentcloud.Provider](data)
	// FIXME: timeweb is not supported yet.
	// case "timeweb":
	// 	return readProvider[timeweb.Provider](data)
	// FIXME: totaluptime is not supported yet.
	// case "totaluptime":
	// 	return readProvider[totaluptime.Provider](data)
	// FIXME: transip is not supported yet: https://github.com/libdns/transip/pull/10
	// case "transip":
	// 	return readProvider[transip.Provider](data)
	// FIXME: vercel is not supported yet: https://github.com/libdns/vercel/issues/2
	// case "vercel":
	// 	return readProvider[vercel.Provider](data)
	// FIXME: vultr is not supported yet: https://github.com/libdns/vultr/issues/4
	// case "vultr":
	// 	return readProvider[vultr.Provider](data)
	case "westcn":
		return readProvider[westcn.Provider](data)
	default:
		return nil, fmt.Errorf("unknown provider %q", providerName)
	}
}

func readProvider[T any](data json.RawMessage) (*T, error) {
	provider := new(T)
	if err := json.Unmarshal(data, provider); err != nil {
		return nil, fmt.Errorf("parse provider settings: %w", err)
	}
	return provider, nil
}
