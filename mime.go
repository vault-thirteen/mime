////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package mime

// Common MIME Types.

// Notes.
//
// This Packages contains all IANA registered MIME Types which have an RFC Reference as of 2019-05-26.
// For more Information visit this URL:
// https://www.iana.org/assignments/media-types/media-types.xhtml
// This List may contain some old deprecated and obsolete MIME Types.

// Application.
const (
	TypeApplication1Dinterleavedparityfec             = "application/1d-interleaved-parityfec"                 // [RFC6015]
	TypeApplicationAltocostmapjson                    = "application/alto-costmap+json"                        // [RFC7285]
	TypeApplicationAltocostmapfilterjson              = "application/alto-costmapfilter+json"                  // [RFC7285]
	TypeApplicationAltodirectoryjson                  = "application/alto-directory+json"                      // [RFC7285]
	TypeApplicationAltoendpointpropjson               = "application/alto-endpointprop+json"                   // [RFC7285]
	TypeApplicationAltoendpointpropparamsjson         = "application/alto-endpointpropparams+json"             // [RFC7285]
	TypeApplicationAltoendpointcostjson               = "application/alto-endpointcost+json"                   // [RFC7285]
	TypeApplicationAltoendpointcostparamsjson         = "application/alto-endpointcostparams+json"             // [RFC7285]
	TypeApplicationAltoerrorjson                      = "application/alto-error+json"                          // [RFC7285]
	TypeApplicationAltonetworkmapfilterjson           = "application/alto-networkmapfilter+json"               // [RFC7285]
	TypeApplicationAltonetworkmapjson                 = "application/alto-networkmap+json"                     // [RFC7285]
	TypeApplicationAtomxml                            = "application/atom+xml"                                 // [RFC4287][RFC5023]
	TypeApplicationAtomcatxml                         = "application/atomcat+xml"                              // [RFC5023]
	TypeApplicationAtomdeletedxml                     = "application/atomdeleted+xml"                          // [RFC6721]
	TypeApplicationAtomsvcxml                         = "application/atomsvc+xml"                              // [RFC5023]
	TypeApplicationAuthpolicyxml                      = "application/auth-policy+xml"                          // [RFC4745]
	TypeApplicationBatchsmtp                          = "application/batch-SMTP"                               // [RFC2442]
	TypeApplicationBeepxml                            = "application/beep+xml"                                 // [RFC3080]
	TypeApplicationCalendarjson                       = "application/calendar+json"                            // [RFC7265]
	TypeApplicationCalendarxml                        = "application/calendar+xml"                             // [RFC6321]
	TypeApplicationCallcompletion                     = "application/call-completion"                          // [RFC6910]
	TypeApplicationCals1840                           = "application/CALS-1840"                                // [RFC1895]
	TypeApplicationCbor                               = "application/cbor"                                     // [RFC7049]
	TypeApplicationCcmpxml                            = "application/ccmp+xml"                                 // [RFC6503]
	TypeApplicationCcxmlxml                           = "application/ccxml+xml"                                // [RFC4267]
	TypeApplicationCdmicapability                     = "application/cdmi-capability"                          // [RFC6208]
	TypeApplicationCdmicontainer                      = "application/cdmi-container"                           // [RFC6208]
	TypeApplicationCdmidomain                         = "application/cdmi-domain"                              // [RFC6208]
	TypeApplicationCdmiobject                         = "application/cdmi-object"                              // [RFC6208]
	TypeApplicationCdmiqueue                          = "application/cdmi-queue"                               // [RFC6208]
	TypeApplicationCdni                               = "application/cdni"                                     // [RFC7736]
	TypeApplicationCellmlxml                          = "application/cellml+xml"                               // [RFC4708]
	TypeApplicationCfw                                = "application/cfw"                                      // [RFC6230]
	TypeApplicationClueinfoxml                        = "application/clue_info+xml"                            // [RFC-ietf-clue-data-model-schema-17]
	TypeApplicationCms                                = "application/cms"                                      // [RFC7193]
	TypeApplicationCnrpxml                            = "application/cnrp+xml"                                 // [RFC3367]
	TypeApplicationCoapgroupjson                      = "application/coap-group+json"                          // [RFC7390]
	TypeApplicationCoappayload                        = "application/coap-payload"                             // [RFC8075]
	TypeApplicationConferenceinfoxml                  = "application/conference-info+xml"                      // [RFC4575]
	TypeApplicationCplxml                             = "application/cpl+xml"                                  // [RFC3880]
	TypeApplicationCose                               = "application/cose"                                     // [RFC8152]
	TypeApplicationCosekey                            = "application/cose-key"                                 // [RFC8152]
	TypeApplicationCosekeyset                         = "application/cose-key-set"                             // [RFC8152]
	TypeApplicationCsrattrs                           = "application/csrattrs"                                 // [RFC7030]
	TypeApplicationCwt                                = "application/cwt"                                      // [RFC8392]
	TypeApplicationDavmountxml                        = "application/davmount+xml"                             // [RFC4709]
	TypeApplicationDialoginfoxml                      = "application/dialog-info+xml"                          // [RFC4235]
	TypeApplicationDicom                              = "application/dicom"                                    // [RFC3240]
	TypeApplicationDns                                = "application/dns"                                      // [RFC4027]
	TypeApplicationDnsjson                            = "application/dns+json"                                 // [RFC8427]
	TypeApplicationDnsmessage                         = "application/dns-message"                              // [RFC8484]
	TypeApplicationDskppxml                           = "application/dskpp+xml"                                // [RFC6063]
	TypeApplicationDsscder                            = "application/dssc+der"                                 // [RFC5698]
	TypeApplicationDsscxml                            = "application/dssc+xml"                                 // [RFC5698]
	TypeApplicationDvcs                               = "application/dvcs"                                     // [RFC3029]
	TypeApplicationEcmascript                         = "application/ecmascript"                               // [RFC4329]
	TypeApplicationEdiconsent                         = "application/EDI-consent"                              // [RFC1767]
	TypeApplicationEdifact                            = "application/EDIFACT"                                  // [RFC1767]
	TypeApplicationEdix12                             = "application/EDI-X12"                                  // [RFC1767]
	TypeApplicationEmergencycalldatacommentxml        = "application/EmergencyCallData.Comment+xml"            // [RFC7852]
	TypeApplicationEmergencycalldatacontrolxml        = "application/EmergencyCallData.Control+xml"            // [RFC8147]
	TypeApplicationEmergencycalldatadeviceinfoxml     = "application/EmergencyCallData.DeviceInfo+xml"         // [RFC7852]
	TypeApplicationEmergencycalldataecallmsd          = "application/EmergencyCallData.eCall.MSD"              // [RFC8147]
	TypeApplicationEmergencycalldataproviderinfoxml   = "application/EmergencyCallData.ProviderInfo+xml"       // [RFC7852]
	TypeApplicationEmergencycalldataserviceinfoxml    = "application/EmergencyCallData.ServiceInfo+xml"        // [RFC7852]
	TypeApplicationEmergencycalldatasubscriberinfoxml = "application/EmergencyCallData.SubscriberInfo+xml"     // [RFC7852]
	TypeApplicationEmergencycalldatavedsxml           = "application/EmergencyCallData.VEDS+xml"               // [RFC8148]
	TypeApplicationEncaprtp                           = "application/encaprtp"                                 // [RFC6849]
	TypeApplicationEppxml                             = "application/epp+xml"                                  // [RFC5730]
	TypeApplicationExample                            = "application/example"                                  // [RFC4735]
	TypeApplicationExpectctreportjson                 = "application/expect-ct-report+json"                    // [RFC-ietf-httpbis-expect-ct-08]
	TypeApplicationFdtxml                             = "application/fdt+xml"                                  // [RFC6726]
	TypeApplicationFits                               = "application/fits"                                     // [RFC4047]
	TypeApplicationFlexfec                            = "application/flexfec"                                  // [RFC-ietf-payload-flexible-fec-scheme-20]
	TypeApplicationFontsfnt                           = "application/font-sfnt"                                // [Levantovsky][ISO-IEC_JTC1][RFC8081], Deprecated.
	TypeApplicationFonttdpfr                          = "application/font-tdpfr"                               // [RFC3073]
	TypeApplicationFontwoff                           = "application/font-woff"                                // [W3C][RFC8081], Deprecated.
	TypeApplicationFrameworkattributesxml             = "application/framework-attributes+xml"                 // [RFC6230]
	TypeApplicationGeojson                            = "application/geo+json"                                 // [RFC7946]
	TypeApplicationGeojsonseq                         = "application/geo+json-seq"                             // [RFC8142]
	TypeApplicationGzip                               = "application/gzip"                                     // [RFC6713]
	TypeApplicationH224                               = "application/H224"                                     // [RFC4573]
	TypeApplicationHeldxml                            = "application/held+xml"                                 // [RFC5985]
	TypeApplicationHttp                               = "application/http"                                     // [RFC7230]
	TypeApplicationIbekeyrequestxml                   = "application/ibe-key-request+xml"                      // [RFC5408]
	TypeApplicationIbepkgreplyxml                     = "application/ibe-pkg-reply+xml"                        // [RFC5408]
	TypeApplicationIbeppdata                          = "application/ibe-pp-data"                              // [RFC5408]
	TypeApplicationImiscomposingxml                   = "application/im-iscomposing+xml"                       // [RFC3994]
	TypeApplicationIndex                              = "application/index"                                    // [RFC2652]
	TypeApplicationIndexcmd                           = "application/index.cmd"                                // [RFC2652]
	TypeApplicationIndexobj                           = "application/index.obj"                                // [RFC2652]
	TypeApplicationIndexresponse                      = "application/index.response"                           // [RFC2652]
	TypeApplicationIndexvnd                           = "application/index.vnd"                                // [RFC2652]
	TypeApplicationIotp                               = "application/IOTP"                                     // [RFC2935]
	TypeApplicationIpfix                              = "application/ipfix"                                    // [RFC5655]
	TypeApplicationIpp                                = "application/ipp"                                      // [RFC8010]
	TypeApplicationIsup                               = "application/isup"                                     // [RFC3204]
	TypeApplicationJavascript                         = "application/javascript"                               // [RFC4329]
	TypeApplicationJose                               = "application/jose"                                     // [RFC7515]
	TypeApplicationJosejson                           = "application/jose+json"                                // [RFC7515]
	TypeApplicationJrdjson                            = "application/jrd+json"                                 // [RFC7033]
	TypeApplicationJson                               = "application/json"                                     // [RFC8259]
	TypeApplicationJsonpatchjson                      = "application/json-patch+json"                          // [RFC6902]
	TypeApplicationJsonseq                            = "application/json-seq"                                 // [RFC7464]
	TypeApplicationJwkjson                            = "application/jwk+json"                                 // [RFC7517]
	TypeApplicationJwksetjson                         = "application/jwk-set+json"                             // [RFC7517]
	TypeApplicationJwt                                = "application/jwt"                                      // [RFC7519]
	TypeApplicationKpmlrequestxml                     = "application/kpml-request+xml"                         // [RFC4730]
	TypeApplicationKpmlresponsexml                    = "application/kpml-response+xml"                        // [RFC4730]
	TypeApplicationLgrxml                             = "application/lgr+xml"                                  // [RFC7940]
	TypeApplicationLinkformat                         = "application/link-format"                              // [RFC6690]
	TypeApplicationLoadcontrolxml                     = "application/load-control+xml"                         // [RFC7200]
	TypeApplicationLostxml                            = "application/lost+xml"                                 // [RFC5222]
	TypeApplicationLostsyncxml                        = "application/lostsync+xml"                             // [RFC6739]
	TypeApplicationMadsxml                            = "application/mads+xml"                                 // [RFC6207]
	TypeApplicationMarc                               = "application/marc"                                     // [RFC2220]
	TypeApplicationMarcxmlxml                         = "application/marcxml+xml"                              // [RFC6207]
	TypeApplicationMbox                               = "application/mbox"                                     // [RFC4155]
	TypeApplicationMediacontrolxml                    = "application/media_control+xml"                        // [RFC5168]
	TypeApplicationMediapolicydatasetxml              = "application/media-policy-dataset+xml"                 // [RFC6796]
	TypeApplicationMediaservercontrolxml              = "application/mediaservercontrol+xml"                   // [RFC5022]
	TypeApplicationMergepatchjson                     = "application/merge-patch+json"                         // [RFC7396]
	TypeApplicationMetalink4Xml                       = "application/metalink4+xml"                            // [RFC5854]
	TypeApplicationMetsxml                            = "application/mets+xml"                                 // [RFC6207]
	TypeApplicationMikey                              = "application/mikey"                                    // [RFC3830]
	TypeApplicationModsxml                            = "application/mods+xml"                                 // [RFC6207]
	TypeApplicationMosskeys                           = "application/moss-keys"                                // [RFC1848]
	TypeApplicationMosssignature                      = "application/moss-signature"                           // [RFC1848]
	TypeApplicationMosskeydata                        = "application/mosskey-data"                             // [RFC1848]
	TypeApplicationMosskeyrequest                     = "application/mosskey-request"                          // [RFC1848]
	TypeApplicationMp21                               = "application/mp21"                                     // [RFC6381]
	TypeApplicationMp4                                = "application/mp4"                                      // [RFC4337][RFC6381]
	TypeApplicationMpeg4Generic                       = "application/mpeg4-generic"                            // [RFC3640]
	TypeApplicationMpeg4Iod                           = "application/mpeg4-iod"                                // [RFC4337]
	TypeApplicationMpeg4Iodxmt                        = "application/mpeg4-iod-xmt"                            // [RFC4337]
	TypeApplicationMrbconsumerxml                     = "application/mrb-consumer+xml"                         // [RFC6917]
	TypeApplicationMrbpublishxml                      = "application/mrb-publish+xml"                          // [RFC6917]
	TypeApplicationMscivrxml                          = "application/msc-ivr+xml"                              // [RFC6231]
	TypeApplicationMscmixerxml                        = "application/msc-mixer+xml"                            // [RFC6505]
	TypeApplicationMudjson                            = "application/mud+json"                                 // [RFC8520]
	TypeApplicationMxf                                = "application/mxf"                                      // [RFC4539]
	TypeApplicationNasdata                            = "application/nasdata"                                  // [RFC4707]
	TypeApplicationNewscheckgroups                    = "application/news-checkgroups"                         // [RFC5537]
	TypeApplicationNewsgroupinfo                      = "application/news-groupinfo"                           // [RFC5537]
	TypeApplicationNewstransmission                   = "application/news-transmission"                        // [RFC5537]
	TypeApplicationNlsmlxml                           = "application/nlsml+xml"                                // [RFC6787]
	TypeApplicationOcsprequest                        = "application/ocsp-request"                             // [RFC6960]
	TypeApplicationOcspresponse                       = "application/ocsp-response"                            // [RFC6960]
	TypeApplicationOctetstream                        = "application/octet-stream"                             // [RFC2045][RFC2046]
	TypeApplicationOda                                = "application/ODA"                                      // [RFC1494]
	TypeApplicationOebpspackagexml                    = "application/oebps-package+xml"                        // [RFC4839]
	TypeApplicationOgg                                = "application/ogg"                                      // [RFC5334][RFC7845]
	TypeApplicationOscore                             = "application/oscore"                                   // [RFC-ietf-core-object-security-16]
	TypeApplicationP2Poverlayxml                      = "application/p2p-overlay+xml"                          // [RFC6940]
	TypeApplicationParityfec                          = "application/parityfec"                                // [RFC5109][RFC3009], ?!
	TypeApplicationPassport                           = "application/passport"                                 // [RFC8225]
	TypeApplicationPatchopserrorxml                   = "application/patch-ops-error+xml"                      // [RFC5261]
	TypeApplicationPdf                                = "application/pdf"                                      // [RFC8118]
	TypeApplicationPemcertificatechain                = "application/pem-certificate-chain"                    // [RFC8555]
	TypeApplicationPgpencrypted                       = "application/pgp-encrypted"                            // [RFC3156]
	TypeApplicationPgpkeys                            = "application/pgp-keys"                                 // [RFC3156], ?!
	TypeApplicationPgpsignature                       = "application/pgp-signature"                            // [RFC3156]
	TypeApplicationPidfdiffxml                        = "application/pidf-diff+xml"                            // [RFC5262]
	TypeApplicationPidfxml                            = "application/pidf+xml"                                 // [RFC3863]
	TypeApplicationPkcs10                             = "application/pkcs10"                                   // [RFC5967]
	TypeApplicationPkcs7Mime                          = "application/pkcs7-mime"                               // [RFC8551][RFC7114]
	TypeApplicationPkcs7Signature                     = "application/pkcs7-signature"                          // [RFC8551]
	TypeApplicationPkcs8                              = "application/pkcs8"                                    // [RFC5958]
	TypeApplicationPkcs8Encrypted                     = "application/pkcs8-encrypted"                          // [RFC8351]
	TypeApplicationPkixattrcert                       = "application/pkix-attr-cert"                           // [RFC5877]
	TypeApplicationPkixcert                           = "application/pkix-cert"                                // [RFC2585]
	TypeApplicationPkixcrl                            = "application/pkix-crl"                                 // [RFC2585]
	TypeApplicationPkixpkipath                        = "application/pkix-pkipath"                             // [RFC6066]
	TypeApplicationPkixcmp                            = "application/pkixcmp"                                  // [RFC2510]
	TypeApplicationPlsxml                             = "application/pls+xml"                                  // [RFC4267]
	TypeApplicationPocsettingsxml                     = "application/poc-settings+xml"                         // [RFC4354]
	TypeApplicationPostscript                         = "application/postscript"                               // [RFC2045][RFC2046]
	TypeApplicationPpsptrackerjson                    = "application/ppsp-tracker+json"                        // [RFC7846]
	TypeApplicationProblemjson                        = "application/problem+json"                             // [RFC7807]
	TypeApplicationProblemxml                         = "application/problem+xml"                              // [RFC7807]
	TypeApplicationPskcxml                            = "application/pskc+xml"                                 // [RFC6030]
	TypeApplicationRdfxml                             = "application/rdf+xml"                                  // [RFC3870]
	TypeApplicationQsig                               = "application/QSIG"                                     // [RFC3204]
	TypeApplicationRaptorfec                          = "application/raptorfec"                                // [RFC6682]
	TypeApplicationRdapjson                           = "application/rdap+json"                                // [RFC7483]
	TypeApplicationReginfoxml                         = "application/reginfo+xml"                              // [RFC3680]
	TypeApplicationRemoteprinting                     = "application/remote-printing"                          // [RFC1486]
	TypeApplicationReputonjson                        = "application/reputon+json"                             // [RFC7071]
	TypeApplicationResourcelistsdiffxml               = "application/resource-lists-diff+xml"                  // [RFC5362]
	TypeApplicationResourcelistsxml                   = "application/resource-lists+xml"                       // [RFC4826]
	TypeApplicationRfcxml                             = "application/rfc+xml"                                  // [RFC7991]
	TypeApplicationRlmixml                            = "application/rlmi+xml"                                 // [RFC4662]
	TypeApplicationRlsservicesxml                     = "application/rls-services+xml"                         // [RFC4826]
	TypeApplicationRpkighostbusters                   = "application/rpki-ghostbusters"                        // [RFC6493]
	TypeApplicationRpkimanifest                       = "application/rpki-manifest"                            // [RFC6481]
	TypeApplicationRpkipublication                    = "application/rpki-publication"                         // [RFC8181]
	TypeApplicationRpkiroa                            = "application/rpki-roa"                                 // [RFC6481]
	TypeApplicationRpkiupdown                         = "application/rpki-updown"                              // [RFC6492]
	TypeApplicationRtploopback                        = "application/rtploopback"                              // [RFC6849]
	TypeApplicationRtx                                = "application/rtx"                                      // [RFC4588]
	TypeApplicationSbmlxml                            = "application/sbml+xml"                                 // [RFC3823]
	TypeApplicationScimjson                           = "application/scim+json"                                // [RFC7644]
	TypeApplicationScvpcvrequest                      = "application/scvp-cv-request"                          // [RFC5055]
	TypeApplicationScvpcvresponse                     = "application/scvp-cv-response"                         // [RFC5055]
	TypeApplicationScvpvprequest                      = "application/scvp-vp-request"                          // [RFC5055]
	TypeApplicationScvpvpresponse                     = "application/scvp-vp-response"                         // [RFC5055]
	TypeApplicationSdp                                = "application/sdp"                                      // [RFC4566]
	TypeApplicationSeceventjwt                        = "application/secevent+jwt"                             // [RFC8417]
	TypeApplicationSenmlexi                           = "application/senml-exi"                                // [RFC8428]
	TypeApplicationSenmlcbor                          = "application/senml+cbor"                               // [RFC8428]
	TypeApplicationSenmljson                          = "application/senml+json"                               // [RFC8428]
	TypeApplicationSenmlxml                           = "application/senml+xml"                                // [RFC8428]
	TypeApplicationSensmlexi                          = "application/sensml-exi"                               // [RFC8428]
	TypeApplicationSensmlcbor                         = "application/sensml+cbor"                              // [RFC8428]
	TypeApplicationSensmljson                         = "application/sensml+json"                              // [RFC8428]
	TypeApplicationSensmlxml                          = "application/sensml+xml"                               // [RFC8428]
	TypeApplicationSgml                               = "application/sgml"                                     // [RFC1874]
	TypeApplicationShfxml                             = "application/shf+xml"                                  // [RFC4194]
	TypeApplicationSieve                              = "application/sieve"                                    // [RFC5228]
	TypeApplicationSimplefilterxml                    = "application/simple-filter+xml"                        // [RFC4661]
	TypeApplicationSimplemessagesummary               = "application/simple-message-summary"                   // [RFC3842]
	TypeApplicationSmil                               = "application/smil"                                     // [RFC4536], Obsolete.
	TypeApplicationSmilxml                            = "application/smil+xml"                                 // [RFC4536]
	TypeApplicationSmpte336M                          = "application/smpte336m"                                // [RFC6597]
	TypeApplicationSoapxml                            = "application/soap+xml"                                 // [RFC3902]
	TypeApplicationSpiritseventxml                    = "application/spirits-event+xml"                        // [RFC3910]
	TypeApplicationSql                                = "application/sql"                                      // [RFC6922]
	TypeApplicationSrgs                               = "application/srgs"                                     // [RFC4267]
	TypeApplicationSrgsxml                            = "application/srgs+xml"                                 // [RFC4267]
	TypeApplicationSruxml                             = "application/sru+xml"                                  // [RFC6207]
	TypeApplicationSsmlxml                            = "application/ssml+xml"                                 // [RFC4267]
	TypeApplicationTampapexupdate                     = "application/tamp-apex-update"                         // [RFC5934]
	TypeApplicationTampapexupdateconfirm              = "application/tamp-apex-update-confirm"                 // [RFC5934]
	TypeApplicationTampcommunityupdate                = "application/tamp-community-update"                    // [RFC5934]
	TypeApplicationTampcommunityupdateconfirm         = "application/tamp-community-update-confirm"            // [RFC5934]
	TypeApplicationTamperror                          = "application/tamp-error"                               // [RFC5934]
	TypeApplicationTampsequenceadjust                 = "application/tamp-sequence-adjust"                     // [RFC5934]
	TypeApplicationTampsequenceadjustconfirm          = "application/tamp-sequence-adjust-confirm"             // [RFC5934]
	TypeApplicationTampstatusquery                    = "application/tamp-status-query"                        // [RFC5934]
	TypeApplicationTampstatusresponse                 = "application/tamp-status-response"                     // [RFC5934]
	TypeApplicationTampupdate                         = "application/tamp-update"                              // [RFC5934]
	TypeApplicationTampupdateconfirm                  = "application/tamp-update-confirm"                      // [RFC5934]
	TypeApplicationTeixml                             = "application/tei+xml"                                  // [RFC6129]
	TypeApplicationThraudxml                          = "application/thraud+xml"                               // [RFC5941]
	TypeApplicationTimestampquery                     = "application/timestamp-query"                          // [RFC3161]
	TypeApplicationTimestampreply                     = "application/timestamp-reply"                          // [RFC3161]
	TypeApplicationTimestampeddata                    = "application/timestamped-data"                         // [RFC5955]
	TypeApplicationTlsrptgzip                         = "application/tlsrpt+gzip"                              // [RFC8460]
	TypeApplicationTlsrptjson                         = "application/tlsrpt+json"                              // [RFC8460]
	TypeApplicationTnauthlist                         = "application/tnauthlist"                               // [RFC8226]
	TypeApplicationTrickleicesdpfrag                  = "application/trickle-ice-sdpfrag"                      // [RFC-ietf-mmusic-trickle-ice-sip-18]
	TypeApplicationTzif                               = "application/tzif"                                     // [RFC8536]
	TypeApplicationTzifleap                           = "application/tzif-leap"                                // [RFC8536]
	TypeApplicationUlpfec                             = "application/ulpfec"                                   // [RFC5109]
	TypeApplicationVcardjson                          = "application/vcard+json"                               // [RFC7095]
	TypeApplicationVcardxml                           = "application/vcard+xml"                                // [RFC6351]
	TypeApplicationVemmi                              = "application/vemmi"                                    // [RFC2122]
	TypeApplicationVndapplempegurl                    = "application/vnd.apple.mpegurl"                        // [RFC8216]
	TypeApplicationVndpwgmultiplexed                  = "application/vnd.pwg-multiplexed"                      // [RFC3391]
	TypeApplicationVndradisysmomlxml                  = "application/vnd.radisys.moml+xml"                     // [RFC5707]
	TypeApplicationVndradisysmsmlauditconfxml         = "application/vnd.radisys.msml-audit-conf+xml"          // [RFC5707]
	TypeApplicationVndradisysmsmlauditconnxml         = "application/vnd.radisys.msml-audit-conn+xml"          // [RFC5707]
	TypeApplicationVndradisysmsmlauditdialogxml       = "application/vnd.radisys.msml-audit-dialog+xml"        // [RFC5707]
	TypeApplicationVndradisysmsmlauditstreamxml       = "application/vnd.radisys.msml-audit-stream+xml"        // [RFC5707]
	TypeApplicationVndradisysmsmlauditxml             = "application/vnd.radisys.msml-audit+xml"               // [RFC5707]
	TypeApplicationVndradisysmsmlconfxml              = "application/vnd.radisys.msml-conf+xml"                // [RFC5707]
	TypeApplicationVndradisysmsmldialogbasexml        = "application/vnd.radisys.msml-dialog-base+xml"         // [RFC5707]
	TypeApplicationVndradisysmsmldialogfaxdetectxml   = "application/vnd.radisys.msml-dialog-fax-detect+xml"   // [RFC5707]
	TypeApplicationVndradisysmsmldialogfaxsendrecvxml = "application/vnd.radisys.msml-dialog-fax-sendrecv+xml" // [RFC5707]
	TypeApplicationVndradisysmsmldialoggroupxml       = "application/vnd.radisys.msml-dialog-group+xml"        // [RFC5707]
	TypeApplicationVndradisysmsmldialogspeechxml      = "application/vnd.radisys.msml-dialog-speech+xml"       // [RFC5707]
	TypeApplicationVndradisysmsmldialogtransformxml   = "application/vnd.radisys.msml-dialog-transform+xml"    // [RFC5707]
	TypeApplicationVndradisysmsmldialogxml            = "application/vnd.radisys.msml-dialog+xml"              // [RFC5707]
	TypeApplicationVndradisysmsmlxml                  = "application/vnd.radisys.msml+xml"                     // [RFC5707]
	TypeApplicationVoicexmlxml                        = "application/voicexml+xml"                             // [RFC4267]
	TypeApplicationVouchercmsjson                     = "application/voucher-cms+json"                         // [RFC8366]
	TypeApplicationVqrtcpxr                           = "application/vq-rtcpxr"                                // [RFC6035]
	TypeApplicationWatcherinfoxml                     = "application/watcherinfo+xml"                          // [RFC3858]
	TypeApplicationWebpushoptionsjson                 = "application/webpush-options+json"                     // [RFC8292]
	TypeApplicationWhoisppquery                       = "application/whoispp-query"                            // [RFC2957]
	TypeApplicationWhoisppresponse                    = "application/whoispp-response"                         // [RFC2958]
	TypeApplicationX400Bp                             = "application/x400-bp"                                  // [RFC1494]
	TypeApplicationXacmlxml                           = "application/xacml+xml"                                // [RFC7061]
	TypeApplicationXcapattxml                         = "application/xcap-att+xml"                             // [RFC4825]
	TypeApplicationXcapcapsxml                        = "application/xcap-caps+xml"                            // [RFC4825]
	TypeApplicationXcapdiffxml                        = "application/xcap-diff+xml"                            // [RFC5874]
	TypeApplicationXcapelxml                          = "application/xcap-el+xml"                              // [RFC4825]
	TypeApplicationXcaperrorxml                       = "application/xcap-error+xml"                           // [RFC4825]
	TypeApplicationXcapnsxml                          = "application/xcap-ns+xml"                              // [RFC4825]
	TypeApplicationXconconferenceinfodiffxml          = "application/xcon-conference-info-diff+xml"            // [RFC6502]
	TypeApplicationXconconferenceinfoxml              = "application/xcon-conference-info+xml"                 // [RFC6502]
	TypeApplicationXml                                = "application/xml"                                      // [RFC7303]
	TypeApplicationXmldtd                             = "application/xml-dtd"                                  // [RFC7303]
	TypeApplicationXmlexternalparsedentity            = "application/xml-external-parsed-entity"               // [RFC7303]
	TypeApplicationXmlpatchxml                        = "application/xml-patch+xml"                            // [RFC7351]
	TypeApplicationXmppxml                            = "application/xmpp+xml"                                 // [RFC3923]
	TypeApplicationXvxml                              = "application/xv+xml"                                   // [RFC4374]
	TypeApplicationYang                               = "application/yang"                                     // [RFC6020]
	TypeApplicationYangdatajson                       = "application/yang-data+json"                           // [RFC8040]
	TypeApplicationYangdataxml                        = "application/yang-data+xml"                            // [RFC8040]
	TypeApplicationYangpatchjson                      = "application/yang-patch+json"                          // [RFC8072]
	TypeApplicationYangpatchxml                       = "application/yang-patch+xml"                           // [RFC8072]
	TypeApplicationYinxml                             = "application/yin+xml"                                  // [RFC6020]
	TypeApplicationZlib                               = "application/zlib"                                     // [RFC6713]
	TypeApplicationZstd                               = "application/zstd"                                     // [RFC8478]
)

