package client

import (
	"context"
	"runtime/debug"
	"strings"
)

const (
	onaSDKModulePath = "github.com/gitpod-io/gitpod-sdk-go"

	userAgentLanguage = "go"
	userAgentLayerRaw = "raw"
	userAgentLayerSDK = "sdk"
)

// SDKVersion is the version encoded in SDK User-Agent values.
//
// Release builds may override this value. When left as dev, the client uses
// Go build information to find the module version when available.
var SDKVersion = "dev"

type userAgentLayerContextKey struct{}

func withSDKUserAgentLayer(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if layer, ok := ctx.Value(userAgentLayerContextKey{}).(string); ok && layer == userAgentLayerSDK {
		return ctx
	}
	return context.WithValue(ctx, userAgentLayerContextKey{}, userAgentLayerSDK)
}

// WithSDKUserAgentLayer marks requests made through the high-level SDK layer.
func WithSDKUserAgentLayer(ctx context.Context) context.Context {
	return withSDKUserAgentLayer(ctx)
}

// SDKUserAgent returns the Ona Go SDK User-Agent token.
func SDKUserAgent() string {
	return onaUserAgent("", userAgentLayerSDK)
}

func userAgentLayerFromContext(ctx context.Context) string {
	if ctx == nil {
		return userAgentLayerRaw
	}
	layer, _ := ctx.Value(userAgentLayerContextKey{}).(string)
	if layer == userAgentLayerSDK {
		return userAgentLayerSDK
	}
	return userAgentLayerRaw
}

func onaUserAgent(prefix string, layer string) string {
	version := sdkUserAgentVersion()
	product := "ona-go-client"
	if layer == userAgentLayerSDK {
		product = "ona-go-sdk"
	}

	ua := product + "/" + version + " (language=" + userAgentLanguage + "; layer=" + layer + ")"
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return ua
	}
	return prefix + " " + ua
}

func sdkUserAgentVersion() string {
	if SDKVersion != "" && SDKVersion != "dev" {
		return sanitizeUserAgentToken(SDKVersion)
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Path == onaSDKModulePath && isBuildInfoVersion(info.Main.Version) {
			return sanitizeUserAgentToken(info.Main.Version)
		}
		for _, dep := range info.Deps {
			if dep.Path == onaSDKModulePath && isBuildInfoVersion(dep.Version) {
				return sanitizeUserAgentToken(dep.Version)
			}
			if dep.Replace != nil && dep.Replace.Path == onaSDKModulePath && isBuildInfoVersion(dep.Replace.Version) {
				return sanitizeUserAgentToken(dep.Replace.Version)
			}
		}
	}

	if SDKVersion != "" {
		return sanitizeUserAgentToken(SDKVersion)
	}
	return "dev"
}

func isBuildInfoVersion(version string) bool {
	return version != "" && version != "(devel)"
}

func sanitizeUserAgentToken(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "dev"
	}

	var b strings.Builder
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '.', r == '-', r == '_', r == '+', r == '~':
			b.WriteRune(r)
		default:
			b.WriteByte('-')
		}
	}
	if b.Len() == 0 {
		return "dev"
	}
	return b.String()
}
