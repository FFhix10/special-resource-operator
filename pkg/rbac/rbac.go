package rbac

// +kubebuilder:rbac:groups=sro.openshift.io,resources=specialresources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sro.openshift.io,resources=specialresources/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=sro.openshift.io,resources=specialresources/finalizers,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/log,verbs=get
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.openshift.io,resources=clusterversions,verbs=get
// +kubebuilder:rbac:groups=config.openshift.io,resources=proxies,verbs=get;list
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=security.openshift.io,resources=securitycontextconstraints,verbs=use;get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=image.openshift.io,resources=imagestreams,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=image.openshift.io,resources=imagestreams/finalizers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=image.openshift.io,resources=imagestreams/layers,verbs=get
// +kubebuilder:rbac:groups=core,resources=imagestreams/layers,verbs=get
// +kubebuilder:rbac:groups=build.openshift.io,resources=buildconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=build.openshift.io,resources=builds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=list;watch;create;update;patch;delete;get
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;update;
// +kubebuilder:rbac:groups=core,resources=persistentvolumes,verbs=get;list;watch;create;delete;update;patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=storage.k8s.io,resources=csinodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=storage.k8s.io,resources=storageclasses,verbs=watch;get;list
// +kubebuilder:rbac:groups=storage.k8s.io,resources=csidrivers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=endpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitoring.coreos.com,resources=servicemonitors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitoring.coreos.com,resources=prometheusrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=route.openshift.io,resources=routes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.openshift.io,resources=clusteroperators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.openshift.io,resources=clusteroperators/status,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers,verbs=get;list;watch;create;update;patch;delete;deletecollection
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete;deletecollection
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=create;patch;delete
// +kubebuilder:rbac:groups=core,resources=services/finalizers,verbs=create;delete;get;list;update;patch;delete;watch
// +kubebuilder:rbac:groups=apps,resources=deployments/finalizers,resourceNames=shipwright-build,verbs=update
// +kubebuilder:rbac:groups=apps,resources=replicasets,verbs=create;delete;get;list;patch;update;watch;get
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=shipwright.io,resources=*,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=shipwright.io,resources=buildruns,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=shipwright.io,resources=buildstrategies,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=shipwright.io,resources=clusterbuildstrategies,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=tekton.dev,resources=taskruns,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=tekton.dev,resources=tasks,verbs=create;delete;get;list;patch;update;watch
// +kubebuilder:rbac:groups=storage.k8s.io,resources=volumeattachments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=snapshot.storage.k8s.io,resources=volumesnapshotclasses,verbs=get;list;watch
// +kubebuilder:rbac:groups=snapshot.storage.k8s.io,resources=volumesnapshots,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=snapshot.storage.k8s.io,resources=volumesnapshotcontents,verbs=create;get;list;watch;update;delete
// +kubebuilder:rbac:groups=snapshot.storage.k8s.io,resources=volumesnapshots/status,verbs=create;get;list;watch;update;delete
// +kubebuilder:rbac:groups=snapshot.storage.k8s.io,resources=volumesnapshotcontents/status,verbs=create;get;list;watch;update;delete
// +kubebuilder:rbac:groups=csi.storage.k8s.io,resources=csidrivers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims/status,verbs=get;list;watch;create;delete;update;patch
// +kubebuilder:rbac:groups=operators.coreos.com,resources=operatorgroups,verbs=get;list;watch;create;delete;update;patch
// +kubebuilder:rbac:groups=operators.coreos.com,resources=subscriptions,verbs=get;list;watch;create;delete;update;patch
// +kubebuilder:rbac:groups=operator.cert-manager.io,resources=certmanagers,verbs=get;list;watch;create;delete;update;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=apiregistration.k8s.io,resources=apiservices,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=auditregistration.k8s.io,resources=auditsinks,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=cert-manager.io,resources=issuers/status,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers/status,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;update;list;watch;deletecollection
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=get;update;list;watch;delete
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/finalizers,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/status,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/finalizers,verbs=update
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates/status,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders,verbs=create;delete;get;list;watch;update;patch;deletecollection
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges,verbs=create;delete;get;list;watch;update;patch;deletecollection
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders/finalizers,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=orders/status,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges/finalizers,verbs=update
// +kubebuilder:rbac:groups=acme.cert-manager.io,resources=challenges/status,verbs=update
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;delete;update
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/finalizers,verbs=update
// +kubebuilder:rbac:groups=route.openshift.io,resources=routes/custom-host,verbs=create
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=create;patch;deletecollection
// +kubebuilder:rbac:groups=cert-manager.io,resources=signers,resourceNames=clusterissuers.cert-manager.io/*,verbs=approve
// +kubebuilder:rbac:groups=cert-manager.io,resources=signers,resourceNames=issuers.cert-manager.io/*,verbs=approve
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=get;list;watch;update
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/status,verbs=update
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,resourceNames=clusterissuers.cert-manager.io/*,verbs=sign
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,resourceNames=issuers.cert-manager.io/*,verbs=sign
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,resourceNames=cert-manager-cainjector-leader-election,verbs=patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,resourceNames=cert-manager-cainjector-election-core,verbs=patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,resourceNames=cert-manager-cainjector-leader-election-core,verbs=patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,resourceNames=cert-manager-controller,verbs=patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=create;get;list;watch;update;delete;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=create;get;list;watch;update;delete;patch
// +kubebuilder:rbac:groups=*,resources=cronjobs,verbs=get;delete;update;list;watch;patch
// +kubebuilder:rbac:groups=*,resources=daemonsets,verbs=get
// +kubebuilder:rbac:groups=*,resources=deployments,verbs=get
// +kubebuilder:rbac:groups=*,resources=imagepolicies,verbs=get;update;delete
// +kubebuilder:rbac:groups=*,resources=jobs,verbs=get;create;delete;update;list;watch;patch
// +kubebuilder:rbac:groups=*,resources=mutatingwebhookconfigurations,verbs=get
// +kubebuilder:rbac:groups=*,resources=pods,verbs=get
// +kubebuilder:rbac:groups=*,resources=replicacontrollers,verbs=get
// +kubebuilder:rbac:groups=*,resources=replicasets,verbs=get
// +kubebuilder:rbac:groups=*,resources=statefulsets,verbs=get
// +kubebuilder:rbac:groups=connaisseur.policy,resources=imagepolicies,verbs=create
// +kubebuilder:rbac:groups=admissionregistration.k8s.io/v1beta1,resources=mutatingwebhookconfigurations,verbs=create;delete;update;list
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=watch;list
// +kubebuilder:rbac:groups="",resources=nodes/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=nodes/status,verbs=update;patch
// +kubebuilder:rbac:groups="",resources=pods,verbs=deletecollection
// +kubebuilder:rbac:groups="",resources=podtemplates,verbs=list;watch;get;create;update
// +kubebuilder:rbac:groups="",resources=podtemplates/finalizers,verbs=update
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=list;watch;get;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=jobs/finalizers,verbs=update
// +kubebuilder:rbac:groups=extensions,resources=jobs,verbs=list;watch;get;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=httproutes,verbs=get;list;watch;create;update;delete
// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=gateways,verbs=get;list;watch
// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=gateways/finalizers,verbs=update
// +kubebuilder:rbac:groups=networking.x-k8s.io,resources=httproutes/finalisers,verbs=update
// +kubebuilder:rbac:groups=infoscale.veritas.com,resources=infoscaleclusters,verbs=update;patch;get;list
