// Copyright 2020 apirator.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apimock

import (
	"context"

	apirator "github.com/apirator/apirator/pkg/apis/apirator/v1alpha1"
)

func (r *ReconcileAPIMock) markAsSuccessful(obj *apirator.APIMock) error {
	log.Info("Updating APIMock with status Provisioned...", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
	err := r.updateStatus(obj, apirator.PROVISIONED)
	if err != nil {
		log.Error(err, "Failed to update APIMock with Provisioned status", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
		return err
	}
	log.Info("Status Provisioned update successfully", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
	return nil
}

func (r *ReconcileAPIMock) markAsFailure(obj *apirator.APIMock) error {
	log.Info("Updating APIMock with status Error...", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
	err := r.updateStatus(obj, apirator.ERROR)
	if err != nil {
		log.Error(err, "Failed to update APIMock with Provisioned status", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
		return err
	}
	log.Info("Status Error update successfully", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
	return nil
}

func (r *ReconcileAPIMock) markAsInvalidOAS(obj *apirator.APIMock) error {
	log.Info("Updating APIMock with status OAS invalid...", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
	err := r.updateStatus(obj, apirator.INVALID_OAS)
	if err != nil {
		log.Error(err, "Failed to update APIMock with Invalid OAS", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
		return err
	}
	log.Info("Status Invalid OAS update successfully", "APIMock.Namespace", obj.Namespace, "APIMock.Name", obj.Name)
	return nil
}

func (r *ReconcileAPIMock) updateStatus(obj *apirator.APIMock, status string) error {
	if obj.Status.Phase != status {
		obj.Status.Phase = status
		err := r.client.Status().Update(context.TODO(), obj)
		if err != nil {
			return err
		}
	}
	return nil
}
