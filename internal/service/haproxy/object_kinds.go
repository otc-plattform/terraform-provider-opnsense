package haproxy

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	ophaproxy "github.com/browningluke/opnsense-go/pkg/haproxy"
)

var haproxyBackendKind = haproxyObjectKind{
	typeSuffix:     "haproxy_backend",
	summaryName:    "HAProxy backend",
	resourceDesc:   "Manage an OPNsense HAProxy backend pool using the raw HAProxy API payload.",
	dataSourceDesc: "Read an OPNsense HAProxy backend pool using the raw HAProxy API payload.",
	search: func(ctx context.Context, c *ophaproxy.Controller) (*ophaproxy.HAProxyObjectSearchResult, error) {
		return c.HAProxySearchBackends(ctx)
	},
	get: func(ctx context.Context, c *ophaproxy.Controller, id string) (ophaproxy.HAProxyObject, error) {
		return c.HAProxyGetBackend(ctx, id)
	},
	add: func(ctx context.Context, c *ophaproxy.Controller, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyAddBackend(ctx, object)
	},
	edit: func(ctx context.Context, c *ophaproxy.Controller, id string, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyEditBackend(ctx, id, object)
	},
	delete: func(ctx context.Context, c *ophaproxy.Controller, id string) (*api.ActionResult, error) {
		return c.HAProxyDeleteBackend(ctx, id)
	},
}

var haproxyFrontendKind = haproxyObjectKind{
	typeSuffix:     "haproxy_frontend",
	summaryName:    "HAProxy frontend",
	resourceDesc:   "Manage an OPNsense HAProxy public service using the raw HAProxy API payload.",
	dataSourceDesc: "Read an OPNsense HAProxy public service using the raw HAProxy API payload.",
	search: func(ctx context.Context, c *ophaproxy.Controller) (*ophaproxy.HAProxyObjectSearchResult, error) {
		return c.HAProxySearchFrontends(ctx)
	},
	get: func(ctx context.Context, c *ophaproxy.Controller, id string) (ophaproxy.HAProxyObject, error) {
		return c.HAProxyGetFrontend(ctx, id)
	},
	add: func(ctx context.Context, c *ophaproxy.Controller, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyAddFrontend(ctx, object)
	},
	edit: func(ctx context.Context, c *ophaproxy.Controller, id string, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyEditFrontend(ctx, id, object)
	},
	delete: func(ctx context.Context, c *ophaproxy.Controller, id string) (*api.ActionResult, error) {
		return c.HAProxyDeleteFrontend(ctx, id)
	},
}

var haproxyACLKind = haproxyObjectKind{
	typeSuffix:     "haproxy_acl",
	summaryName:    "HAProxy ACL",
	resourceDesc:   "Manage an OPNsense HAProxy condition/ACL using the raw HAProxy API payload.",
	dataSourceDesc: "Read an OPNsense HAProxy condition/ACL using the raw HAProxy API payload.",
	search: func(ctx context.Context, c *ophaproxy.Controller) (*ophaproxy.HAProxyObjectSearchResult, error) {
		return c.HAProxySearchACLs(ctx)
	},
	get: func(ctx context.Context, c *ophaproxy.Controller, id string) (ophaproxy.HAProxyObject, error) {
		return c.HAProxyGetACL(ctx, id)
	},
	add: func(ctx context.Context, c *ophaproxy.Controller, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyAddACL(ctx, object)
	},
	edit: func(ctx context.Context, c *ophaproxy.Controller, id string, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyEditACL(ctx, id, object)
	},
	delete: func(ctx context.Context, c *ophaproxy.Controller, id string) (*api.ActionResult, error) {
		return c.HAProxyDeleteACL(ctx, id)
	},
}

var haproxyActionKind = haproxyObjectKind{
	typeSuffix:     "haproxy_action",
	summaryName:    "HAProxy action",
	resourceDesc:   "Manage an OPNsense HAProxy rule/action using the raw HAProxy API payload.",
	dataSourceDesc: "Read an OPNsense HAProxy rule/action using the raw HAProxy API payload.",
	search: func(ctx context.Context, c *ophaproxy.Controller) (*ophaproxy.HAProxyObjectSearchResult, error) {
		return c.HAProxySearchActions(ctx)
	},
	get: func(ctx context.Context, c *ophaproxy.Controller, id string) (ophaproxy.HAProxyObject, error) {
		return c.HAProxyGetAction(ctx, id)
	},
	add: func(ctx context.Context, c *ophaproxy.Controller, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyAddAction(ctx, object)
	},
	edit: func(ctx context.Context, c *ophaproxy.Controller, id string, object ophaproxy.HAProxyObject) (*api.ActionResult, error) {
		return c.HAProxyEditAction(ctx, id, object)
	},
	delete: func(ctx context.Context, c *ophaproxy.Controller, id string) (*api.ActionResult, error) {
		return c.HAProxyDeleteAction(ctx, id)
	},
}
