/*
Copyright 2024 YANDEX LLC

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

package services

const MultiNodeReaderOnlyPublishFlag = "MULTI_NODE_READER_ONLY"
const MultiNodeReaderOnlyPublishFlagValue = "true"

func ROPublishContext(ro bool) map[string]string {
	if ro {
		return map[string]string{
			MultiNodeReaderOnlyPublishFlag: MultiNodeReaderOnlyPublishFlagValue,
		}
	}

	return nil
}

func NodeStageRO(ctx map[string]string) bool {
	if ro, ok := ctx[MultiNodeReaderOnlyPublishFlag]; ok && ro == MultiNodeReaderOnlyPublishFlagValue {
		return true
	}

	return false
}
