package testing

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	grace              = int64(30)
	enableServiceLinks = corev1.DefaultEnableServiceLinks
)

func TestData() (*corev1.PodList, *corev1.ServiceList, *corev1.ReplicationControllerList) {
	pods := &corev1.PodList{
		ListMeta: metav1.ListMeta{
			ResourceVersion: "15",
		},
		Items: []corev1.Pod{
			{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "test", ResourceVersion: "10"},
				Spec: corev1.PodSpec{
					RestartPolicy:                 corev1.RestartPolicyAlways,
					DNSPolicy:                     corev1.DNSClusterFirst,
					TerminationGracePeriodSeconds: &grace,
					SecurityContext:               &corev1.PodSecurityContext{},
					EnableServiceLinks:            &enableServiceLinks,
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "bar", Namespace: "test", ResourceVersion: "11"},
				Spec: corev1.PodSpec{
					RestartPolicy:                 corev1.RestartPolicyAlways,
					DNSPolicy:                     corev1.DNSClusterFirst,
					TerminationGracePeriodSeconds: &grace,
					SecurityContext:               &corev1.PodSecurityContext{},
					EnableServiceLinks:            &enableServiceLinks,
				},
			},
		},
	}
	svc := &corev1.ServiceList{
		ListMeta: metav1.ListMeta{
			ResourceVersion: "16",
		},
		Items: []corev1.Service{
			{
				ObjectMeta: metav1.ObjectMeta{Name: "baz", Namespace: "test", ResourceVersion: "12"},
				Spec: corev1.ServiceSpec{
					SessionAffinity: "None",
					Type:            corev1.ServiceTypeClusterIP,
				},
			},
		},
	}

	one := int32(1)
	rc := &corev1.ReplicationControllerList{
		ListMeta: metav1.ListMeta{
			ResourceVersion: "17",
		},
		Items: []corev1.ReplicationController{
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rc1", Namespace: "test", ResourceVersion: "18"},
				Spec: corev1.ReplicationControllerSpec{
					Replicas: &one,
				},
			},
		},
	}
	return pods, svc, rc
}
