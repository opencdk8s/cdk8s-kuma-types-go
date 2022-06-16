// @opencdk8s/cdk8s-kuma-types
package opencdk8scdk8skumatypes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-kuma-types-go/opencdk8scdk8skumatypes/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/opencdk8s/cdk8s-kuma-types-go/opencdk8scdk8skumatypes/internal"
)

// CircuitBreaker is the Schema for the circuitbreaker API.
type CircuitBreaker interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for CircuitBreaker
type jsiiProxy_CircuitBreaker struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_CircuitBreaker) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CircuitBreaker) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CircuitBreaker) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CircuitBreaker) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CircuitBreaker) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CircuitBreaker) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CircuitBreaker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "CircuitBreaker" API object.
func NewCircuitBreaker(scope constructs.Construct, id *string, props *CircuitBreakerProps) CircuitBreaker {
	_init_.Initialize()

	j := jsiiProxy_CircuitBreaker{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.CircuitBreaker",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "CircuitBreaker" API object.
func NewCircuitBreaker_Override(c CircuitBreaker, scope constructs.Construct, id *string, props *CircuitBreakerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.CircuitBreaker",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CircuitBreaker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.CircuitBreaker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "CircuitBreaker".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func CircuitBreaker_Manifest(props *CircuitBreakerProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.CircuitBreaker",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func CircuitBreaker_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.CircuitBreaker",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func CircuitBreaker_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.CircuitBreaker",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_CircuitBreaker) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (c *jsiiProxy_CircuitBreaker) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (c *jsiiProxy_CircuitBreaker) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CircuitBreaker) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// CircuitBreaker is the Schema for the circuitbreaker API.
type CircuitBreakerProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Dataplane is the Schema for the dataplanes API.
type Dataplane interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Dataplane
type jsiiProxy_Dataplane struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Dataplane) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataplane) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataplane) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataplane) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataplane) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataplane) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dataplane) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "Dataplane" API object.
func NewDataplane(scope constructs.Construct, id *string, props *DataplaneProps) Dataplane {
	_init_.Initialize()

	j := jsiiProxy_Dataplane{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Dataplane",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Dataplane" API object.
func NewDataplane_Override(d Dataplane, scope constructs.Construct, id *string, props *DataplaneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Dataplane",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Dataplane_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Dataplane",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "Dataplane".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Dataplane_Manifest(props *DataplaneProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Dataplane",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Dataplane_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Dataplane",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Dataplane_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.Dataplane",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (d *jsiiProxy_Dataplane) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (d *jsiiProxy_Dataplane) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (d *jsiiProxy_Dataplane) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_Dataplane) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// DataplaneInsight is the Schema for the dataplane insights API.
type DataplaneInsight interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for DataplaneInsight
type jsiiProxy_DataplaneInsight struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_DataplaneInsight) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplaneInsight) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplaneInsight) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplaneInsight) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplaneInsight) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplaneInsight) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplaneInsight) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "DataplaneInsight" API object.
func NewDataplaneInsight(scope constructs.Construct, id *string, props *DataplaneInsightProps) DataplaneInsight {
	_init_.Initialize()

	j := jsiiProxy_DataplaneInsight{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.DataplaneInsight",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "DataplaneInsight" API object.
func NewDataplaneInsight_Override(d DataplaneInsight, scope constructs.Construct, id *string, props *DataplaneInsightProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.DataplaneInsight",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataplaneInsight_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.DataplaneInsight",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "DataplaneInsight".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func DataplaneInsight_Manifest(props *DataplaneInsightProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.DataplaneInsight",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func DataplaneInsight_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.DataplaneInsight",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func DataplaneInsight_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.DataplaneInsight",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (d *jsiiProxy_DataplaneInsight) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (d *jsiiProxy_DataplaneInsight) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (d *jsiiProxy_DataplaneInsight) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataplaneInsight) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// DataplaneInsight is the Schema for the dataplane insights API.
type DataplaneInsightProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
}

// Dataplane is the Schema for the dataplanes API.
type DataplaneProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

type ExternalService interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ExternalService
type jsiiProxy_ExternalService struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ExternalService) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "ExternalService" API object.
func NewExternalService(scope constructs.Construct, id *string, props *ExternalServiceProps) ExternalService {
	_init_.Initialize()

	j := jsiiProxy_ExternalService{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ExternalService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ExternalService" API object.
func NewExternalService_Override(e ExternalService, scope constructs.Construct, id *string, props *ExternalServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ExternalService",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ExternalService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ExternalService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ExternalService".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ExternalService_Manifest(props *ExternalServiceProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ExternalService",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ExternalService_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ExternalService",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ExternalService_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.ExternalService",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (e *jsiiProxy_ExternalService) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (e *jsiiProxy_ExternalService) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (e *jsiiProxy_ExternalService) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ExternalService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ExternalServiceProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// FaultInjection is the Schema for the faultinjections API.
type FaultInjection interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for FaultInjection
type jsiiProxy_FaultInjection struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_FaultInjection) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FaultInjection) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FaultInjection) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FaultInjection) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FaultInjection) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FaultInjection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FaultInjection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "FaultInjection" API object.
func NewFaultInjection(scope constructs.Construct, id *string, props *FaultInjectionProps) FaultInjection {
	_init_.Initialize()

	j := jsiiProxy_FaultInjection{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.FaultInjection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "FaultInjection" API object.
func NewFaultInjection_Override(f FaultInjection, scope constructs.Construct, id *string, props *FaultInjectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.FaultInjection",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FaultInjection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.FaultInjection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "FaultInjection".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func FaultInjection_Manifest(props *FaultInjectionProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.FaultInjection",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func FaultInjection_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.FaultInjection",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func FaultInjection_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.FaultInjection",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (f *jsiiProxy_FaultInjection) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (f *jsiiProxy_FaultInjection) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (f *jsiiProxy_FaultInjection) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FaultInjection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// FaultInjection is the Schema for the faultinjections API.
type FaultInjectionProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// HealthCheck is the Schema for the healthchecks API.
type HealthCheck interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_HealthCheck) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "HealthCheck" API object.
func NewHealthCheck(scope constructs.Construct, id *string, props *HealthCheckProps) HealthCheck {
	_init_.Initialize()

	j := jsiiProxy_HealthCheck{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.HealthCheck",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "HealthCheck" API object.
func NewHealthCheck_Override(h HealthCheck, scope constructs.Construct, id *string, props *HealthCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.HealthCheck",
		[]interface{}{scope, id, props},
		h,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func HealthCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.HealthCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "HealthCheck".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func HealthCheck_Manifest(props *HealthCheckProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.HealthCheck",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func HealthCheck_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.HealthCheck",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func HealthCheck_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.HealthCheck",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (h *jsiiProxy_HealthCheck) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		h,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (h *jsiiProxy_HealthCheck) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		h,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (h *jsiiProxy_HealthCheck) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (h *jsiiProxy_HealthCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// HealthCheck is the Schema for the healthchecks API.
type HealthCheckProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Mesh is the Schema for the meshes API.
type Mesh interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Mesh
type jsiiProxy_Mesh struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Mesh) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "Mesh" API object.
func NewMesh(scope constructs.Construct, id *string, props *MeshProps) Mesh {
	_init_.Initialize()

	j := jsiiProxy_Mesh{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Mesh",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Mesh" API object.
func NewMesh_Override(m Mesh, scope constructs.Construct, id *string, props *MeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Mesh",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Mesh_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Mesh",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "Mesh".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Mesh_Manifest(props *MeshProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Mesh",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Mesh_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Mesh",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Mesh_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.Mesh",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (m *jsiiProxy_Mesh) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (m *jsiiProxy_Mesh) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (m *jsiiProxy_Mesh) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_Mesh) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// MeshInsight is the Schema for the meshes insights API.
type MeshInsight interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for MeshInsight
type jsiiProxy_MeshInsight struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_MeshInsight) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MeshInsight) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MeshInsight) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MeshInsight) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MeshInsight) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MeshInsight) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MeshInsight) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "MeshInsight" API object.
func NewMeshInsight(scope constructs.Construct, id *string, props *MeshInsightProps) MeshInsight {
	_init_.Initialize()

	j := jsiiProxy_MeshInsight{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.MeshInsight",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "MeshInsight" API object.
func NewMeshInsight_Override(m MeshInsight, scope constructs.Construct, id *string, props *MeshInsightProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.MeshInsight",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func MeshInsight_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.MeshInsight",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "MeshInsight".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func MeshInsight_Manifest(props *MeshInsightProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.MeshInsight",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func MeshInsight_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.MeshInsight",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func MeshInsight_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.MeshInsight",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (m *jsiiProxy_MeshInsight) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (m *jsiiProxy_MeshInsight) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (m *jsiiProxy_MeshInsight) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MeshInsight) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// MeshInsight is the Schema for the meshes insights API.
type MeshInsightProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Mesh is the Schema for the meshes API.
type MeshProps struct {
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// ProxyTemplate is the Schema for the proxytemplates API.
type ProxyTemplate interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ProxyTemplate
type jsiiProxy_ProxyTemplate struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ProxyTemplate) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyTemplate) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyTemplate) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyTemplate) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyTemplate) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProxyTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "ProxyTemplate" API object.
func NewProxyTemplate(scope constructs.Construct, id *string, props *ProxyTemplateProps) ProxyTemplate {
	_init_.Initialize()

	j := jsiiProxy_ProxyTemplate{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ProxyTemplate",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ProxyTemplate" API object.
func NewProxyTemplate_Override(p ProxyTemplate, scope constructs.Construct, id *string, props *ProxyTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ProxyTemplate",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ProxyTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ProxyTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ProxyTemplate".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ProxyTemplate_Manifest(props *ProxyTemplateProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ProxyTemplate",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ProxyTemplate_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ProxyTemplate",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ProxyTemplate_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.ProxyTemplate",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (p *jsiiProxy_ProxyTemplate) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (p *jsiiProxy_ProxyTemplate) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (p *jsiiProxy_ProxyTemplate) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_ProxyTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ProxyTemplate is the Schema for the proxytemplates API.
type ProxyTemplateProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// RateLimit is the Schema for the ratelimits API.
type RateLimit interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for RateLimit
type jsiiProxy_RateLimit struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_RateLimit) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimit) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimit) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimit) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimit) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimit) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RateLimit) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "RateLimit" API object.
func NewRateLimit(scope constructs.Construct, id *string, props *RateLimitProps) RateLimit {
	_init_.Initialize()

	j := jsiiProxy_RateLimit{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.RateLimit",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "RateLimit" API object.
func NewRateLimit_Override(r RateLimit, scope constructs.Construct, id *string, props *RateLimitProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.RateLimit",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func RateLimit_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.RateLimit",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "RateLimit".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func RateLimit_Manifest(props *RateLimitProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.RateLimit",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func RateLimit_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.RateLimit",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func RateLimit_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.RateLimit",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (r *jsiiProxy_RateLimit) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (r *jsiiProxy_RateLimit) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (r *jsiiProxy_RateLimit) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_RateLimit) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// RateLimit is the Schema for the ratelimits API.
type RateLimitProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Retry is the Schema for the retries API.
type Retry interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Retry
type jsiiProxy_Retry struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Retry) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Retry) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Retry) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Retry) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Retry) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Retry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Retry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "Retry" API object.
func NewRetry(scope constructs.Construct, id *string, props *RetryProps) Retry {
	_init_.Initialize()

	j := jsiiProxy_Retry{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Retry",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Retry" API object.
func NewRetry_Override(r Retry, scope constructs.Construct, id *string, props *RetryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Retry",
		[]interface{}{scope, id, props},
		r,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Retry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Retry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "Retry".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Retry_Manifest(props *RetryProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Retry",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Retry_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Retry",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Retry_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.Retry",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (r *jsiiProxy_Retry) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (r *jsiiProxy_Retry) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		r,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (r *jsiiProxy_Retry) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_Retry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Retry is the Schema for the retries API.
type RetryProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// ServiceInsight is the Schema for the services insights API.
type ServiceInsight interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ServiceInsight
type jsiiProxy_ServiceInsight struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ServiceInsight) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceInsight) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceInsight) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceInsight) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceInsight) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceInsight) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceInsight) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "ServiceInsight" API object.
func NewServiceInsight(scope constructs.Construct, id *string, props *ServiceInsightProps) ServiceInsight {
	_init_.Initialize()

	j := jsiiProxy_ServiceInsight{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ServiceInsight",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ServiceInsight" API object.
func NewServiceInsight_Override(s ServiceInsight, scope constructs.Construct, id *string, props *ServiceInsightProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ServiceInsight",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ServiceInsight_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ServiceInsight",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ServiceInsight".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ServiceInsight_Manifest(props *ServiceInsightProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ServiceInsight",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ServiceInsight_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ServiceInsight",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ServiceInsight_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.ServiceInsight",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (s *jsiiProxy_ServiceInsight) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (s *jsiiProxy_ServiceInsight) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (s *jsiiProxy_ServiceInsight) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_ServiceInsight) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ServiceInsight is the Schema for the services insights API.
type ServiceInsightProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Timeout is the Schema for the timeout API.
type Timeout interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Timeout
type jsiiProxy_Timeout struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Timeout) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "Timeout" API object.
func NewTimeout(scope constructs.Construct, id *string, props *TimeoutProps) Timeout {
	_init_.Initialize()

	j := jsiiProxy_Timeout{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Timeout",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Timeout" API object.
func NewTimeout_Override(t Timeout, scope constructs.Construct, id *string, props *TimeoutProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Timeout",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Timeout_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Timeout",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "Timeout".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Timeout_Manifest(props *TimeoutProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Timeout",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Timeout_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Timeout",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Timeout_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.Timeout",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (t *jsiiProxy_Timeout) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (t *jsiiProxy_Timeout) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (t *jsiiProxy_Timeout) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_Timeout) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Timeout is the Schema for the timeout API.
type TimeoutProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// TrafficLog is the Schema for the trafficlogs API.
type TrafficLog interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for TrafficLog
type jsiiProxy_TrafficLog struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_TrafficLog) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficLog) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficLog) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficLog) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficLog) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficLog) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficLog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "TrafficLog" API object.
func NewTrafficLog(scope constructs.Construct, id *string, props *TrafficLogProps) TrafficLog {
	_init_.Initialize()

	j := jsiiProxy_TrafficLog{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficLog",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "TrafficLog" API object.
func NewTrafficLog_Override(t TrafficLog, scope constructs.Construct, id *string, props *TrafficLogProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficLog",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TrafficLog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficLog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "TrafficLog".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func TrafficLog_Manifest(props *TrafficLogProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficLog",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func TrafficLog_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficLog",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func TrafficLog_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.TrafficLog",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (t *jsiiProxy_TrafficLog) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (t *jsiiProxy_TrafficLog) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (t *jsiiProxy_TrafficLog) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TrafficLog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TrafficLog is the Schema for the trafficlogs API.
type TrafficLogProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// TrafficPermission is the Schema for the trafficpermissions API.
type TrafficPermission interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for TrafficPermission
type jsiiProxy_TrafficPermission struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_TrafficPermission) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficPermission) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficPermission) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficPermission) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficPermission) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficPermission) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "TrafficPermission" API object.
func NewTrafficPermission(scope constructs.Construct, id *string, props *TrafficPermissionProps) TrafficPermission {
	_init_.Initialize()

	j := jsiiProxy_TrafficPermission{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficPermission",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "TrafficPermission" API object.
func NewTrafficPermission_Override(t TrafficPermission, scope constructs.Construct, id *string, props *TrafficPermissionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficPermission",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TrafficPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "TrafficPermission".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func TrafficPermission_Manifest(props *TrafficPermissionProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficPermission",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func TrafficPermission_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficPermission",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func TrafficPermission_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.TrafficPermission",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (t *jsiiProxy_TrafficPermission) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (t *jsiiProxy_TrafficPermission) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (t *jsiiProxy_TrafficPermission) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TrafficPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TrafficPermission is the Schema for the trafficpermissions API.
type TrafficPermissionProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// TrafficRoute is the Schema for the trafficroutes API.
type TrafficRoute interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for TrafficRoute
type jsiiProxy_TrafficRoute struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_TrafficRoute) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficRoute) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficRoute) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficRoute) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficRoute) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficRoute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "TrafficRoute" API object.
func NewTrafficRoute(scope constructs.Construct, id *string, props *TrafficRouteProps) TrafficRoute {
	_init_.Initialize()

	j := jsiiProxy_TrafficRoute{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "TrafficRoute" API object.
func NewTrafficRoute_Override(t TrafficRoute, scope constructs.Construct, id *string, props *TrafficRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficRoute",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TrafficRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "TrafficRoute".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func TrafficRoute_Manifest(props *TrafficRouteProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficRoute",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func TrafficRoute_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficRoute",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func TrafficRoute_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.TrafficRoute",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (t *jsiiProxy_TrafficRoute) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (t *jsiiProxy_TrafficRoute) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (t *jsiiProxy_TrafficRoute) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TrafficRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TrafficRoute is the Schema for the trafficroutes API.
type TrafficRouteProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// TrafficTrace is the Schema for the traffictraces API.
type TrafficTrace interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for TrafficTrace
type jsiiProxy_TrafficTrace struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_TrafficTrace) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficTrace) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficTrace) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficTrace) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficTrace) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficTrace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrafficTrace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "TrafficTrace" API object.
func NewTrafficTrace(scope constructs.Construct, id *string, props *TrafficTraceProps) TrafficTrace {
	_init_.Initialize()

	j := jsiiProxy_TrafficTrace{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficTrace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "TrafficTrace" API object.
func NewTrafficTrace_Override(t TrafficTrace, scope constructs.Construct, id *string, props *TrafficTraceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.TrafficTrace",
		[]interface{}{scope, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TrafficTrace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficTrace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "TrafficTrace".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func TrafficTrace_Manifest(props *TrafficTraceProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficTrace",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func TrafficTrace_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.TrafficTrace",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func TrafficTrace_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.TrafficTrace",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (t *jsiiProxy_TrafficTrace) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (t *jsiiProxy_TrafficTrace) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (t *jsiiProxy_TrafficTrace) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TrafficTrace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TrafficTrace is the Schema for the traffictraces API.
type TrafficTraceProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// VirtualOutbound is the Schema for the virtualoutbounds API.
type VirtualOutbound interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for VirtualOutbound
type jsiiProxy_VirtualOutbound struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_VirtualOutbound) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualOutbound) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualOutbound) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualOutbound) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualOutbound) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualOutbound) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualOutbound) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "VirtualOutbound" API object.
func NewVirtualOutbound(scope constructs.Construct, id *string, props *VirtualOutboundProps) VirtualOutbound {
	_init_.Initialize()

	j := jsiiProxy_VirtualOutbound{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.VirtualOutbound",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "VirtualOutbound" API object.
func NewVirtualOutbound_Override(v VirtualOutbound, scope constructs.Construct, id *string, props *VirtualOutboundProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.VirtualOutbound",
		[]interface{}{scope, id, props},
		v,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func VirtualOutbound_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.VirtualOutbound",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "VirtualOutbound".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func VirtualOutbound_Manifest(props *VirtualOutboundProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.VirtualOutbound",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func VirtualOutbound_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.VirtualOutbound",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func VirtualOutbound_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.VirtualOutbound",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (v *jsiiProxy_VirtualOutbound) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (v *jsiiProxy_VirtualOutbound) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (v *jsiiProxy_VirtualOutbound) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (v *jsiiProxy_VirtualOutbound) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// VirtualOutbound is the Schema for the virtualoutbounds API.
type VirtualOutboundProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Zone is the Schema for the zone API.
type Zone interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for Zone
type jsiiProxy_Zone struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_Zone) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zone) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zone) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zone) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zone) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zone) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "Zone" API object.
func NewZone(scope constructs.Construct, id *string, props *ZoneProps) Zone {
	_init_.Initialize()

	j := jsiiProxy_Zone{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Zone",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "Zone" API object.
func NewZone_Override(z Zone, scope constructs.Construct, id *string, props *ZoneProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.Zone",
		[]interface{}{scope, id, props},
		z,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Zone_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Zone",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "Zone".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func Zone_Manifest(props *ZoneProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Zone",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func Zone_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.Zone",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func Zone_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.Zone",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (z *jsiiProxy_Zone) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (z *jsiiProxy_Zone) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (z *jsiiProxy_Zone) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (z *jsiiProxy_Zone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ZoneIngress is the Schema for the zone ingress API.
type ZoneIngress interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ZoneIngress
type jsiiProxy_ZoneIngress struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ZoneIngress) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngress) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngress) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngress) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngress) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngress) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "ZoneIngress" API object.
func NewZoneIngress(scope constructs.Construct, id *string, props *ZoneIngressProps) ZoneIngress {
	_init_.Initialize()

	j := jsiiProxy_ZoneIngress{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngress",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ZoneIngress" API object.
func NewZoneIngress_Override(z ZoneIngress, scope constructs.Construct, id *string, props *ZoneIngressProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngress",
		[]interface{}{scope, id, props},
		z,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ZoneIngress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ZoneIngress".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ZoneIngress_Manifest(props *ZoneIngressProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngress",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ZoneIngress_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngress",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ZoneIngress_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngress",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (z *jsiiProxy_ZoneIngress) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (z *jsiiProxy_ZoneIngress) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (z *jsiiProxy_ZoneIngress) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (z *jsiiProxy_ZoneIngress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ZoneIngressInsight is the Schema for the zone ingress insight API.
type ZoneIngressInsight interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ZoneIngressInsight
type jsiiProxy_ZoneIngressInsight struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ZoneIngressInsight) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngressInsight) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngressInsight) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngressInsight) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngressInsight) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngressInsight) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneIngressInsight) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "ZoneIngressInsight" API object.
func NewZoneIngressInsight(scope constructs.Construct, id *string, props *ZoneIngressInsightProps) ZoneIngressInsight {
	_init_.Initialize()

	j := jsiiProxy_ZoneIngressInsight{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngressInsight",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ZoneIngressInsight" API object.
func NewZoneIngressInsight_Override(z ZoneIngressInsight, scope constructs.Construct, id *string, props *ZoneIngressInsightProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngressInsight",
		[]interface{}{scope, id, props},
		z,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ZoneIngressInsight_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngressInsight",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ZoneIngressInsight".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ZoneIngressInsight_Manifest(props *ZoneIngressInsightProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngressInsight",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ZoneIngressInsight_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngressInsight",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ZoneIngressInsight_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.ZoneIngressInsight",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (z *jsiiProxy_ZoneIngressInsight) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (z *jsiiProxy_ZoneIngressInsight) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (z *jsiiProxy_ZoneIngressInsight) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (z *jsiiProxy_ZoneIngressInsight) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ZoneIngressInsight is the Schema for the zone ingress insight API.
type ZoneIngressInsightProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// ZoneIngress is the Schema for the zone ingress API.
type ZoneIngressProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// ZoneInsight is the Schema for the zone insight API.
type ZoneInsight interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	Node() constructs.Node
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ZoneInsight
type jsiiProxy_ZoneInsight struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ZoneInsight) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneInsight) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneInsight) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneInsight) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneInsight) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneInsight) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneInsight) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Defines a "ZoneInsight" API object.
func NewZoneInsight(scope constructs.Construct, id *string, props *ZoneInsightProps) ZoneInsight {
	_init_.Initialize()

	j := jsiiProxy_ZoneInsight{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ZoneInsight",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ZoneInsight" API object.
func NewZoneInsight_Override(z ZoneInsight, scope constructs.Construct, id *string, props *ZoneInsightProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-kuma-types.ZoneInsight",
		[]interface{}{scope, id, props},
		z,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ZoneInsight_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneInsight",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ZoneInsight".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ZoneInsight_Manifest(props *ZoneInsightProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneInsight",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ZoneInsight_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-kuma-types.ZoneInsight",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ZoneInsight_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-kuma-types.ZoneInsight",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (z *jsiiProxy_ZoneInsight) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (z *jsiiProxy_ZoneInsight) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		z,
		"addJsonPatch",
		args,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (z *jsiiProxy_ZoneInsight) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		z,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (z *jsiiProxy_ZoneInsight) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// ZoneInsight is the Schema for the zone insight API.
type ZoneInsightProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

// Zone is the Schema for the zone API.
type ZoneProps struct {
	Mesh *string `json:"mesh"`
	Metadata *cdk8s.ApiObjectMetadata `json:"metadata"`
	Spec interface{} `json:"spec"`
}

