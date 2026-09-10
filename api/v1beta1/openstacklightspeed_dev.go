/*
Copyright 2026

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import "encoding/json"

// ParseDevConfig unmarshals the Dev RawExtension into a DevSpec.
// Returns a zero-value DevSpec and an error on malformed input.
func (instance *OpenStackLightspeed) ParseDevConfig() (DevSpec, error) {
	var devConfig DevSpec
	if len(instance.Spec.Dev.Raw) > 0 {
		if err := json.Unmarshal(instance.Spec.Dev.Raw, &devConfig); err != nil {
			return devConfig, err
		}
	}
	return devConfig, nil
}
