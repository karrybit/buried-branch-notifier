package command

import (
	"testing"
)

const target = `origin/HEAD -> origin/develop
				origin/develop
				origin/hoge
				origin/fuga
				origin/piyo`

func TestSplitBranches(t *testing.T) {
	expect := []string{"origin/HEAD -> origin/develop", "origin/develop", "origin/hoge", "origin/fuga", "origin/piyo"}
	splited := splitBranches(target)
	for i := range splited {
		if splited[i] != expect[i] {
			t.Errorf("got %+v, expect %+v", splited, expect)
		}
	}
}

func TestRemoveHeadBranchIndex(t *testing.T) {
	expect := []string{"origin/develop", "origin/hoge", "origin/fuga", "origin/piyo"}
	removed := removeHeadBranch(splitBranches(target))
	for i := range removed {
		if removed[i] != expect[i] {
			t.Errorf("got %+v, expect %+v", removed, expect)
		}
	}
}
