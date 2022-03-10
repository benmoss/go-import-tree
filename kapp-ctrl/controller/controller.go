package foo

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

func Stuff() {
	ctrl.Log.WithName("builder-examples")

}
