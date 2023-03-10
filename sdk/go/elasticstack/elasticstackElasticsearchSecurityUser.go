// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticstack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds and updates users in the native realm. These users are commonly referred to as native users. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-elasticstack/sdk/go/elasticstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"env":    "testing",
//				"open":   false,
//				"number": 49,
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = elasticstack.NewElasticstackElasticsearchSecurityUser(ctx, "user", &elasticstack.ElasticstackElasticsearchSecurityUserArgs{
//				Username:     pulumi.String("testuser"),
//				PasswordHash: pulumi.String(fmt.Sprintf("$2a$10$rMZe6TdsUwBX/TA8vRDz0OLwKAZeCzXM4jT3tfCjpSTB8HoFuq8xO")),
//				Roles: pulumi.StringArray{
//					pulumi.String("kibana_user"),
//				},
//				Metadata: pulumi.String(json0),
//				ElasticsearchConnection: &elasticstack.ElasticstackElasticsearchSecurityUserElasticsearchConnectionArgs{
//					Endpoints: pulumi.StringArray{
//						pulumi.String("http://localhost:9200"),
//					},
//					Username: pulumi.String("elastic"),
//					Password: pulumi.String("changeme"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"env":    "testing",
//				"open":   false,
//				"number": 49,
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = elasticstack.NewElasticstackElasticsearchSecurityUser(ctx, "dev", &elasticstack.ElasticstackElasticsearchSecurityUserArgs{
//				Username: pulumi.String("devuser"),
//				Password: pulumi.String("1234567890"),
//				Roles: pulumi.StringArray{
//					pulumi.String("kibana_user"),
//				},
//				Metadata: pulumi.String(json1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import elasticstack:index/elasticstackElasticsearchSecurityUser:ElasticstackElasticsearchSecurityUser user <cluster_uuid>/elastic
//
// ```
type ElasticstackElasticsearchSecurityUser struct {
	pulumi.CustomResourceState

	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSecurityUserElasticsearchConnectionPtrOutput `pulumi:"elasticsearchConnection"`
	// The email of the user.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Specifies whether the user is enabled. The default value is true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The full name of the user.
	FullName pulumi.StringPtrOutput `pulumi:"fullName"`
	// Arbitrary metadata that you want to associate with the user.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The user???s password. Passwords must be at least 6 characters long.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
	PasswordHash pulumi.StringPtrOutput `pulumi:"passwordHash"`
	// A set of roles the user has. The roles determine the user???s access permissions. Default is [].
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewElasticstackElasticsearchSecurityUser registers a new resource with the given unique name, arguments, and options.
func NewElasticstackElasticsearchSecurityUser(ctx *pulumi.Context,
	name string, args *ElasticstackElasticsearchSecurityUserArgs, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchSecurityUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PasswordHash != nil {
		args.PasswordHash = pulumi.ToSecret(args.PasswordHash).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"passwordHash",
	})
	opts = append(opts, secrets)
	var resource ElasticstackElasticsearchSecurityUser
	err := ctx.RegisterResource("elasticstack:index/elasticstackElasticsearchSecurityUser:ElasticstackElasticsearchSecurityUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticstackElasticsearchSecurityUser gets an existing ElasticstackElasticsearchSecurityUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticstackElasticsearchSecurityUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticstackElasticsearchSecurityUserState, opts ...pulumi.ResourceOption) (*ElasticstackElasticsearchSecurityUser, error) {
	var resource ElasticstackElasticsearchSecurityUser
	err := ctx.ReadResource("elasticstack:index/elasticstackElasticsearchSecurityUser:ElasticstackElasticsearchSecurityUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticstackElasticsearchSecurityUser resources.
type elasticstackElasticsearchSecurityUserState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchSecurityUserElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// The email of the user.
	Email *string `pulumi:"email"`
	// Specifies whether the user is enabled. The default value is true.
	Enabled *bool `pulumi:"enabled"`
	// The full name of the user.
	FullName *string `pulumi:"fullName"`
	// Arbitrary metadata that you want to associate with the user.
	Metadata *string `pulumi:"metadata"`
	// The user???s password. Passwords must be at least 6 characters long.
	Password *string `pulumi:"password"`
	// A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
	PasswordHash *string `pulumi:"passwordHash"`
	// A set of roles the user has. The roles determine the user???s access permissions. Default is [].
	Roles []string `pulumi:"roles"`
	// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
	Username *string `pulumi:"username"`
}

type ElasticstackElasticsearchSecurityUserState struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSecurityUserElasticsearchConnectionPtrInput
	// The email of the user.
	Email pulumi.StringPtrInput
	// Specifies whether the user is enabled. The default value is true.
	Enabled pulumi.BoolPtrInput
	// The full name of the user.
	FullName pulumi.StringPtrInput
	// Arbitrary metadata that you want to associate with the user.
	Metadata pulumi.StringPtrInput
	// The user???s password. Passwords must be at least 6 characters long.
	Password pulumi.StringPtrInput
	// A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
	PasswordHash pulumi.StringPtrInput
	// A set of roles the user has. The roles determine the user???s access permissions. Default is [].
	Roles pulumi.StringArrayInput
	// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
	Username pulumi.StringPtrInput
}

func (ElasticstackElasticsearchSecurityUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchSecurityUserState)(nil)).Elem()
}

type elasticstackElasticsearchSecurityUserArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection *ElasticstackElasticsearchSecurityUserElasticsearchConnection `pulumi:"elasticsearchConnection"`
	// The email of the user.
	Email *string `pulumi:"email"`
	// Specifies whether the user is enabled. The default value is true.
	Enabled *bool `pulumi:"enabled"`
	// The full name of the user.
	FullName *string `pulumi:"fullName"`
	// Arbitrary metadata that you want to associate with the user.
	Metadata *string `pulumi:"metadata"`
	// The user???s password. Passwords must be at least 6 characters long.
	Password *string `pulumi:"password"`
	// A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
	PasswordHash *string `pulumi:"passwordHash"`
	// A set of roles the user has. The roles determine the user???s access permissions. Default is [].
	Roles []string `pulumi:"roles"`
	// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a ElasticstackElasticsearchSecurityUser resource.
type ElasticstackElasticsearchSecurityUserArgs struct {
	// Elasticsearch connection configuration block.
	ElasticsearchConnection ElasticstackElasticsearchSecurityUserElasticsearchConnectionPtrInput
	// The email of the user.
	Email pulumi.StringPtrInput
	// Specifies whether the user is enabled. The default value is true.
	Enabled pulumi.BoolPtrInput
	// The full name of the user.
	FullName pulumi.StringPtrInput
	// Arbitrary metadata that you want to associate with the user.
	Metadata pulumi.StringPtrInput
	// The user???s password. Passwords must be at least 6 characters long.
	Password pulumi.StringPtrInput
	// A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
	PasswordHash pulumi.StringPtrInput
	// A set of roles the user has. The roles determine the user???s access permissions. Default is [].
	Roles pulumi.StringArrayInput
	// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
	Username pulumi.StringInput
}

func (ElasticstackElasticsearchSecurityUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticstackElasticsearchSecurityUserArgs)(nil)).Elem()
}

type ElasticstackElasticsearchSecurityUserInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSecurityUserOutput() ElasticstackElasticsearchSecurityUserOutput
	ToElasticstackElasticsearchSecurityUserOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserOutput
}