// Audio
const (
	TypeAudio1Dinterleavedparityfec = "audio/1d-interleaved-parityfec" // [RFC6015]
	TypeAudio32Kadpcm               = "audio/32kadpcm"                 // [RFC3802][RFC2421]
	TypeAudio3Gpp                   = "audio/3gpp"                     // [RFC3839][RFC6381]
	TypeAudio3Gpp2                  = "audio/3gpp2"                    // [RFC4393][RFC6381]
	TypeAudioAc3                    = "audio/ac3"                      // [RFC4184]
	TypeAudioAmr                    = "audio/AMR"                      // [RFC4867]
	TypeAudioAmrwb                  = "audio/AMR-WB"                   // [RFC4867]
	TypeAudioAmrwbplus              = "audio/amr-wb+"                  // [RFC4352]
	TypeAudioAptx                   = "audio/aptx"                     // [RFC7310]
	TypeAudioAsc                    = "audio/asc"                      // [RFC6295]
	TypeAudioAtracadvancedlossless  = "audio/ATRAC-ADVANCED-LOSSLESS"  // [RFC5584]
	TypeAudioAtracx                 = "audio/ATRAC-X"                  // [RFC5584]
	TypeAudioAtrac3                 = "audio/ATRAC3"                   // [RFC5584]
	TypeAudioBasic                  = "audio/basic"                    // [RFC2045][RFC2046]
	TypeAudioBv16                   = "audio/BV16"                     // [RFC4298]
	TypeAudioBv32                   = "audio/BV32"                     // [RFC4298]
	TypeAudioClearmode              = "audio/clearmode"                // [RFC4040]
	TypeAudioCn                     = "audio/CN"                       // [RFC3389]
	TypeAudioDat12                  = "audio/DAT12"                    // [RFC3190]
	TypeAudioDls                    = "audio/dls"                      // [RFC4613]
	TypeAudioDsres201108            = "audio/dsr-es201108"             // [RFC3557]
	TypeAudioDsres202050            = "audio/dsr-es202050"             // [RFC4060]
	TypeAudioDsres202211            = "audio/dsr-es202211"             // [RFC4060]
	TypeAudioDsres202212            = "audio/dsr-es202212"             // [RFC4060]
	TypeAudioDv                     = "audio/DV"                       // [RFC6469]
	TypeAudioDvi4                   = "audio/DVI4"                     // [RFC4856]
	TypeAudioEac3                   = "audio/eac3"                     // [RFC4598]
	TypeAudioEncaprtp               = "audio/encaprtp"                 // [RFC6849]
	TypeAudioEvrc                   = "audio/EVRC"                     // [RFC4788]
	TypeAudioEvrcqcp                = "audio/EVRC-QCP"                 // [RFC3625]
	TypeAudioEvrc0                  = "audio/EVRC0"                    // [RFC4788]
	TypeAudioEvrc1                  = "audio/EVRC1"                    // [RFC4788]
	TypeAudioEvrcb                  = "audio/EVRCB"                    // [RFC5188]
	TypeAudioEvrcb0                 = "audio/EVRCB0"                   // [RFC5188]
	TypeAudioEvrcb1                 = "audio/EVRCB1"                   // [RFC4788]
	TypeAudioEvrcnw                 = "audio/EVRCNW"                   // [RFC6884]
	TypeAudioEvrcnw0                = "audio/EVRCNW0"                  // [RFC6884]
	TypeAudioEvrcnw1                = "audio/EVRCNW1"                  // [RFC6884]
	TypeAudioEvrcwb                 = "audio/EVRCWB"                   // [RFC5188]
	TypeAudioEvrcwb0                = "audio/EVRCWB0"                  // [RFC5188]
	TypeAudioEvrcwb1                = "audio/EVRCWB1"                  // [RFC5188]
	TypeAudioExample                = "audio/example"                  // [RFC4735]
	TypeAudioFlexfec                = "audio/flexfec"                  // [RFC-ietf-payload-flexible-fec-scheme-20]
	TypeAudioFwdred                 = "audio/fwdred"                   // [RFC6354]
	TypeAudioG7110                  = "audio/G711-0"                   // [RFC7655]
	TypeAudioG719                   = "audio/G719"                     // [RFC5404][RFC Errata 3245]
	TypeAudioG7221                  = "audio/G7221"                    // [RFC5577]
	TypeAudioG722                   = "audio/G722"                     // [RFC4856]
	TypeAudioG723                   = "audio/G723"                     // [RFC4856]
	TypeAudioG72616                 = "audio/G726-16"                  // [RFC4856]
	TypeAudioG72624                 = "audio/G726-24"                  // [RFC4856]
	TypeAudioG72632                 = "audio/G726-32"                  // [RFC4856]
	TypeAudioG72640                 = "audio/G726-40"                  // [RFC4856]
	TypeAudioG728                   = "audio/G728"                     // [RFC4856]
	TypeAudioG729                   = "audio/G729"                     // [RFC4856]
	TypeAudioG7291                  = "audio/G7291"                    // [RFC4749][RFC5459], ?!
	TypeAudioG729D                  = "audio/G729D"                    // [RFC4856]
	TypeAudioG729E                  = "audio/G729E"                    // [RFC4856]
	TypeAudioGsm                    = "audio/GSM"                      // [RFC4856]
	TypeAudioGsmefr                 = "audio/GSM-EFR"                  // [RFC4856]
	TypeAudioGsmhr08                = "audio/GSM-HR-08"                // [RFC5993]
	TypeAudioIlbc                   = "audio/iLBC"                     // [RFC3952]
	TypeAudioIpmrv25                = "audio/ip-mr_v2.5"               // [RFC6262]
	TypeAudioL8                     = "audio/L8"                       // [RFC4856]
	TypeAudioL16                    = "audio/L16"                      // [RFC4856]
	TypeAudioL20                    = "audio/L20"                      // [RFC3190]
	TypeAudioL24                    = "audio/L24"                      // [RFC3190]
	TypeAudioLpc                    = "audio/LPC"                      // [RFC4856]
	TypeAudioMelp                   = "audio/MELP"                     // [RFC8130]
	TypeAudioMelp600                = "audio/MELP600"                  // [RFC8130]
	TypeAudioMelp1200               = "audio/MELP1200"                 // [RFC8130]
	TypeAudioMelp2400               = "audio/MELP2400"                 // [RFC8130]
	TypeAudioMobilexmf              = "audio/mobile-xmf"               // [RFC4723]
	TypeAudioMpa                    = "audio/MPA"                      // [RFC3555]
	TypeAudioMp4                    = "audio/mp4"                      // [RFC4337][RFC6381]
	TypeAudioMp4Alatm               = "audio/MP4A-LATM"                // [RFC6416]
	TypeAudioMparobust              = "audio/mpa-robust"               // [RFC5219]
	TypeAudioMpeg                   = "audio/mpeg"                     // [RFC3003]
	TypeAudioMpeg4Generic           = "audio/mpeg4-generic"            // [RFC3640][RFC5691][RFC6295]
	TypeAudioOgg                    = "audio/ogg"                      // [RFC5334][RFC7845]
	TypeAudioOpus                   = "audio/opus"                     // [RFC7587]
	TypeAudioParityfec              = "audio/parityfec"                // [RFC5109][RFC3009], ?!
	TypeAudioPcma                   = "audio/PCMA"                     // [RFC4856]
	TypeAudioPcmawb                 = "audio/PCMA-WB"                  // [RFC5391]
	TypeAudioPcmu                   = "audio/PCMU"                     // [RFC4856]
	TypeAudioPcmuwb                 = "audio/PCMU-WB"                  // [RFC5391]
	TypeAudioQcelp                  = "audio/QCELP"                    // [RFC3555][RFC3625], ?!
	TypeAudioRaptorfec              = "audio/raptorfec"                // [RFC6682]
	TypeAudioRed                    = "audio/RED"                      // [RFC3555]
	TypeAudioRtploopback            = "audio/rtploopback"              // [RFC6849]
	TypeAudioRtpmidi                = "audio/rtp-midi"                 // [RFC6295]
	TypeAudioRtx                    = "audio/rtx"                      // [RFC4588]
	TypeAudioSmv                    = "audio/SMV"                      // [RFC3558]
	TypeAudioSmv0                   = "audio/SMV0"                     // [RFC3558]
	TypeAudioSmvqcp                 = "audio/SMV-QCP"                  // [RFC3625]
	TypeAudioSpeex                  = "audio/speex"                    // [RFC5574]
	TypeAudioT140C                  = "audio/t140c"                    // [RFC4351]
	TypeAudioT38                    = "audio/t38"                      // [RFC4612]
	TypeAudioTelephoneevent         = "audio/telephone-event"          // [RFC4733]
	TypeAudioTone                   = "audio/tone"                     // [RFC4733]
	TypeAudioUemclip                = "audio/UEMCLIP"                  // [RFC5686]
	TypeAudioUlpfec                 = "audio/ulpfec"                   // [RFC5109]
	TypeAudioVdvi                   = "audio/VDVI"                     // [RFC4856]
	TypeAudioVmrwb                  = "audio/VMR-WB"                   // [RFC4348][RFC4424]
	TypeAudioVndqcelp               = "audio/vnd.qcelp"                // [RFC3625], Deprecated.
	TypeAudioVorbis                 = "audio/vorbis"                   // [RFC5215]
	TypeAudioVorbisconfig           = "audio/vorbis-config"            // [RFC5215]
)

