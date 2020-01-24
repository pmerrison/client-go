// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1beta1

import (
	securityv1beta1 "istio.io/api/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationPolicy enables access control on workloads.
//
// For example, the following authorization policy denies all requests to workloads
// in namespace foo.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: deny-all
//  namespace: foo
// spec:
//   {}
// ```
//
// The following authorization policy allows all requests to workloads in namespace
// foo.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: allow-all
//  namespace: foo
// spec:
//  rules:
//  - {}
// ```
//
// <!-- crd generation tags
// +cue-gen:AuthorizationPolicy:groupName:security.istio.io
// +cue-gen:AuthorizationPolicy:version:v1beta1
// +cue-gen:AuthorizationPolicy:storageVersion
// +cue-gen:AuthorizationPolicy:annotations:helm.sh/resource-policy=keep
// +cue-gen:AuthorizationPolicy:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:AuthorizationPolicy:subresource:status
// +cue-gen:AuthorizationPolicy:scope:Namespaced
// +cue-gen:AuthorizationPolicy:resource:categories=istio-io,security-istio-io,plural=authorizationpolicies
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type AuthorizationPolicy struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1beta1.AuthorizationPolicy `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthorizationPolicyList is a collection of AuthorizationPolicies.
type AuthorizationPolicyList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []AuthorizationPolicy `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RequestAuthentication defines what request authentication methods are supported by a workload.
// If will reject a request if the request contains invalid authentication information, based on the
// configured authentication rules. A request that does not contain any authentication credentials
// will be accepted but will not have any authenticated identity. To restrict access to authenticated
// requests only, this should be accompanied by an authorization rule.
// Examples:
//
// - Require JWT for all request for workloads that have label `app:httpbin`
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   jwtRules:
//   - issuer: "issuer-foo"
//     jwksUri: https://example.com/.well-known/jwks.json
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
//  rules:
//  - from:
//    - source:
//        requestPrincipals: ["*"]
// ```
//
// - The next example shows how to set a different JWT requirement for a different `host`. The `RequestAuthentication`
// declares it can accpet JWTs issuer by either `issuer-foo` or `issuer-bar` (the public key set is implicitly
// set from the OpenID Connect spec).
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   jwtRules:
//   - issuer: "issuer-foo"
//   - issuer: "issuer-bar"
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
//  rules:
//  - from:
//    - source:
//        requestPrincipals: ["issuer-foo/*"]
//    to:
//      hosts: ["example.com"]
//  - from:
//    - source:
//        requestPrincipals: ["issuer-bar/*"]
//    to:
//      hosts: ["another-host.com"]
// ```
//
// - You can fine tune the authorization policy to set different requirement per path. For example,
// to require JWT on all paths, except /healthz, the same `RequestAuthentication` can be used, but the
// authorization policy could be:
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
//  rules:
//  - from:
//    - source:
//        requestPrincipals: ["*"]
//  - to:
//    - operation:
//        paths: ["/healthz]
// ```
//
// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:version:v1beta1
// +cue-gen:RequestAuthentication:storageVersion
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RequestAuthentication struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec securityv1beta1.RequestAuthentication `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RequestAuthenticationList is a collection of RequestAuthentications.
type RequestAuthenticationList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RequestAuthentication `json:"items" protobuf:"bytes,2,rep,name=items"`
}