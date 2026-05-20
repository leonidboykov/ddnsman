package ddnsman

import (
	"encoding/json"
	"fmt"

	"github.com/libdns/libdns"

	"github.com/libdns/acmedns"
	"github.com/libdns/acmeproxy"
	"github.com/libdns/alidns"
	allinkl "github.com/libdns/all-inkl"
	"github.com/libdns/arvancloud"
	"github.com/libdns/autodns"
	"github.com/libdns/azure"
	"github.com/libdns/bluecat"
	"github.com/libdns/bunny"
	"github.com/libdns/cloudflare"
	"github.com/libdns/cloudns"
	conohav3 "github.com/libdns/conoha"
	"github.com/libdns/desec"
	"github.com/libdns/digitalocean"
	"github.com/libdns/dinahosting/v2"
	"github.com/libdns/directadmin"
	"github.com/libdns/dnsexit"
	"github.com/libdns/dnsimple"
	"github.com/libdns/dnsupdate"
	"github.com/libdns/domainnameshop"
	"github.com/libdns/dreamhost"
	"github.com/libdns/duckdns"
	"github.com/libdns/dynu"
	"github.com/libdns/dynv6"
	"github.com/libdns/easydns"
	"github.com/libdns/edgeone"
	"github.com/libdns/gandi"
	"github.com/libdns/gcore"
	"github.com/libdns/glesys"
	"github.com/libdns/godaddy"
	"github.com/libdns/googleclouddns"
	"github.com/libdns/he"
	"github.com/libdns/hetzner"
	"github.com/libdns/httpnet"
	"github.com/libdns/huaweicloud"
	"github.com/libdns/infomaniak"
	"github.com/libdns/inwx"
	"github.com/libdns/ionos"
	"github.com/libdns/linode"
	"github.com/libdns/loopia"
	"github.com/libdns/luadns"
	"github.com/libdns/mailinabox"
	"github.com/libdns/metaname"
	"github.com/libdns/mijnhost"
	"github.com/libdns/mythicbeasts"
	"github.com/libdns/namecheap"
	"github.com/libdns/namesilo"
	"github.com/libdns/netcup"
	"github.com/libdns/netlify"
	"github.com/libdns/netnod"
	"github.com/libdns/nfsn"
	"github.com/libdns/njalla"
	"github.com/libdns/oraclecloud"
	"github.com/libdns/ovh"
	"github.com/libdns/porkbun"
	"github.com/libdns/powerdns"
	"github.com/libdns/pph"
	regery "github.com/libdns/regery"
	"github.com/libdns/regfish"
	"github.com/libdns/rfc2136"
	"github.com/libdns/route53"
	"github.com/libdns/scaleway"
	"github.com/libdns/servercow"
	"github.com/libdns/simplydotcom"
	spaceship "github.com/libdns/spaceship"
	"github.com/libdns/tecnocratica"
	"github.com/libdns/tencentcloud"
	"github.com/libdns/thelittlehost"
	"github.com/libdns/timeweb"
	"github.com/libdns/transip"
	"github.com/libdns/unifi"
	"github.com/libdns/volcengine"
	"github.com/libdns/vultr/v2"
	"github.com/libdns/wedos"
	"github.com/libdns/westcn"
)

// Providers allows get and set records to DNS provider.
type Provider interface {
	libdns.RecordGetter
	libdns.RecordSetter
}