// Font.
const (
	TypeFontCollection = "font/collection" // [RFC8081]
	TypeFontOtf        = "font/otf"        // [RFC8081]
	TypeFontSfnt       = "font/sfnt"       // [RFC8081]
	TypeFontTtf        = "font/ttf"        // [RFC8081]
	TypeFontWoff       = "font/woff"       // [RFC8081]
	TypeFontWoff2      = "font/woff2"      // [RFC8081]
)

// Image.
const (
	TypeImageBmp     = "image/bmp"     // [RFC7903]
	TypeImageEmf     = "image/emf"     // [RFC7903]
	TypeImageExample = "image/example" // [RFC4735]
	TypeImageFits    = "image/fits"    // [RFC4047]
	TypeImageG3Fax   = "image/g3fax"   // [RFC1494]
	TypeImageGif     = "image/gif"     // [RFC2045][RFC2046], ?!
	TypeImageIef     = "image/ief"     // [RFC1314], ?!
	TypeImageJp2     = "image/jp2"     // [RFC3745]
	TypeImageJpeg    = "image/jpeg"    // [RFC2045][RFC2046], ?!
	TypeImageJpm     = "image/jpm"     // [RFC3745]
	TypeImageJpx     = "image/jpx"     // [RFC3745]
	TypeImageT38     = "image/t38"     // [RFC3362]
	TypeImageTiff    = "image/tiff"    // [RFC3302]
	TypeImageTifffx  = "image/tiff-fx" // [RFC3950]
	TypeImageWmf     = "image/wmf"     // [RFC7903]
	TypeImageXemf    = "image/emf"     // [RFC7903], Deprecated.
	TypeImageXwmf    = "image/wmf"     // [RFC7903], Deprecated.
)

