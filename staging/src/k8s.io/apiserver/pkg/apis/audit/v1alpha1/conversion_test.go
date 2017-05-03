package v1alpha1

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/apis/audit"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	authnv1 "k8s.io/client-go/pkg/apis/authentication/v1"
)

func TestConversion(t *testing.T) {
	e := Event{
		Level:      LevelRequestObject,
		Timestamp:  metav1.Now(),
		AuditID:    "123",
		RequestURI: "foo.bar/com",
		Verb:       "get",
		User: authnv1.UserInfo{
			Username: "bob",
			UID:      "456",
			Groups:   []string{"users", "admins"},
			Extra:    map[string]authnv1.ExtraValue{"key": authnv1.ExtraValue{"1", "2"}},
		},
		Impersonate: &authnv1.UserInfo{
			Username: "sam",
			UID:      "456789",
			Groups:   []string{"sams"},
			Extra:    map[string]authnv1.ExtraValue{"stuff": authnv1.ExtraValue{"3", "4"}},
		},
		SourceIP: "12.12.12.12",
		ObjectRef: &apiv1.ObjectReference{
			Kind:            "foo",
			Namespace:       "bar",
			Name:            "baz",
			UID:             "9087",
			APIVersion:      "v1",
			ResourceVersion: "987234",
		},
	}

	internal := audit.Event{}
	if err := runtime.NewScheme().Convert(&e, &internal, nil); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