func newProvider(providerName string, data json.RawMessage) (Provider, error) {
	// Missing providers:
	// - civo – not compatible with current libdns, public archive.
	// - ddnss – not compatible with current libdns.
	// - dode – missing [libdns.RecordGetter]/[libdns.RecordSetter] support.
	// - nanelo – missing [libdns.RecordGetter]/[libdns.RecordSetter] support.
	// - websupport – missing [libdns.RecordSetter] support.

	switch providerName {
	case "acmedns":
		return readProvider[acmedns.Provider](data)
	case "acmeproxy":
		return readProvider[acmeproxy.Provider](data)
	case "alidns":
		return readProvider[alidns.Provider](data)
	case "all-inkl":
		return readProvider[allinkl.Provider](data)
	case "arvancloud":
		return readProvider[arvancloud.Provider](data)
	case "autodns":
		return readProvider[autodns.Provider](data)
	case "azure":
		return readProvider[azure.Provider](data)
	case "bluecat":
		return readProvider[bluecat.Provider](data)
	case "bunny":
		return readProvider[bunny.Provider](data)
	case "cloudflare":
		return readProvider[cloudflare.Provider](data)
	case "cloudns":
		return readProvider[cloudns.Provider](data)
	case "cohona":
		return readProvider[conohav3.Provider](data)
	case "desec":
		return readProvider[desec.Provider](data)
	case "digitalocean":
		return readProvider[digitalocean.Provider](data)
	case "dinahosting":
		return readProvider[dinahosting.Provider](data)
	case "directadmin":
		return readProvider[directadmin.Provider](data)
	case "dnsexit":
		return readProvider[dnsexit.Provider](data)
	case "dnsimple":
		return readProvider[dnsimple.Provider](data)
	// FIXME: dnsmadeeasy is not supported yet: https://github.com/libdns/dnsmadeeasy/issues/8
	// case "dnsmadeeasy":
	// 	return readProvider[dnsmadeeasy.Provider](data)
	// FIXME: dnspod is not supported yet: https://github.com/libdns/dnspod/pull/9
	// case "dnspod":
	// 	return readProvider[dnspod.Provider](data)
	case "dnsupdate":
		return readProvider[dnsupdate.Provider](data)
	case "domainnameshop":
		return readProvider[domainnameshop.Provider](data)
	case "dreamhost":
		return readProvider[dreamhost.Provider](data)
	case "duckdns":
		return readProvider[duckdns.Provider](data)
	case "dynu":
		return readProvider[dynu.Provider](data)
	case "dynv6":
		return readProvider[dynv6.Provider](data)
	case "easydns":
		return readProvider[easydns.Provider](data)
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
	case "godaddy":
		return readProvider[godaddy.Provider](data)
	case "googleclouddns":
		return readProvider[googleclouddns.Provider](data)
	case "he":
		return readProvider[he.Provider](data)
	case "hetzner":
		return readProvider[hetzner.Provider](data)
	// FIXME: hexonet is not supported yet: https://github.com/libdns/hexonet/issues/1
	// case "hexonet":
	// 	return readProvider[hexonet.Provider](data)
	// FIXME: hosttech is not supported yet: https://github.com/libdns/hosttech/pull/20
	// case "hosttech":
	// 	return readProvider[hosttech.Provider](data)
	case "httpnet":
		return readProvider[httpnet.Provider](data)
	case "huaweicloud":
		return readProvider[huaweicloud.Provider](data)
	case "infomaniak":
		return readProvider[infomaniak.Provider](data)
	case "inwx":
		return readProvider[inwx.Provider](data)
	case "ionos":
		return readProvider[ionos.Provider](data)
	// FIXME: katapult is not supported yet: https://github.com/libdns/katapult/pull/1
	// case "katapult":
	// 	return readProvider[katapult.Provider](data)
	// FIXME: leaseweb is not supported yet.
	// case "leaseweb":
	// 	return readProvider[leaseweb.Provider](data)
	case "linode":
		return readProvider[linode.Provider](data)
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
	case "mythicbeasts":
		return readProvider[mythicbeasts.Provider](data)
	case "namecheap":
		return readProvider[namecheap.Provider](data)
	// FIXME: namedotcom is not supported yet: https://github.com/libdns/namedotcom/pull/5
	// case "namedotcom":
	// 	return readProvider[namedotcom.Provider](data)
	case "namesilo":
		return readProvider[namesilo.Provider](data)
	// FIXME: neoserv is not supported yet.
	// case "neoserv":
	// 	return readProvider[neoserv.Provider](data)
	case "netcup":
		return readProvider[netcup.Provider](data)
	case "netlify":
		return readProvider[netlify.Provider](data)
	case "netnod":
		return readProvider[netnod.Provider](data)
	case "nfsn":
		return readProvider[nfsn.Provider](data)
	// FIXME: nicrudns is not supported yet.
	// case "nicrudns":
	// 	return readProvider[nicrudns.Provider](data)
	case "njalla":
		return readProvider[njalla.Provider](data)
	// FIXME: openstack-designate is not supported yet: https://github.com/libdns/openstack-designate/issues/2
	// case "openstack-designate":
	// 	return readProvider[openstack.Provider](data)
	case "oraclecloud":
		return readProvider[oraclecloud.Provider](data)
	case "ovh":
		return readProvider[ovh.Provider](data)
	case "porkbun":
		return readProvider[porkbun.Provider](data)
	case "powerdns":
		return readProvider[powerdns.Provider](data)
	case "pph":
		return readProvider[pph.Provider](data)
	case "regery":
		return readProvider[regery.Provider](data)
	case "regfish":
		return readProvider[regfish.Provider](data)
	case "rfc2136":
		return readProvider[rfc2136.Provider](data)
	case "route53":
		return readProvider[route53.Provider](data)
	case "scaleway":
		return readProvider[scaleway.Provider](data)
	// FIXME: selectel is not supported yet.
	// case "selectel":
	// 	return readProvider[selectel.Provider](data)
	case "servercow":
		return readProvider[servercow.Provider](data)
	case "simplydotcom":
		return readProvider[simplydotcom.Provider](data)
	case "spaceship":
		return readProvider[spaceship.Provider](data)
	case "tecnocratica":
		return readProvider[tecnocratica.Provider](data)
	case "tencentcloud":
		return readProvider[tencentcloud.Provider](data)
	case "thelittlehost":
		return readProvider[thelittlehost.Provider](data)
	case "timeweb":
		return readProvider[timeweb.Provider](data)
	// FIXME: totaluptime is not supported yet.
	// case "totaluptime":
	// 	return readProvider[totaluptime.Provider](data)
	case "transip":
		return readProvider[transip.Provider](data)
	case "unifi":
		return readProvider[unifi.Provider](data)
	// FIXME: vercel is not supported yet: https://github.com/libdns/vercel/issues/2
	// case "vercel":
	// 	return readProvider[vercel.Provider](data)
	case "volcengine":
		return readProvider[volcengine.Provider](data)
	case "vultr":
		return readProvider[vultr.Provider](data)
	case "wedos":
		return readProvider[wedos.Provider](data)
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