// Message.
const (
	TypeMessageCpim                          = "message/CPIM"                            // [RFC3862]
	TypeMessageDeliverystatus                = "message/delivery-status"                 // [RFC1894]
	TypeMessageDispositionnotification       = "message/disposition-notification"        // [RFC8098]
	TypeMessageExample                       = "message/example"                         // [RFC4735]
	TypeMessageExternalbody                  = "message/external-body"                   // [RFC2045][RFC2046], ?!
	TypeMessageFeedbackreport                = "message/feedback-report"                 // [RFC5965]
	TypeMessageGlobal                        = "message/global"                          // [RFC6532]
	TypeMessageGlobaldeliverystatus          = "message/global-delivery-status"          // [RFC6533]
	TypeMessageGlobaldispositionnotification = "message/global-disposition-notification" // [RFC6533]
	TypeMessageGlobalheaders                 = "message/global-headers"                  // [RFC6533]
	TypeMessageHttp                          = "message/http"                            // [RFC7230]
	TypeMessageImdnxml                       = "message/imdn+xml"                        // [RFC5438]
	TypeMessageNews                          = "message/news"                            // [RFC5537], Obsolete.
	TypeMessagePartial                       = "message/partial"                         // [RFC2045][RFC2046], ?!
	TypeMessageRfc822                        = "message/rfc822"                          // [RFC2045][RFC2046], ?!
	TypeMessageShttp                         = "message/s-http"                          // [RFC2660]
	TypeMessageSip                           = "message/sip"                             // [RFC3261]
	TypeMessageSipfrag                       = "message/sipfrag"                         // [RFC3420]
	TypeMessageTrackingstatus                = "message/tracking-status"                 // [RFC3886]
)

