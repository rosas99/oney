// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

// Package app does all of the work necessary to create a OneX
// APIServer by binding together the API, master and APIServer infrastructure.
//
//nolint:nakedret
package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	extensionsapiserver "k8s.io/apiextensions-apiserver/pkg/apiserver"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/admission"
	genericapifilters "k8s.io/apiserver/pkg/endpoints/filters"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/apiserver/pkg/util/notfoundhandler"
	"k8s.io/apiserver/pkg/util/webhook"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/rest"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"
	aggregatorscheme "k8s.io/kube-aggregator/pkg/apiserver/scheme"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/features"

	"github.com/rosas/onex/cmd/onex-apiserver/app/options"
	"github.com/rosas/onex/internal/controlplane"
	controlplaneapiserver "github.com/rosas/onex/internal/controlplane/apiserver"
	"github.com/rosas/onex/pkg/apiserver/storage"
	"github.com/rosas/onex/pkg/version"
)

const appName = "onex-apiserver"

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
}

type Option func(*options.ServerRunOptions)
type RegisterFunc func(plugins *admission.Plugins)

// WithLegacyCode returns an Option that sets the external group versions in the ServerRunOptions.
func WithEtcdOptions(prefix string, versions ...schema.GroupVersion) Option {
	return func(s *options.ServerRunOptions) {
		codec := legacyscheme.Codecs.LegacyCodec(versions...)
		s.RecommendedOptions.Etcd = genericoptions.NewEtcdOptions(storagebackend.NewDefaultConfig(prefix, codec))
		// Note: DefaultStorageMediaType needs to be reset here.
		s.RecommendedOptions.Etcd.DefaultStorageMediaType = "application/vnd.kubernetes.protobuf"
		/*
			s.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(
				v1beta1.SchemeGroupVersion,
				schema.GroupKind{Group: v1beta1.GroupName},
			)
		*/
		for _, version := range versions {
			controlplane.AddStableAPIGroupVersionsEnabledByDefault(version)
		}
	}
}

// WithAdmissionPlugin returns an Option function that adds an admission plugin to the recommended plugin order list
// and registers the plugin using the provided RegisterFunc in the ServerRunOptions.
func WithAdmissionPlugin(name string, registerFunc RegisterFunc) Option {
	return func(s *options.ServerRunOptions) {
		// Note: Need to add this to the RecommendedPluginOrder list.
		s.RecommendedOptions.Admission.RecommendedPluginOrder = append(s.RecommendedOptions.Admission.RecommendedPluginOrder, name)
		registerFunc(s.RecommendedOptions.Admission.Plugins)
	}
}

// WithOpenAPIDefinitions sets the OpenAPI definitions for the server.
func WithGetOpenAPIDefinitions(getOpenAPIDefinitions common.GetOpenAPIDefinitions) Option {
	return func(s *options.ServerRunOptions) {
		s.GetOpenAPIDefinitions = getOpenAPIDefinitions
	}
}

// WithAdmissionInitializers returns an Option function that sets the external admission initializers in the ServerRunOptions.
func WithAdmissionInitializers(initializer func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error)) Option {
	return func(s *options.ServerRunOptions) {
		s.RecommendedOptions.ExternalAdmissionInitializers = initializer
	}
}

// WithPostStartHook returns an Option function that sets the external post-start hook with the given name in the ServerRunOptions.
func WithPostStartHook(name string, hook genericapiserver.PostStartHookFunc) Option {
	return func(s *options.ServerRunOptions) {
		s.ExternalPostStartHooks[name] = hook
	}
}

// WithSharedInformerFactory returns an Option function that sets the external SharedInformerFactory in the ServerRunOptions.
func WithSharedInformerFactory(informers controlplane.ExternalSharedInformerFactory) Option {
	return func(s *options.ServerRunOptions) {
		s.ExternalVersionedInformers = informers
	}
}

// WithRESTStorageProviders returns an Option function that sets the external REST storage providers in the ServerRunOptions.
func WithRESTStorageProviders(providers ...storage.RESTStorageProvider) Option {
	return func(s *options.ServerRunOptions) {
		s.ExternalRESTStorageProviders = providers
	}
}

// WithAlternateDNS returns an Option function that sets the alternate DNS configurations in the ServerRunOptions.
func WithAlternateDNS(dns ...string) Option {
	return func(s *options.ServerRunOptions) {
		s.AlternateDNS = dns
	}
}

