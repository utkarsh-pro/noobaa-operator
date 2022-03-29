package bucketclass

import (
	nbv1 "github.com/noobaa/noobaa-operator/v5/pkg/apis/noobaa/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Verify Bucketclass provisioner actions", func() {
	Context("When bucketclass is in the same namespace as NooBaa system", func() {
		It("should allow object for the provisioner", func() {
			systemNS := "test"

			obj := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: systemNS,
				},
			}

			Expect(isObjectForProvisioner(obj, systemNS)).To(BeTrue())
		})
	})

	Context("When bucketclass is not in the same namespace as NooBaa system", func() {
		It("should disallow object for the provisioner", func() {
			systemNS := "test"

			obj := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "random",
				},
			}

			Expect(isObjectForProvisioner(obj, systemNS)).To(BeFalse())
		})
	})

	Context("When bucketclass is not in the same namespace as NooBaa system: Valid provisioner label", func() {
		It("should allow object for the provisioner", func() {
			systemNS := "test"

			obj := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "random",
					Labels: map[string]string{
						"provisioner": systemNS,
					},
				},
			}

			Expect(isObjectForProvisioner(obj, systemNS)).To(BeTrue())
		})
	})

	Context("When bucketclass is not in the same namespace as NooBaa system: Invalid provisioner label", func() {
		It("should disallow object for the provisioner", func() {
			systemNS := "test"

			obj := &nbv1.BucketClass{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "random",
					Labels: map[string]string{
						"provisioner": "xyz",
					},
				},
			}

			Expect(isObjectForProvisioner(obj, systemNS)).To(BeFalse())
		})
	})
})
