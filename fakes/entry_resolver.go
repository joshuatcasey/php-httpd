package fakes

import (
	"sync"

	packit "github.com/paketo-buildpacks/packit/v2"
)

type EntryResolver struct {
	MergeLayerTypesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Name    string
			Entries []packit.BuildpackPlanEntry
		}
		Returns struct {
			Launch bool
			Build  bool
		}
		Stub func(string, []packit.BuildpackPlanEntry) (bool, bool)
	}
}

func (f *EntryResolver) MergeLayerTypes(param1 string, param2 []packit.BuildpackPlanEntry) (bool, bool) {
	f.MergeLayerTypesCall.mutex.Lock()
	defer f.MergeLayerTypesCall.mutex.Unlock()
	f.MergeLayerTypesCall.CallCount++
	f.MergeLayerTypesCall.Receives.Name = param1
	f.MergeLayerTypesCall.Receives.Entries = param2
	if f.MergeLayerTypesCall.Stub != nil {
		return f.MergeLayerTypesCall.Stub(param1, param2)
	}
	return f.MergeLayerTypesCall.Returns.Launch, f.MergeLayerTypesCall.Returns.Build
}