func (*ElasticstackElasticsearchSecurityUser) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchSecurityUser)(nil)).Elem()
}

func (i *ElasticstackElasticsearchSecurityUser) ToElasticstackElasticsearchSecurityUserOutput() ElasticstackElasticsearchSecurityUserOutput {
	return i.ToElasticstackElasticsearchSecurityUserOutputWithContext(context.Background())
}

func (i *ElasticstackElasticsearchSecurityUser) ToElasticstackElasticsearchSecurityUserOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSecurityUserOutput)
}

// ElasticstackElasticsearchSecurityUserArrayInput is an input type that accepts ElasticstackElasticsearchSecurityUserArray and ElasticstackElasticsearchSecurityUserArrayOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchSecurityUserArrayInput` via:
//
//	ElasticstackElasticsearchSecurityUserArray{ ElasticstackElasticsearchSecurityUserArgs{...} }
type ElasticstackElasticsearchSecurityUserArrayInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSecurityUserArrayOutput() ElasticstackElasticsearchSecurityUserArrayOutput
	ToElasticstackElasticsearchSecurityUserArrayOutputWithContext(context.Context) ElasticstackElasticsearchSecurityUserArrayOutput
}

type ElasticstackElasticsearchSecurityUserArray []ElasticstackElasticsearchSecurityUserInput

func (ElasticstackElasticsearchSecurityUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchSecurityUser)(nil)).Elem()
}

func (i ElasticstackElasticsearchSecurityUserArray) ToElasticstackElasticsearchSecurityUserArrayOutput() ElasticstackElasticsearchSecurityUserArrayOutput {
	return i.ToElasticstackElasticsearchSecurityUserArrayOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchSecurityUserArray) ToElasticstackElasticsearchSecurityUserArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSecurityUserArrayOutput)
}

// ElasticstackElasticsearchSecurityUserMapInput is an input type that accepts ElasticstackElasticsearchSecurityUserMap and ElasticstackElasticsearchSecurityUserMapOutput values.
// You can construct a concrete instance of `ElasticstackElasticsearchSecurityUserMapInput` via:
//
//	ElasticstackElasticsearchSecurityUserMap{ "key": ElasticstackElasticsearchSecurityUserArgs{...} }
type ElasticstackElasticsearchSecurityUserMapInput interface {
	pulumi.Input

	ToElasticstackElasticsearchSecurityUserMapOutput() ElasticstackElasticsearchSecurityUserMapOutput
	ToElasticstackElasticsearchSecurityUserMapOutputWithContext(context.Context) ElasticstackElasticsearchSecurityUserMapOutput
}

