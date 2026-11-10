/*
Copyright 2026 Peter Rozsa.

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

package main

import (
	"maps"
	"slices"
	"testing"
)

func TestWatchNamespaces(t *testing.T) {
	for _, tc := range []struct {
		value string
		want  []string // nil: watch every namespace
	}{
		{value: "", want: nil},
		{value: " , ", want: nil},
		{value: "impala", want: []string{"impala"}},
		{value: "impala-operator, team-a ,team-a,", want: []string{"impala-operator", "team-a"}},
	} {
		got := watchNamespaces(tc.value)
		if tc.want == nil {
			if got != nil {
				t.Errorf("watchNamespaces(%q) = %v, want nil", tc.value, got)
			}
			continue
		}
		if keys := slices.Sorted(maps.Keys(got)); !slices.Equal(keys, tc.want) {
			t.Errorf("watchNamespaces(%q) = %v, want %v", tc.value, keys, tc.want)
		}
	}
}