// NewAPIServerCommand creates a *cobra.Command object with default parameters.
func NewAPIServerCommand(serverRunOptions ...Option) *cobra.Command {
	s := options.NewServerRunOptions()
	for _, opt := range serverRunOptions {
		opt(s)
	}

	cmd := &cobra.Command{
		Use:   appName,
		Short: "Launch a onex API server",
		Long: `The OneX API server validates and configures data
for the api objects which include miners, minersets, configmaps, and
others. The API Server services REST operations and provides the frontend to the
onex's shared state through which all other components interact.`,

		// stop printing usage when the command errors
		SilenceUsage: true,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			// silence client-go warnings.
			// onex-apiserver loopback clients should not log self-issued warnings.
			rest.SetDefaultWarningHandler(rest.NoWarnings{})
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			version.PrintAndExitIfRequested(appName)
			fs := cmd.Flags()

			// Activate logging as soon as possible, after that
			// show flags with the final logging configuration.
			if err := logsapi.ValidateAndApply(s.Logs, utilfeature.DefaultFeatureGate); err != nil {
				return err
			}
			cliflag.PrintFlags(fs)

			// set default options
			completedOptions, err := s.Complete()
			if err != nil {
				return err
			}

			// validate options
			if errs := completedOptions.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}
			// add feature enablement metrics
			utilfeature.DefaultMutableFeatureGate.AddMetrics()
			return Run(completedOptions, genericapiserver.SetupSignalHandler())
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	fs := cmd.Flags()
	namedFlagSets := s.Flags()
	version.AddFlags(namedFlagSets.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name(), logs.SkipLoggingConfigurationFlags())
	// The custom flag is actually not used. It is just a placeholder. In order to be consistent with
	// the kube-apiserver code, learning onex-apiserver is equivalent to learning kube-apiserver.
	options.AddCustomGlobalFlags(namedFlagSets.FlagSet("generic"))
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, namedFlagSets, cols)

	return cmd
}

// Run runs the specified APIServer. This should never exit.
func Run(opts options.CompletedOptions, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	klog.Infof("Version: %+v", version.Get().String())

	klog.InfoS("Golang settings", "GOGC", os.Getenv("GOGC"), "GOMAXPROCS", os.Getenv("GOMAXPROCS"), "GOTRACEBACK", os.Getenv("GOTRACEBACK"))

	config, err := NewConfig(opts)
	if err != nil {
		return err
	}
	completed, err := config.Complete()
	if err != nil {
		return err
	}
	server, err := CreateServerChain(completed)
	if err != nil {
		return err
	}

	prepared, err := server.PrepareRun()
	if err != nil {
		return err
	}

	return prepared.Run(stopCh)
}

// CreateServerChain creates the apiservers connected via delegation.
func CreateServerChain(config CompletedConfig) (*aggregatorapiserver.APIAggregator, error) {
	notFoundHandler := notfoundhandler.New(config.ControlPlane.GenericConfig.Serializer, genericapifilters.NoMuxAndDiscoveryIncompleteKey)
	apiExtensionsServer, err := config.ApiExtensions.New(genericapiserver.NewEmptyDelegateWithCustomHandler(notFoundHandler))
	if err != nil {
		return nil, err
	}
	crdAPIEnabled := config.ApiExtensions.GenericConfig.MergedResourceConfig.ResourceEnabled(apiextensionsv1.SchemeGroupVersion.WithResource("customresourcedefinitions"))

	onexAPIServer, err := config.ControlPlane.New(apiExtensionsServer.GenericAPIServer)
	if err != nil {
		return nil, err
	}

	// aggregator comes last in the chain
	aggregatorServer, err := createAggregatorServer(config.Aggregator, onexAPIServer.GenericAPIServer, apiExtensionsServer.Informers, crdAPIEnabled)
	if err != nil {
		// we don't need special handling for innerStopCh because the aggregator server doesn't create any go routines
		return nil, err
	}

	return aggregatorServer, nil
}

// CreateProxyTransport creates the dialer infrastructure to connect to the nodes.
func CreateProxyTransport() *http.Transport {
	var proxyDialerFn utilnet.DialFunc
	// Proxying to pods and services is IP-based... don't expect to be able to verify the hostname
	proxyTLSClientConfig := &tls.Config{InsecureSkipVerify: true}
	proxyTransport := utilnet.SetTransportDefaults(&http.Transport{
		DialContext:     proxyDialerFn,
		TLSClientConfig: proxyTLSClientConfig,
	})
	return proxyTransport
}

