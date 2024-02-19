// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package cluster

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if len(a.ko.Spec.CapacityProviders) != len(b.ko.Spec.CapacityProviders) {
		delta.Add("Spec.CapacityProviders", a.ko.Spec.CapacityProviders, b.ko.Spec.CapacityProviders)
	} else if len(a.ko.Spec.CapacityProviders) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CapacityProviders, b.ko.Spec.CapacityProviders) {
			delta.Add("Spec.CapacityProviders", a.ko.Spec.CapacityProviders, b.ko.Spec.CapacityProviders)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Configuration, b.ko.Spec.Configuration) {
		delta.Add("Spec.Configuration", a.ko.Spec.Configuration, b.ko.Spec.Configuration)
	} else if a.ko.Spec.Configuration != nil && b.ko.Spec.Configuration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration, b.ko.Spec.Configuration.ExecuteCommandConfiguration) {
			delta.Add("Spec.Configuration.ExecuteCommandConfiguration", a.ko.Spec.Configuration.ExecuteCommandConfiguration, b.ko.Spec.Configuration.ExecuteCommandConfiguration)
		} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID, b.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID) {
				delta.Add("Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID", a.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID, b.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID)
			} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID != nil {
				if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID", a.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID, b.ko.Spec.Configuration.ExecuteCommandConfiguration.KMSKeyID)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration) {
				delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration)
			} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled) {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled)
				} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != nil {
					if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled {
						delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchEncryptionEnabled)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName) {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName)
				} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != nil {
					if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName {
						delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.CloudWatchLogGroupName)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName) {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName)
				} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != nil {
					if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName {
						delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3BucketName)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled) {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled)
				} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != nil {
					if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled {
						delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3EncryptionEnabled)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix) {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix)
				} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != nil {
					if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix {
						delta.Add("Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix", a.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix, b.ko.Spec.Configuration.ExecuteCommandConfiguration.LogConfiguration.S3KeyPrefix)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging, b.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging) {
				delta.Add("Spec.Configuration.ExecuteCommandConfiguration.Logging", a.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging, b.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging)
			} else if a.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging != nil && b.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging != nil {
				if *a.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging != *b.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging {
					delta.Add("Spec.Configuration.ExecuteCommandConfiguration.Logging", a.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging, b.ko.Spec.Configuration.ExecuteCommandConfiguration.Logging)
				}
			}
		}
	}
	if len(a.ko.Spec.DefaultCapacityProviderStrategy) != len(b.ko.Spec.DefaultCapacityProviderStrategy) {
		delta.Add("Spec.DefaultCapacityProviderStrategy", a.ko.Spec.DefaultCapacityProviderStrategy, b.ko.Spec.DefaultCapacityProviderStrategy)
	} else if len(a.ko.Spec.DefaultCapacityProviderStrategy) > 0 {
		if !reflect.DeepEqual(a.ko.Spec.DefaultCapacityProviderStrategy, b.ko.Spec.DefaultCapacityProviderStrategy) {
			delta.Add("Spec.DefaultCapacityProviderStrategy", a.ko.Spec.DefaultCapacityProviderStrategy, b.ko.Spec.DefaultCapacityProviderStrategy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ServiceConnectDefaults, b.ko.Spec.ServiceConnectDefaults) {
		delta.Add("Spec.ServiceConnectDefaults", a.ko.Spec.ServiceConnectDefaults, b.ko.Spec.ServiceConnectDefaults)
	} else if a.ko.Spec.ServiceConnectDefaults != nil && b.ko.Spec.ServiceConnectDefaults != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ServiceConnectDefaults.Namespace, b.ko.Spec.ServiceConnectDefaults.Namespace) {
			delta.Add("Spec.ServiceConnectDefaults.Namespace", a.ko.Spec.ServiceConnectDefaults.Namespace, b.ko.Spec.ServiceConnectDefaults.Namespace)
		} else if a.ko.Spec.ServiceConnectDefaults.Namespace != nil && b.ko.Spec.ServiceConnectDefaults.Namespace != nil {
			if *a.ko.Spec.ServiceConnectDefaults.Namespace != *b.ko.Spec.ServiceConnectDefaults.Namespace {
				delta.Add("Spec.ServiceConnectDefaults.Namespace", a.ko.Spec.ServiceConnectDefaults.Namespace, b.ko.Spec.ServiceConnectDefaults.Namespace)
			}
		}
	}
	if len(a.ko.Spec.Settings) != len(b.ko.Spec.Settings) {
		delta.Add("Spec.Settings", a.ko.Spec.Settings, b.ko.Spec.Settings)
	} else if len(a.ko.Spec.Settings) > 0 {
		if !reflect.DeepEqual(a.ko.Spec.Settings, b.ko.Spec.Settings) {
			delta.Add("Spec.Settings", a.ko.Spec.Settings, b.ko.Spec.Settings)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}