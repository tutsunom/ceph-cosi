/*
Copyright 2021 The Ceph-COSI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

import (
	"context"

	"github.com/ceph/cosi-driver-ceph/pkg/util/s3client"
	rgwadmin "github.com/ceph/go-ceph/rgw/admin"
	"k8s.io/klog/v2"
)

func NewDriver(ctx context.Context, provisioner, rgwEndpoint, accessKey, secretKey string) (*IdentityServer, *ProvisionerServer, error) {
	// TODO : use different user this operation
	klog.Info(accessKey, " ", secretKey, " ", rgwEndpoint)
	s3Client, err := s3client.NewS3Agent(accessKey, secretKey, rgwEndpoint, true)
	if err != nil {
		klog.Fatalln(err)
	}
	//TODO : add support for TLS endpoint
	rgwAdminClient, err := rgwadmin.New(rgwEndpoint, accessKey, secretKey, nil)
	if err != nil {
		klog.Fatalln(err)
	}
	return &IdentityServer{
			provisioner: provisioner,
		}, &ProvisionerServer{
			provisioner:    provisioner,
			s3Client:       s3Client,
			rgwAdminClient: rgwAdminClient,
		}, nil
}