// CreateOneXAPIServerConfig creates all the resources for running kube-apiserver, but runs none of them.
func CreateOneXAPIServerConfig(opts options.CompletedOptions) (
	*controlplane.Config,
	aggregatorapiserver.ServiceResolver,
	error,
) {
	proxyTransport := CreateProxyTransport()

	genericConfig, _, kubeSharedInformers, storageFactory, err := controlplaneapiserver.BuildGenericConfig(
		opts.CompletedOptions,
		[]*runtime.Scheme{legacyscheme.Scheme, extensionsapiserver.Scheme, aggregatorscheme.Scheme},
		opts.GetOpenAPIDefinitions,
	)
	if err != nil {
		return nil, nil, err
	}

	opts.Metrics.Apply()

	config := &controlplane.Config{
		GenericConfig: genericConfig,
		ExtraConfig: controlplane.ExtraConfig{
			APIResourceConfigSource: storageFactory.APIResourceConfigSource,
			StorageFactory:          storageFactory,
			EventTTL:                opts.EventTTL,
			EnableLogsSupport:       opts.EnableLogsHandler,
			ProxyTransport:          proxyTransport,
			//ExternalGroupResources: opts.ExternalGroupResources,
			ExternalRESTStorageProviders: opts.ExternalRESTStorageProviders,
			MasterCount:                  opts.MasterCount,
			//VersionedInformers:           opts.SharedInformerFactory,
			// Here we will use the config file of "onex" to create a client-go informers.
			KubeVersionedInformers:     kubeSharedInformers,
			InternalVersionedInformers: opts.InternalVersionedInformers,
			ExternalVersionedInformers: opts.ExternalVersionedInformers,
			ExternalPostStartHooks:     opts.ExternalPostStartHooks,
		},
	}

	if utilfeature.DefaultFeatureGate.Enabled(features.UnknownVersionInteroperabilityProxy) {
		config.ExtraConfig.PeerEndpointLeaseReconciler, err = controlplaneapiserver.CreatePeerEndpointLeaseReconciler(genericConfig.Config, storageFactory)
		if err != nil {
			return nil, nil, err
		}
		// build peer proxy config only if peer ca file exists
		if opts.PeerCAFile != "" {
			config.ExtraConfig.PeerProxy, err = controlplaneapiserver.BuildPeerProxy(
				kubeSharedInformers,
				genericConfig.StorageVersionManager,
				opts.ProxyClientCertFile,
				opts.ProxyClientKeyFile,
				opts.PeerCAFile,
				opts.PeerAdvertiseAddress,
				genericConfig.APIServerID,
				config.ExtraConfig.PeerEndpointLeaseReconciler,
				config.GenericConfig.Serializer,
			)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	/* UPDATEME: when add authentication features.
	clientCAProvider, err := opts.Authentication.ClientCert.GetClientCAContentProvider()
	if err != nil {
		return nil, nil, nil, err
	}
	config.ExtraConfig.ClusterAuthenticationInfo.ClientCA = clientCAProvider

	requestHeaderConfig, err := opts.Authentication.RequestHeader.ToAuthenticationRequestHeaderConfig()
	if err != nil {
		return nil, nil, nil, err
	}
	if requestHeaderConfig != nil {
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderCA = requestHeaderConfig.CAContentProvider
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderAllowedNames = requestHeaderConfig.AllowedClientNames
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderExtraHeaderPrefixes = requestHeaderConfig.ExtraHeaderPrefixes
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderGroupHeaders = requestHeaderConfig.GroupHeaders
		config.ExtraConfig.ClusterAuthenticationInfo.RequestHeaderUsernameHeaders = requestHeaderConfig.UsernameHeaders
	}
	*/

	serviceResolver := buildServiceResolver(opts.EnableAggregatorRouting, genericConfig.LoopbackClientConfig.Host, kubeSharedInformers)

	return config, serviceResolver, nil
}

var testServiceResolver webhook.ServiceResolver

func buildServiceResolver(enabledAggregatorRouting bool, hostname string, informer kubeinformers.SharedInformerFactory) webhook.ServiceResolver {
	if testServiceResolver != nil {
		return testServiceResolver
	}

	var serviceResolver webhook.ServiceResolver
	if enabledAggregatorRouting {
		serviceResolver = aggregatorapiserver.NewEndpointServiceResolver(
			informer.Core().V1().Services().Lister(),
			informer.Core().V1().Endpoints().Lister(),
		)
	} else {
		serviceResolver = aggregatorapiserver.NewClusterIPServiceResolver(
			informer.Core().V1().Services().Lister(),
		)
	}

	// resolve kubernetes.default.svc locally
	if localHost, err := url.Parse(hostname); err == nil {
		serviceResolver = aggregatorapiserver.NewLoopbackServiceResolver(serviceResolver, localHost)
	}
	return serviceResolver
}