type ElasticstackElasticsearchSecurityUserMap map[string]ElasticstackElasticsearchSecurityUserInput

func (ElasticstackElasticsearchSecurityUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchSecurityUser)(nil)).Elem()
}

func (i ElasticstackElasticsearchSecurityUserMap) ToElasticstackElasticsearchSecurityUserMapOutput() ElasticstackElasticsearchSecurityUserMapOutput {
	return i.ToElasticstackElasticsearchSecurityUserMapOutputWithContext(context.Background())
}

func (i ElasticstackElasticsearchSecurityUserMap) ToElasticstackElasticsearchSecurityUserMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticstackElasticsearchSecurityUserMapOutput)
}

type ElasticstackElasticsearchSecurityUserOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSecurityUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticstackElasticsearchSecurityUser)(nil)).Elem()
}

func (o ElasticstackElasticsearchSecurityUserOutput) ToElasticstackElasticsearchSecurityUserOutput() ElasticstackElasticsearchSecurityUserOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityUserOutput) ToElasticstackElasticsearchSecurityUserOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserOutput {
	return o
}

// Elasticsearch connection configuration block.
func (o ElasticstackElasticsearchSecurityUserOutput) ElasticsearchConnection() ElasticstackElasticsearchSecurityUserElasticsearchConnectionPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) ElasticstackElasticsearchSecurityUserElasticsearchConnectionPtrOutput {
		return v.ElasticsearchConnection
	}).(ElasticstackElasticsearchSecurityUserElasticsearchConnectionPtrOutput)
}

// The email of the user.
func (o ElasticstackElasticsearchSecurityUserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// Specifies whether the user is enabled. The default value is true.
func (o ElasticstackElasticsearchSecurityUserOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The full name of the user.
func (o ElasticstackElasticsearchSecurityUserOutput) FullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringPtrOutput { return v.FullName }).(pulumi.StringPtrOutput)
}

// Arbitrary metadata that you want to associate with the user.
func (o ElasticstackElasticsearchSecurityUserOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringOutput { return v.Metadata }).(pulumi.StringOutput)
}

// The user???s password. Passwords must be at least 6 characters long.
func (o ElasticstackElasticsearchSecurityUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// A hash of the user???s password. This must be produced using the same hashing algorithm as has been configured for password storage (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#hashing-settings).
func (o ElasticstackElasticsearchSecurityUserOutput) PasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringPtrOutput { return v.PasswordHash }).(pulumi.StringPtrOutput)
}

// A set of roles the user has. The roles determine the user???s access permissions. Default is [].
func (o ElasticstackElasticsearchSecurityUserOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// An identifier for the user (see https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html#security-api-put-user-path-params).
func (o ElasticstackElasticsearchSecurityUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticstackElasticsearchSecurityUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type ElasticstackElasticsearchSecurityUserArrayOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSecurityUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElasticstackElasticsearchSecurityUser)(nil)).Elem()
}

func (o ElasticstackElasticsearchSecurityUserArrayOutput) ToElasticstackElasticsearchSecurityUserArrayOutput() ElasticstackElasticsearchSecurityUserArrayOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityUserArrayOutput) ToElasticstackElasticsearchSecurityUserArrayOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserArrayOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityUserArrayOutput) Index(i pulumi.IntInput) ElasticstackElasticsearchSecurityUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchSecurityUser {
		return vs[0].([]*ElasticstackElasticsearchSecurityUser)[vs[1].(int)]
	}).(ElasticstackElasticsearchSecurityUserOutput)
}

type ElasticstackElasticsearchSecurityUserMapOutput struct{ *pulumi.OutputState }

func (ElasticstackElasticsearchSecurityUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElasticstackElasticsearchSecurityUser)(nil)).Elem()
}

func (o ElasticstackElasticsearchSecurityUserMapOutput) ToElasticstackElasticsearchSecurityUserMapOutput() ElasticstackElasticsearchSecurityUserMapOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityUserMapOutput) ToElasticstackElasticsearchSecurityUserMapOutputWithContext(ctx context.Context) ElasticstackElasticsearchSecurityUserMapOutput {
	return o
}

func (o ElasticstackElasticsearchSecurityUserMapOutput) MapIndex(k pulumi.StringInput) ElasticstackElasticsearchSecurityUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElasticstackElasticsearchSecurityUser {
		return vs[0].(map[string]*ElasticstackElasticsearchSecurityUser)[vs[1].(string)]
	}).(ElasticstackElasticsearchSecurityUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSecurityUserInput)(nil)).Elem(), &ElasticstackElasticsearchSecurityUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSecurityUserArrayInput)(nil)).Elem(), ElasticstackElasticsearchSecurityUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElasticstackElasticsearchSecurityUserMapInput)(nil)).Elem(), ElasticstackElasticsearchSecurityUserMap{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSecurityUserOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSecurityUserArrayOutput{})
	pulumi.RegisterOutputType(ElasticstackElasticsearchSecurityUserMapOutput{})
}
