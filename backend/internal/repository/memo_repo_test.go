package repository

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gridea-pro/backend/internal/domain"
)

func TestMemoListSortedByCreatedAtDesc(t *testing.T) {
	dir := t.TempDir()
	_ = os.MkdirAll(filepath.Join(dir, "config"), 0755)
	repo := NewMemoRepository(dir)
	ctx := context.Background()

	now := time.Now()
	stored := []domain.Memo{
		{ID: "mid", CreatedAt: now.Add(-2 * time.Hour), UpdatedAt: now, Tags: []string{}, Images: []string{}},
		{ID: "new", CreatedAt: now, UpdatedAt: now, Tags: []string{}, Images: []string{}},
		{ID: "old", CreatedAt: now.Add(-5 * time.Hour), UpdatedAt: now, Tags: []string{}, Images: []string{}},
	}
	if err := repo.SaveAll(ctx, stored); err != nil {
		t.Fatalf("SaveAll: %v", err)
	}

	got, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	want := []string{"new", "mid", "old"}
	if len(got) != len(want) {
		t.Fatalf("len=%d want=%d", len(got), len(want))
	}
	for i, id := range want {
		if got[i].ID != id {
			gotIDs := make([]string, len(got))
			for j, m := range got {
				gotIDs[j] = m.ID
			}
			t.Fatalf("排序错误：got=%v want=%v", gotIDs, want)
		}
	}
}

func TestMemoListReturnsCopy(t *testing.T) {
	dir := t.TempDir()
	_ = os.MkdirAll(filepath.Join(dir, "config"), 0755)
	repo := NewMemoRepository(dir)
	ctx := context.Background()

	now := time.Now()
	stored := []domain.Memo{
		{ID: "a", CreatedAt: now.Add(-3 * time.Hour), Tags: []string{}, Images: []string{}},
		{ID: "b", CreatedAt: now.Add(-2 * time.Hour), Tags: []string{}, Images: []string{}},
		{ID: "c", CreatedAt: now.Add(-1 * time.Hour), Tags: []string{}, Images: []string{}},
	}
	if err := repo.SaveAll(ctx, stored); err != nil {
		t.Fatalf("SaveAll: %v", err)
	}

	out1, _ := repo.List(ctx)
	if len(out1) > 0 {
		out1[0].ID = "X"
	}
	out2, _ := repo.List(ctx)
	for _, m := range out2 {
		if m.ID == "X" {
			t.Fatalf("List 返回的不是副本：外部修改污染了 cache")
		}
	}
}