// Model.
const (
	TypeModelExample = "model/example" // [RFC4735]
	TypeModelMesh    = "model/mesh"    // [RFC2077], ?!
	TypeModelVrml    = "model/vrml"    // [RFC2077], ?!
)

// Multipart.
const (
	TypeMultipartAlternative  = "multipart/alternative"   // [RFC2046][RFC2045], ?!
	TypeMultipartByteranges   = "multipart/byteranges"    // [RFC7233]
	TypeMultipartDigest       = "multipart/digest"        // [RFC2046][RFC2045], ?!
	TypeMultipartEncrypted    = "multipart/encrypted"     // [RFC1847]
	TypeMultipartExample      = "multipart/example"       // [RFC4735]
	TypeMultipartFormdata     = "multipart/form-data"     // [RFC7578]
	TypeMultipartMixed        = "multipart/mixed"         // [RFC2046][RFC2045]
	TypeMultipartMultilingual = "multipart/multilingual"  // [RFC8255]
	TypeMultipartParallel     = "multipart/parallel"      // [RFC2046][RFC2045], ?!
	TypeMultipartRelated      = "multipart/related"       // [RFC2387]
	TypeMultipartReport       = "multipart/report"        // [RFC6522]
	TypeMultipartSigned       = "multipart/signed"        // [RFC1847]
	TypeMultipartVoicemessage = "multipart/voice-message" // [RFC3801]
)

