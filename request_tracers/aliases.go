// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package request_tracers

import (
	"github.com/khulnasoft/khulnasoft-go/internal/apierror"
	"github.com/khulnasoft/khulnasoft-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type ASN = shared.ASN

// This is an alias to an internal type.
type ASNParam = shared.ASNParam

// This is an alias to an internal type.
type AuditLog = shared.AuditLog

// This is an alias to an internal type.
type AuditLogAction = shared.AuditLogAction

// This is an alias to an internal type.
type AuditLogActor = shared.AuditLogActor

// The type of actor, whether a User, Khulnasoft Admin, or an Automated System.
//
// This is an alias to an internal type.
type AuditLogActorType = shared.AuditLogActorType

// This is an alias to an internal value.
const AuditLogActorTypeUser = shared.AuditLogActorTypeUser

// This is an alias to an internal value.
const AuditLogActorTypeAdmin = shared.AuditLogActorTypeAdmin

// This is an alias to an internal value.
const AuditLogActorTypeKhulnasoft = shared.AuditLogActorTypeKhulnasoft

// This is an alias to an internal type.
type AuditLogOwner = shared.AuditLogOwner

// This is an alias to an internal type.
type AuditLogResource = shared.AuditLogResource

// The Certificate Authority that will issue the certificate
//
// This is an alias to an internal type.
type CertificateCA = shared.CertificateCA

// This is an alias to an internal value.
const CertificateCADigicert = shared.CertificateCADigicert

// This is an alias to an internal value.
const CertificateCAGoogle = shared.CertificateCAGoogle

// This is an alias to an internal value.
const CertificateCALetsEncrypt = shared.CertificateCALetsEncrypt

// This is an alias to an internal value.
const CertificateCASSLCom = shared.CertificateCASSLCom

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
//
// This is an alias to an internal type.
type CertificateRequestType = shared.CertificateRequestType

// This is an alias to an internal value.
const CertificateRequestTypeOriginRSA = shared.CertificateRequestTypeOriginRSA

// This is an alias to an internal value.
const CertificateRequestTypeOriginECC = shared.CertificateRequestTypeOriginECC

// This is an alias to an internal value.
const CertificateRequestTypeKeylessCertificate = shared.CertificateRequestTypeKeylessCertificate

// A Khulnasoft Tunnel that connects your origin to Khulnasoft's edge.
//
// This is an alias to an internal type.
type KhulnasoftTunnel = shared.KhulnasoftTunnel

// This is an alias to an internal type.
type KhulnasoftTunnelConnection = shared.KhulnasoftTunnelConnection

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Khulnasoft Edge).
//
// This is an alias to an internal type.
type KhulnasoftTunnelStatus = shared.KhulnasoftTunnelStatus

// This is an alias to an internal value.
const KhulnasoftTunnelStatusInactive = shared.KhulnasoftTunnelStatusInactive

// This is an alias to an internal value.
const KhulnasoftTunnelStatusDegraded = shared.KhulnasoftTunnelStatusDegraded

// This is an alias to an internal value.
const KhulnasoftTunnelStatusHealthy = shared.KhulnasoftTunnelStatusHealthy

// This is an alias to an internal value.
const KhulnasoftTunnelStatusDown = shared.KhulnasoftTunnelStatusDown

// The type of tunnel.
//
// This is an alias to an internal type.
type KhulnasoftTunnelTunType = shared.KhulnasoftTunnelTunType

// This is an alias to an internal value.
const KhulnasoftTunnelTunTypeCfdTunnel = shared.KhulnasoftTunnelTunTypeCfdTunnel

// This is an alias to an internal value.
const KhulnasoftTunnelTunTypeWARPConnector = shared.KhulnasoftTunnelTunTypeWARPConnector

// This is an alias to an internal value.
const KhulnasoftTunnelTunTypeIPSec = shared.KhulnasoftTunnelTunTypeIPSec

// This is an alias to an internal value.
const KhulnasoftTunnelTunTypeGRE = shared.KhulnasoftTunnelTunTypeGRE

// This is an alias to an internal value.
const KhulnasoftTunnelTunTypeCNI = shared.KhulnasoftTunnelTunTypeCNI

// This is an alias to an internal type.
type ErrorData = shared.ErrorData

// This is an alias to an internal type.
type MemberParam = shared.MemberParam

// This is an alias to an internal type.
type MemberRoleParam = shared.MemberRoleParam

// This is an alias to an internal type.
type MemberRolesPermissionsParam = shared.MemberRolesPermissionsParam

// A member's status in the account.
//
// This is an alias to an internal type.
type MemberStatus = shared.MemberStatus

// This is an alias to an internal value.
const MemberStatusAccepted = shared.MemberStatusAccepted

// This is an alias to an internal value.
const MemberStatusPending = shared.MemberStatusPending

// Details of the user associated to the membership.
//
// This is an alias to an internal type.
type MemberUserParam = shared.MemberUserParam

// This is an alias to an internal type.
type Permission = shared.Permission

// This is an alias to an internal type.
type PermissionGrant = shared.PermissionGrant

// This is an alias to an internal type.
type PermissionGrantParam = shared.PermissionGrantParam

// The rate plan applied to the subscription.
//
// This is an alias to an internal type.
type RatePlan = shared.RatePlan

// The rate plan applied to the subscription.
//
// This is an alias to an internal type.
type RatePlanParam = shared.RatePlanParam

// This is an alias to an internal type.
type ResponseInfo = shared.ResponseInfo

// This is an alias to an internal type.
type Role = shared.Role

// Direction to order DNS records in.
//
// This is an alias to an internal type.
type SortDirection = shared.SortDirection

// This is an alias to an internal value.
const SortDirectionAsc = shared.SortDirectionAsc

// This is an alias to an internal value.
const SortDirectionDesc = shared.SortDirectionDesc

// This is an alias to an internal type.
type Subscription = shared.Subscription

// How often the subscription is renewed automatically.
//
// This is an alias to an internal type.
type SubscriptionFrequency = shared.SubscriptionFrequency

// This is an alias to an internal value.
const SubscriptionFrequencyWeekly = shared.SubscriptionFrequencyWeekly

// This is an alias to an internal value.
const SubscriptionFrequencyMonthly = shared.SubscriptionFrequencyMonthly

// This is an alias to an internal value.
const SubscriptionFrequencyQuarterly = shared.SubscriptionFrequencyQuarterly

// This is an alias to an internal value.
const SubscriptionFrequencyYearly = shared.SubscriptionFrequencyYearly

// The state that the subscription is in.
//
// This is an alias to an internal type.
type SubscriptionState = shared.SubscriptionState

// This is an alias to an internal value.
const SubscriptionStateTrial = shared.SubscriptionStateTrial

// This is an alias to an internal value.
const SubscriptionStateProvisioned = shared.SubscriptionStateProvisioned

// This is an alias to an internal value.
const SubscriptionStatePaid = shared.SubscriptionStatePaid

// This is an alias to an internal value.
const SubscriptionStateAwaitingPayment = shared.SubscriptionStateAwaitingPayment

// This is an alias to an internal value.
const SubscriptionStateCancelled = shared.SubscriptionStateCancelled

// This is an alias to an internal value.
const SubscriptionStateFailed = shared.SubscriptionStateFailed

// This is an alias to an internal value.
const SubscriptionStateExpired = shared.SubscriptionStateExpired

// This is an alias to an internal type.
type SubscriptionParam = shared.SubscriptionParam
