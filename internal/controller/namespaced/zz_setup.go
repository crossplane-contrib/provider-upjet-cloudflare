// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accessrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/access/accessrule"
	account "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/account"
	accounttoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/accounttoken"
	apitoken "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/apitoken"
	member "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/account/member"
	apishield "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishield"
	apishieldoperation "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldoperation"
	apishieldoperationschemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldoperationschemavalidationsettings"
	apishieldschema "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldschema"
	apishieldschemavalidationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/apishield/apishieldschemavalidationsettings"
	magicwangretunnel "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicwangretunnel"
	magicwanipsectunnel "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicwanipsectunnel"
	magicwanstaticroute "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/magicwanstaticroute"
	regionaltieredcache "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/regionaltieredcache"
	tieredcache "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/argo/tieredcache"
	authenticatedoriginpulls "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/authenticatedoriginpulls"
	authenticatedoriginpullscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/authenticatedoriginpullscertificate"
	certificatepack "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/certificatepack"
	customhostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/customhostname"
	customhostnamefallbackorigin "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/customhostnamefallbackorigin"
	customssl "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/customssl"
	hostnametlssetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/hostnametlssetting"
	keylesscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/keylesscertificate"
	mtlscertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/mtlscertificate"
	origincacertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/origincacertificate"
	totaltls "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/certificate/totaltls"
	dnsfirewall "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnsfirewall"
	dnsrecord "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/dns/dnsrecord"
	emailroutingaddress "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingaddress"
	emailroutingcatchall "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingcatchall"
	emailroutingdns "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingdns"
	emailroutingrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingrule"
	emailroutingsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/email/emailroutingsettings"
	botmanagement "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/botmanagement"
	filter "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/filter"
	firewallrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/firewallrule"
	leakedcredentialcheck "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/leakedcredentialcheck"
	leakedcredentialcheckrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/leakedcredentialcheckrule"
	pagerule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/pagerule"
	ratelimit "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/ratelimit"
	ruleset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/ruleset"
	urlnormalizationsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/urlnormalizationsettings"
	useragentblockingrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/firewall/useragentblockingrule"
	healthcheck "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/healthcheck/healthcheck"
	list "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/list/list"
	listitem "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/list/listitem"
	loadbalancer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/loadbalancer/loadbalancer"
	loadbalancermonitor "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/loadbalancer/loadbalancermonitor"
	loadbalancerpool "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/loadbalancer/loadbalancerpool"
	logpullretention "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/logpush/logpullretention"
	logpushjob "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/logpush/logpushjob"
	addressmap "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/addressmap"
	byoipprefix "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/byoipprefix"
	callssfuapp "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/callssfuapp"
	callsturnapp "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/callsturnapp"
	cloudconnectorrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudconnectorrules"
	cloudforceonerequest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequest"
	cloudforceonerequestasset "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequestasset"
	cloudforceonerequestmessage "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequestmessage"
	cloudforceonerequestpriority "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/cloudforceonerequestpriority"
	contentscanning "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/contentscanning"
	contentscanningexpression "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/contentscanningexpression"
	d1database "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/d1database"
	hyperdriveconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/hyperdriveconfig"
	observatoryscheduledtest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/observatoryscheduledtest"
	queue "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/queue"
	queueconsumer "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/queueconsumer"
	regionalhostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/regionalhostname"
	snippet "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/snippet"
	snippetrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/snippetrules"
	turnstilewidget "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/turnstilewidget"
	web3hostname "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/web3hostname"
	webanalyticsrule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/webanalyticsrule"
	webanalyticssite "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/misc/webanalyticssite"
	notificationpolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/notification/notificationpolicy"
	notificationpolicywebhooks "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/notification/notificationpolicywebhooks"
	providerconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/providerconfig"
	r2bucket "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucket"
	r2bucketcors "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketcors"
	r2bucketlifecycle "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketlifecycle"
	r2bucketlock "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketlock"
	r2bucketsippy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2bucketsippy"
	r2customdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2customdomain"
	r2manageddomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/r2/r2manageddomain"
	spectrumapplication "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/spectrum/spectrumapplication"
	stream "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/stream"
	streamaudiotrack "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamaudiotrack"
	streamdownload "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamdownload"
	streamkey "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamkey"
	streamliveinput "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamliveinput"
	streamwatermark "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamwatermark"
	streamwebhook "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/stream/streamwebhook"
	waitingroom "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroom"
	waitingroomevent "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroomevent"
	waitingroomrules "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroomrules"
	waitingroomsettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/waitingroom/waitingroomsettings"
	pagesdomain "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/pagesdomain"
	pagesproject "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/pagesproject"
	workerskv "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerskv"
	workerskvnamespace "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workerskvnamespace"
	workersscript "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/workers/workersscript"
	zerotrustaccessshortlivedcertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustaccessshortlivedcertificate"
	zerotrustdevicemanagednetworks "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicemanagednetworks"
	zerotrustdevicepostureintegration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdevicepostureintegration"
	zerotrustdeviceposturerule "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdeviceposturerule"
	zerotrustdextest "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdextest"
	zerotrustdlpcustomentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpcustomentry"
	zerotrustdlpcustomprofile "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpcustomprofile"
	zerotrustdlpentry "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustdlpentry"
	zerotrustgatewaycertificate "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaycertificate"
	zerotrustgatewaypolicy "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaypolicy"
	zerotrustgatewaysettings "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustgatewaysettings"
	zerotrustlist "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustlist"
	zerotrustriskbehavior "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustriskbehavior"
	zerotrustriskscoringintegration "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrustriskscoringintegration"
	zerotrusttunnelcloudflared "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelcloudflared"
	zerotrusttunnelcloudflaredconfig "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zerotrust/zerotrusttunnelcloudflaredconfig"
	zone "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zone"
	zonecachereserve "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonecachereserve"
	zonecachevariants "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonecachevariants"
	zonednssec "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonednssec"
	zonehold "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonehold"
	zonelockdown "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonelockdown"
	zonesetting "github.com/crossplane-contrib/provider-upjet-cloudflare/internal/controller/namespaced/zone/zonesetting"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessrule.Setup,
		account.Setup,
		accounttoken.Setup,
		apitoken.Setup,
		member.Setup,
		apishield.Setup,
		apishieldoperation.Setup,
		apishieldoperationschemavalidationsettings.Setup,
		apishieldschema.Setup,
		apishieldschemavalidationsettings.Setup,
		magicwangretunnel.Setup,
		magicwanipsectunnel.Setup,
		magicwanstaticroute.Setup,
		regionaltieredcache.Setup,
		tieredcache.Setup,
		authenticatedoriginpulls.Setup,
		authenticatedoriginpullscertificate.Setup,
		certificatepack.Setup,
		customhostname.Setup,
		customhostnamefallbackorigin.Setup,
		customssl.Setup,
		hostnametlssetting.Setup,
		keylesscertificate.Setup,
		mtlscertificate.Setup,
		origincacertificate.Setup,
		totaltls.Setup,
		dnsfirewall.Setup,
		dnsrecord.Setup,
		emailroutingaddress.Setup,
		emailroutingcatchall.Setup,
		emailroutingdns.Setup,
		emailroutingrule.Setup,
		emailroutingsettings.Setup,
		botmanagement.Setup,
		filter.Setup,
		firewallrule.Setup,
		leakedcredentialcheck.Setup,
		leakedcredentialcheckrule.Setup,
		pagerule.Setup,
		ratelimit.Setup,
		ruleset.Setup,
		urlnormalizationsettings.Setup,
		useragentblockingrule.Setup,
		healthcheck.Setup,
		list.Setup,
		listitem.Setup,
		loadbalancer.Setup,
		loadbalancermonitor.Setup,
		loadbalancerpool.Setup,
		logpullretention.Setup,
		logpushjob.Setup,
		addressmap.Setup,
		byoipprefix.Setup,
		callssfuapp.Setup,
		callsturnapp.Setup,
		cloudconnectorrules.Setup,
		cloudforceonerequest.Setup,
		cloudforceonerequestasset.Setup,
		cloudforceonerequestmessage.Setup,
		cloudforceonerequestpriority.Setup,
		contentscanning.Setup,
		contentscanningexpression.Setup,
		d1database.Setup,
		hyperdriveconfig.Setup,
		observatoryscheduledtest.Setup,
		queue.Setup,
		queueconsumer.Setup,
		regionalhostname.Setup,
		snippet.Setup,
		snippetrules.Setup,
		turnstilewidget.Setup,
		web3hostname.Setup,
		webanalyticsrule.Setup,
		webanalyticssite.Setup,
		notificationpolicy.Setup,
		notificationpolicywebhooks.Setup,
		providerconfig.Setup,
		r2bucket.Setup,
		r2bucketcors.Setup,
		r2bucketlifecycle.Setup,
		r2bucketlock.Setup,
		r2bucketsippy.Setup,
		r2customdomain.Setup,
		r2manageddomain.Setup,
		spectrumapplication.Setup,
		stream.Setup,
		streamaudiotrack.Setup,
		streamdownload.Setup,
		streamkey.Setup,
		streamliveinput.Setup,
		streamwatermark.Setup,
		streamwebhook.Setup,
		waitingroom.Setup,
		waitingroomevent.Setup,
		waitingroomrules.Setup,
		waitingroomsettings.Setup,
		pagesdomain.Setup,
		pagesproject.Setup,
		workerskv.Setup,
		workerskvnamespace.Setup,
		workersscript.Setup,
		zerotrustaccessshortlivedcertificate.Setup,
		zerotrustdevicemanagednetworks.Setup,
		zerotrustdevicepostureintegration.Setup,
		zerotrustdeviceposturerule.Setup,
		zerotrustdextest.Setup,
		zerotrustdlpcustomentry.Setup,
		zerotrustdlpcustomprofile.Setup,
		zerotrustdlpentry.Setup,
		zerotrustgatewaycertificate.Setup,
		zerotrustgatewaypolicy.Setup,
		zerotrustgatewaysettings.Setup,
		zerotrustlist.Setup,
		zerotrustriskbehavior.Setup,
		zerotrustriskscoringintegration.Setup,
		zerotrusttunnelcloudflared.Setup,
		zerotrusttunnelcloudflaredconfig.Setup,
		zone.Setup,
		zonecachereserve.Setup,
		zonecachevariants.Setup,
		zonednssec.Setup,
		zonehold.Setup,
		zonelockdown.Setup,
		zonesetting.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessrule.SetupGated,
		account.SetupGated,
		accounttoken.SetupGated,
		apitoken.SetupGated,
		member.SetupGated,
		apishield.SetupGated,
		apishieldoperation.SetupGated,
		apishieldoperationschemavalidationsettings.SetupGated,
		apishieldschema.SetupGated,
		apishieldschemavalidationsettings.SetupGated,
		magicwangretunnel.SetupGated,
		magicwanipsectunnel.SetupGated,
		magicwanstaticroute.SetupGated,
		regionaltieredcache.SetupGated,
		tieredcache.SetupGated,
		authenticatedoriginpulls.SetupGated,
		authenticatedoriginpullscertificate.SetupGated,
		certificatepack.SetupGated,
		customhostname.SetupGated,
		customhostnamefallbackorigin.SetupGated,
		customssl.SetupGated,
		hostnametlssetting.SetupGated,
		keylesscertificate.SetupGated,
		mtlscertificate.SetupGated,
		origincacertificate.SetupGated,
		totaltls.SetupGated,
		dnsfirewall.SetupGated,
		dnsrecord.SetupGated,
		emailroutingaddress.SetupGated,
		emailroutingcatchall.SetupGated,
		emailroutingdns.SetupGated,
		emailroutingrule.SetupGated,
		emailroutingsettings.SetupGated,
		botmanagement.SetupGated,
		filter.SetupGated,
		firewallrule.SetupGated,
		leakedcredentialcheck.SetupGated,
		leakedcredentialcheckrule.SetupGated,
		pagerule.SetupGated,
		ratelimit.SetupGated,
		ruleset.SetupGated,
		urlnormalizationsettings.SetupGated,
		useragentblockingrule.SetupGated,
		healthcheck.SetupGated,
		list.SetupGated,
		listitem.SetupGated,
		loadbalancer.SetupGated,
		loadbalancermonitor.SetupGated,
		loadbalancerpool.SetupGated,
		logpullretention.SetupGated,
		logpushjob.SetupGated,
		addressmap.SetupGated,
		byoipprefix.SetupGated,
		callssfuapp.SetupGated,
		callsturnapp.SetupGated,
		cloudconnectorrules.SetupGated,
		cloudforceonerequest.SetupGated,
		cloudforceonerequestasset.SetupGated,
		cloudforceonerequestmessage.SetupGated,
		cloudforceonerequestpriority.SetupGated,
		contentscanning.SetupGated,
		contentscanningexpression.SetupGated,
		d1database.SetupGated,
		hyperdriveconfig.SetupGated,
		observatoryscheduledtest.SetupGated,
		queue.SetupGated,
		queueconsumer.SetupGated,
		regionalhostname.SetupGated,
		snippet.SetupGated,
		snippetrules.SetupGated,
		turnstilewidget.SetupGated,
		web3hostname.SetupGated,
		webanalyticsrule.SetupGated,
		webanalyticssite.SetupGated,
		notificationpolicy.SetupGated,
		notificationpolicywebhooks.SetupGated,
		providerconfig.SetupGated,
		r2bucket.SetupGated,
		r2bucketcors.SetupGated,
		r2bucketlifecycle.SetupGated,
		r2bucketlock.SetupGated,
		r2bucketsippy.SetupGated,
		r2customdomain.SetupGated,
		r2manageddomain.SetupGated,
		spectrumapplication.SetupGated,
		stream.SetupGated,
		streamaudiotrack.SetupGated,
		streamdownload.SetupGated,
		streamkey.SetupGated,
		streamliveinput.SetupGated,
		streamwatermark.SetupGated,
		streamwebhook.SetupGated,
		waitingroom.SetupGated,
		waitingroomevent.SetupGated,
		waitingroomrules.SetupGated,
		waitingroomsettings.SetupGated,
		pagesdomain.SetupGated,
		pagesproject.SetupGated,
		workerskv.SetupGated,
		workerskvnamespace.SetupGated,
		workersscript.SetupGated,
		zerotrustaccessshortlivedcertificate.SetupGated,
		zerotrustdevicemanagednetworks.SetupGated,
		zerotrustdevicepostureintegration.SetupGated,
		zerotrustdeviceposturerule.SetupGated,
		zerotrustdextest.SetupGated,
		zerotrustdlpcustomentry.SetupGated,
		zerotrustdlpcustomprofile.SetupGated,
		zerotrustdlpentry.SetupGated,
		zerotrustgatewaycertificate.SetupGated,
		zerotrustgatewaypolicy.SetupGated,
		zerotrustgatewaysettings.SetupGated,
		zerotrustlist.SetupGated,
		zerotrustriskbehavior.SetupGated,
		zerotrustriskscoringintegration.SetupGated,
		zerotrusttunnelcloudflared.SetupGated,
		zerotrusttunnelcloudflaredconfig.SetupGated,
		zone.SetupGated,
		zonecachereserve.SetupGated,
		zonecachevariants.SetupGated,
		zonednssec.SetupGated,
		zonehold.SetupGated,
		zonelockdown.SetupGated,
		zonesetting.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		accessrule.SetupWebhookWithManager,
		account.SetupWebhookWithManager,
		accounttoken.SetupWebhookWithManager,
		apitoken.SetupWebhookWithManager,
		member.SetupWebhookWithManager,
		apishield.SetupWebhookWithManager,
		apishieldoperation.SetupWebhookWithManager,
		apishieldoperationschemavalidationsettings.SetupWebhookWithManager,
		apishieldschema.SetupWebhookWithManager,
		apishieldschemavalidationsettings.SetupWebhookWithManager,
		magicwangretunnel.SetupWebhookWithManager,
		magicwanipsectunnel.SetupWebhookWithManager,
		magicwanstaticroute.SetupWebhookWithManager,
		regionaltieredcache.SetupWebhookWithManager,
		tieredcache.SetupWebhookWithManager,
		authenticatedoriginpulls.SetupWebhookWithManager,
		authenticatedoriginpullscertificate.SetupWebhookWithManager,
		certificatepack.SetupWebhookWithManager,
		customhostname.SetupWebhookWithManager,
		customhostnamefallbackorigin.SetupWebhookWithManager,
		customssl.SetupWebhookWithManager,
		hostnametlssetting.SetupWebhookWithManager,
		keylesscertificate.SetupWebhookWithManager,
		mtlscertificate.SetupWebhookWithManager,
		origincacertificate.SetupWebhookWithManager,
		totaltls.SetupWebhookWithManager,
		dnsfirewall.SetupWebhookWithManager,
		dnsrecord.SetupWebhookWithManager,
		emailroutingaddress.SetupWebhookWithManager,
		emailroutingcatchall.SetupWebhookWithManager,
		emailroutingdns.SetupWebhookWithManager,
		emailroutingrule.SetupWebhookWithManager,
		emailroutingsettings.SetupWebhookWithManager,
		botmanagement.SetupWebhookWithManager,
		filter.SetupWebhookWithManager,
		firewallrule.SetupWebhookWithManager,
		leakedcredentialcheck.SetupWebhookWithManager,
		leakedcredentialcheckrule.SetupWebhookWithManager,
		pagerule.SetupWebhookWithManager,
		ratelimit.SetupWebhookWithManager,
		ruleset.SetupWebhookWithManager,
		urlnormalizationsettings.SetupWebhookWithManager,
		useragentblockingrule.SetupWebhookWithManager,
		healthcheck.SetupWebhookWithManager,
		list.SetupWebhookWithManager,
		listitem.SetupWebhookWithManager,
		loadbalancer.SetupWebhookWithManager,
		loadbalancermonitor.SetupWebhookWithManager,
		loadbalancerpool.SetupWebhookWithManager,
		logpullretention.SetupWebhookWithManager,
		logpushjob.SetupWebhookWithManager,
		addressmap.SetupWebhookWithManager,
		byoipprefix.SetupWebhookWithManager,
		callssfuapp.SetupWebhookWithManager,
		callsturnapp.SetupWebhookWithManager,
		cloudconnectorrules.SetupWebhookWithManager,
		cloudforceonerequest.SetupWebhookWithManager,
		cloudforceonerequestasset.SetupWebhookWithManager,
		cloudforceonerequestmessage.SetupWebhookWithManager,
		cloudforceonerequestpriority.SetupWebhookWithManager,
		contentscanning.SetupWebhookWithManager,
		contentscanningexpression.SetupWebhookWithManager,
		d1database.SetupWebhookWithManager,
		hyperdriveconfig.SetupWebhookWithManager,
		observatoryscheduledtest.SetupWebhookWithManager,
		queue.SetupWebhookWithManager,
		queueconsumer.SetupWebhookWithManager,
		regionalhostname.SetupWebhookWithManager,
		snippet.SetupWebhookWithManager,
		snippetrules.SetupWebhookWithManager,
		turnstilewidget.SetupWebhookWithManager,
		web3hostname.SetupWebhookWithManager,
		webanalyticsrule.SetupWebhookWithManager,
		webanalyticssite.SetupWebhookWithManager,
		notificationpolicy.SetupWebhookWithManager,
		notificationpolicywebhooks.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		r2bucket.SetupWebhookWithManager,
		r2bucketcors.SetupWebhookWithManager,
		r2bucketlifecycle.SetupWebhookWithManager,
		r2bucketlock.SetupWebhookWithManager,
		r2bucketsippy.SetupWebhookWithManager,
		r2customdomain.SetupWebhookWithManager,
		r2manageddomain.SetupWebhookWithManager,
		spectrumapplication.SetupWebhookWithManager,
		stream.SetupWebhookWithManager,
		streamaudiotrack.SetupWebhookWithManager,
		streamdownload.SetupWebhookWithManager,
		streamkey.SetupWebhookWithManager,
		streamliveinput.SetupWebhookWithManager,
		streamwatermark.SetupWebhookWithManager,
		streamwebhook.SetupWebhookWithManager,
		waitingroom.SetupWebhookWithManager,
		waitingroomevent.SetupWebhookWithManager,
		waitingroomrules.SetupWebhookWithManager,
		waitingroomsettings.SetupWebhookWithManager,
		pagesdomain.SetupWebhookWithManager,
		pagesproject.SetupWebhookWithManager,
		workerskv.SetupWebhookWithManager,
		workerskvnamespace.SetupWebhookWithManager,
		workersscript.SetupWebhookWithManager,
		zerotrustaccessshortlivedcertificate.SetupWebhookWithManager,
		zerotrustdevicemanagednetworks.SetupWebhookWithManager,
		zerotrustdevicepostureintegration.SetupWebhookWithManager,
		zerotrustdeviceposturerule.SetupWebhookWithManager,
		zerotrustdextest.SetupWebhookWithManager,
		zerotrustdlpcustomentry.SetupWebhookWithManager,
		zerotrustdlpcustomprofile.SetupWebhookWithManager,
		zerotrustdlpentry.SetupWebhookWithManager,
		zerotrustgatewaycertificate.SetupWebhookWithManager,
		zerotrustgatewaypolicy.SetupWebhookWithManager,
		zerotrustgatewaysettings.SetupWebhookWithManager,
		zerotrustlist.SetupWebhookWithManager,
		zerotrustriskbehavior.SetupWebhookWithManager,
		zerotrustriskscoringintegration.SetupWebhookWithManager,
		zerotrusttunnelcloudflared.SetupWebhookWithManager,
		zerotrusttunnelcloudflaredconfig.SetupWebhookWithManager,
		zone.SetupWebhookWithManager,
		zonecachereserve.SetupWebhookWithManager,
		zonecachevariants.SetupWebhookWithManager,
		zonednssec.SetupWebhookWithManager,
		zonehold.SetupWebhookWithManager,
		zonelockdown.SetupWebhookWithManager,
		zonesetting.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