// Text.
const (
	TypeText1Dinterleavedparityfec    = "text/1d-interleaved-parityfec"      // [RFC6015]
	TypeTextCalendar                  = "text/calendar"                      // [RFC5545]
	TypeTextCss                       = "text/css"                           // [RFC2318]
	TypeTextCsv                       = "text/csv"                           // [RFC4180][RFC7111]
	TypeTextDirectory                 = "text/directory"                     // [RFC2425][RFC6350], Deprecated by [RFC6350].
	TypeTextDns                       = "text/dns"                           // [RFC4027]
	TypeTextEcmascript                = "text/ecmascript"                    // [RFC4329], Obsolete.
	TypeTextEncaprtp                  = "text/encaprtp"                      // [RFC6849]
	TypeTextEnriched                  = "text/enriched"                      // [RFC1896]
	TypeTextExample                   = "text/example"                       // [RFC4735]
	TypeTextFlexfec                   = "text/flexfec"                       // [RFC-ietf-payload-flexible-fec-scheme-20]
	TypeTextFwdred                    = "text/fwdred"                        // [RFC6354]
	TypeTextGrammarreflist            = "text/grammar-ref-list"              // [RFC6787]
	TypeTextHtml                      = "text/html"                          // [RFC2854] ?!
	TypeTextJavascript                = "text/javascript"                    // [RFC4329], Obsolete.
	TypeTextJscript                   = "text/jscript"                       // ?! Obsolete.
	TypeTextMarkdown                  = "text/markdown"                      // [RFC7763]
	TypeTextParameters                = "text/parameters"                    // [RFC7826]
	TypeTextParityfec                 = "text/parityfec"                     // [RFC5109][RFC3009], ?!
	TypeTextPlain                     = "text/plain"                         // [RFC2046][RFC3676][RFC5147]
	TypeTextRaptorfec                 = "text/raptorfec"                     // [RFC6682]
	TypeTextRed                       = "text/RED"                           // [RFC4102]
	TypeTextRfc822Headers             = "text/rfc822-headers"                // [RFC6522]
	TypeTextRichtext                  = "text/richtext"                      // [RFC2045][RFC2046] // ?!
	TypeTextRtploopback               = "text/rtploopback"                   // [RFC6849]
	TypeTextRtx                       = "text/rtx"                           // [RFC4588]
	TypeTextSgml                      = "text/sgml"                          // [RFC1874]
	TypeTextT140                      = "text/t140"                          // [RFC4103]
	TypeTextTroff                     = "text/troff"                         // [RFC4263]
	TypeTextUlpfec                    = "text/ulpfec"                        // [RFC5109]
	TypeTextUrilist                   = "text/uri-list"                      // [RFC2483]
	TypeTextVcard                     = "text/vcard"                         // [RFC6350]
	TypeTextVndradisysmsmlbasiclayout = "text/vnd.radisys.msml-basic-layout" // [RFC5707]
	TypeTextXml                       = "text/xml"                           // [RFC7303]
	TypeTextXmlexternalparsedentity   = "text/xml-external-parsed-entity"    // [RFC7303]
)

// Video.
const (
	TypeVideo1Dinterleavedparityfec = "video/1d-interleaved-parityfec" // [RFC6015]
	TypeVideo3Gpp                   = "video/3gpp"                     // [RFC3839][RFC6381]
	TypeVideo3Gpp2                  = "video/3gpp2"                    // [RFC4393][RFC6381]
	TypeVideo3Gpptt                 = "video/3gpp-tt"                  // [RFC4396]
	TypeVideoBmpeg                  = "video/BMPEG"                    // [RFC3555]
	TypeVideoBt656                  = "video/BT656"                    // [RFC3555]
	TypeVideoCelb                   = "video/CelB"                     // [RFC3555]
	TypeVideoDv                     = "video/DV"                       // [RFC6469]
	TypeVideoEncaprtp               = "video/encaprtp"                 // [RFC6849]
	TypeVideoExample                = "video/example"                  // [RFC4735]
	TypeVideoFlexfec                = "video/flexfec"                  // [RFC-ietf-payload-flexible-fec-scheme-20]
	TypeVideoH261                   = "video/H261"                     // [RFC4587]
	TypeVideoH263                   = "video/H263"                     // [RFC3555]
	TypeVideoH2631998               = "video/H263-1998"                // [RFC4629]
	TypeVideoH2632000               = "video/H263-2000"                // [RFC4629]
	TypeVideoH264                   = "video/H264"                     // [RFC6184]
	TypeVideoH264Rcdo               = "video/H264-RCDO"                // [RFC6185]
	TypeVideoH264Svc                = "video/H264-SVC"                 // [RFC6190]
	TypeVideoH265                   = "video/H265"                     // [RFC7798]
	TypeVideoJpeg                   = "video/JPEG"                     // [RFC3555]
	TypeVideoJpeg2000               = "video/jpeg2000"                 // [RFC5371][RFC5372]
	TypeVideoMj2                    = "video/mj2"                      // [RFC3745]
	TypeVideoMp1S                   = "video/MP1S"                     // [RFC3555]
	TypeVideoMp2P                   = "video/MP2P"                     // [RFC3555]
	TypeVideoMp2T                   = "video/MP2T"                     // [RFC3555]
	TypeVideoMp4                    = "video/mp4"                      // [RFC4337][RFC6381]
	TypeVideoMp4Ves                 = "video/MP4V-ES"                  // [RFC6416]
	TypeVideoMpv                    = "video/MPV"                      // [RFC3555]
	TypeVideoMpeg                   = "video/mpeg"                     // [RFC2045][RFC2046], ?!
	TypeVideoMpeg4Generic           = "video/mpeg4-generic"            // [RFC3640]
	TypeVideoNv                     = "video/nv"                       // [RFC4856]
	TypeVideoOgg                    = "video/ogg"                      // [RFC5334][RFC7845]
	TypeVideoParityfec              = "video/parityfec"                // [RFC5109][RFC3009], ?!
	TypeVideoPointer                = "video/pointer"                  // [RFC2862]
	TypeVideoQuicktime              = "video/quicktime"                // [RFC6381][Paul_Lindner]
	TypeVideoRaptorfec              = "video/raptorfec"                // [RFC6682]
	TypeVideoRaw                    = "video/raw"                      // [RFC4175], ?!
	TypeVideoRtploopback            = "video/rtploopback"              // [RFC6849]
	TypeVideoRtx                    = "video/rtx"                      // [RFC4588]
	TypeVideoSmpte291               = "video/smpte291"                 // [RFC8331]
	TypeVideoSmpte292M              = "video/SMPTE292M"                // [RFC3497]
	TypeVideoUlpfec                 = "video/ulpfec"                   // [RFC5109]
	TypeVideoVc1                    = "video/vc1"                      // [RFC4425]
	TypeVideoVc2                    = "video/vc2"                      // [RFC8450]
	TypeVideoVp8                    = "video/VP8"                      // [RFC7741]
)

// Vendor.
const (
	TypeVendorKml              = "application/vnd.google-earth.kml+xml"
	TypeVendorMozillaXul       = "application/vnd.mozilla.xul+xml"
	TypeVendorMsExcel          = "application/vnd.ms-excel"
	TypeVendorMsExcel2007      = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	TypeVendorMsPowerpoint     = "application/vnd.ms-powerpoint"
	TypeVendorMsPowerpoint2007 = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	TypeVendorMsWord           = "application/msword"
	TypeVendorMsWord2007       = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	TypeVendorOdg              = "application/vnd.oasis.opendocument.graphics"
	TypeVendorOdp              = "application/vnd.oasis.opendocument.presentation"
	TypeVendorOds              = "application/vnd.oasis.opendocument.spreadsheet"
	TypeVendorOdt              = "application/vnd.oasis.opendocument.text"
)

// Non-standard.
const (
	TypeXDvi            = "application/x-dvi"
	TypeXEcmascript     = "application/x-ecmascript"
	TypeXFormUrlencoded = "application/x-www-form-urlencoded"
	TypeXJavascript     = "application/x-javascript"
	TypeXJquery         = "text/x-jquery-tmpl"
	TypeXLatex          = "application/x-latex"
	TypeXRar            = "application/x-rar-compressed"
	TypeXShockwaveflash = "application/x-shockwave-flash"
	TypeXStuffit        = "application/x-stuffit"
	TypeXTarball        = "application/x-tar"
	TypeXTtf            = "application/x-font-ttf"
)

// Public Key Cryptography Standards.
const (
	TypePkcsP12 = "application/x-pkcs12"
	TypePkcsPfx = "application/x-pkcs12"
	TypePkcsP7b = "application/x-pkcs7-certificates"
	TypePkcsSpc = "application/x-pkcs7-certificates"
	TypePkcsP7r = "application/x-pkcs7-certreqresp"
	TypePkcsP7c = "application/x-pkcs7-mime"
	TypePkcsP7m = "application/x-pkcs7-mime"
	TypePkcsP7s = "application/x-pkcs7-signature"
)
